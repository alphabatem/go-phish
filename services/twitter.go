package services

import (
	go_phish "github.com/alphabatem/go-phish"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type TwitterService struct {
	Service

	client  *twitter.Client
	stream  *twitter.Stream
	demux   twitter.SwitchDemux
	filters []string

	mediaEnabledUsers map[string]bool

	asvc *AnalyserService
}

const TWITTER_SVC = "twitter_svc"

// Service ID
func (svc *TwitterService) Id() string {
	return TWITTER_SVC
}

func (svc *TwitterService) Start() error {

	config := oauth1.NewConfig(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"))
	token := oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	httpClient.Timeout = 10 * 1000 * time.Millisecond

	svc.filters = []string{"$btc", "$eth", "$ada", "crypto", "#crypto", "#ifo"} //TODO populate from market service
	svc.mediaEnabledUsers = map[string]bool{}

	svc.client = twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	user, _, err := svc.client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return err
	}
	log.Printf("[TWITTER] Authenticated as: %s", user.Name)

	svc.demux = twitter.NewSwitchDemux()
	svc.demux.Tweet = svc.tweetHandler
	svc.demux.Event = svc.eventHandler

	svc.asvc = svc.Service(ANALYSER_SVC).(*AnalyserService)

	return nil
}

// Enables OCR procssing for meida items a given user posts when it comes through the stream
func (svc *TwitterService) EnableMediaProcessing(users ...string) {
	for _, u := range users {
		svc.mediaEnabledUsers[u] = true
	}
}

func (svc *TwitterService) UserTweets(username string) ([]twitter.Tweet, error) {
	replies := false
	userTimelineParams := &twitter.UserTimelineParams{ScreenName: username, ExcludeReplies: &replies}
	tweets, resp, err := svc.client.Timelines.UserTimeline(userTimelineParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.Printf("Tweet Len: %v", len(tweets))
	if len(tweets) == 0 {
		dat, _ := ioutil.ReadAll(resp.Body)
		log.Printf("%s", dat)
		return tweets, nil
	}

	for {
		userTimelineParams = &twitter.UserTimelineParams{ScreenName: username, ExcludeReplies: &replies, Count: 200, MaxID: tweets[len(tweets)-1].ID}
		tweets2, _, err := svc.client.Timelines.UserTimeline(userTimelineParams)
		if err != nil {
			return tweets, err
		}

		if len(tweets2) <= 1 { //Out of tweets
			return tweets, nil
		}

		log.Printf("Tweet2 Len: %v", len(tweets2))
		tweets = append(tweets, tweets2...)
		log.Printf("Total Len: %v", len(tweets))
	}

	return tweets, nil
}

func (svc *TwitterService) UserStream(filters []string) error {
	params := &twitter.StreamFilterParams{
		Follow: filters,
	}
	stream, err := svc.client.Streams.Filter(params)
	if err != nil {
		return err
	}

	log.Printf("Starting User stream handler")
	go svc.demux.HandleChan(stream.Messages)

	return nil
}

func (svc *TwitterService) startStream() error {
	filterParams := &twitter.StreamFilterParams{
		Track:         svc.filters,
		Language:      []string{"en", "enUs"},
		StallWarnings: twitter.Bool(true),
	}

	var err error
	svc.stream, err = svc.client.Streams.Filter(filterParams)
	if err != nil {
		return err
	}

	log.Printf("Starting Filter stream handler")
	go svc.demux.HandleChan(svc.stream.Messages)

	return nil
}

func (svc *TwitterService) tweetHandler(tweet *twitter.Tweet) {
	text := tweet.Text
	if tweet.ExtendedTweet != nil {
		text = tweet.ExtendedTweet.FullText
	}

	in := go_phish.FeedInput{
		Text:     text,
		Platform: "twitter",
		User:     tweet.User.IDStr,
	}

	svc.asvc.InboundChan <- &in
}

func (svc *TwitterService) eventHandler(event *twitter.Event) {
	log.Printf("%#v\n", event)
}

func (svc *TwitterService) Shutdown() {
	svc.stream.Stop()
}
