package services

import (
	"log"
	"strings"
)

type SentimentCallback func(in *market_sentiment.MarketInput) error

type AnalyserService struct {
	Service

	InboundChan chan *market_sentiment.MarketInput

	tokenListeners map[string][]SentimentCallback
	userListeners  map[string][]SentimentCallback

	bloksvc *BlockListService
}

const ANALYSER_SVC = "analyser_svc"

func (svc AnalyserService) Id() string {
	return ANALYSER_SVC
}

func (svc *AnalyserService) Start() error {
	svc.bloksvc = svc.Service(BLOCKLIST_SVC).(*BlockListService)

	svc.tokenListeners = map[string][]SentimentCallback{}
	svc.userListeners = map[string][]SentimentCallback{}

	svc.InboundChan = make(chan *market_sentiment.MarketInput, 1000) //TODO listen
	go svc.Worker()                                                  //Handles inbound analysis reqs

	return nil
}

func (svc *AnalyserService) Worker() {
	log.Printf("Analyser worker started")
	for r := range svc.InboundChan {
		_, err := svc.Analyse(r)
		if err != nil {
			log.Printf("Analyser error: %s", err)
		}
	}
	log.Printf("Analyser worker stopped")
}

// TODO Input
func (svc *AnalyserService) Analyse(in *market_sentiment.MarketInput) (*market_sentiment.MarketInput, error) {

	sentiment, err := svc.ssvc.Score(in, in.Text)
	if err != nil {
		return in, nil
	}

	// Get the context sentiment first (+ correlation)
	// if the reply is positive they agree with this sentiment (this is the resultant sentiment), otherwise its the opposite
	if in.ContextText != "" {
		contextSentiment, err := svc.ssvc.Score(in, in.ContextText) //TODO Add HASH cache here
		if err != nil {
			return in, nil
		}

		if sentiment >= 0 {
			in.Sentiment = contextSentiment
		} else {
			in.Sentiment = sentiment
		}
	} else {
		in.Sentiment = sentiment //We just use the authors sentiment
	}

	in.Buzz, err = svc.isvc.ScoreBuzz(in)
	if err != nil {
		return in, nil
	}

	cleanText := strings.TrimSpace(in.Text)
	cleanText = strings.ReplaceAll(cleanText, "\n", "")
	//log.Printf("[TWITTER] %s || Symbols: %v - Sentiment: %v (%v) - Buzz: %v", cleanText, tok, svc.ssvc.ToEmoji(in.Sentiment), in.Sentiment, in.Buzz)

	go svc.Notify(in)

	return in, nil
}

func (svc *AnalyserService) SummarizeMinute(minStart int) {
	//TODO Token Loop
	//TODO score sentiment for each token for the minutes input data

	//TODO add to global market sentiment score
}

func (svc *AnalyserService) ParseCorrelations(in *market_sentiment.MarketInput) []string {
	tok := []string{}
	tokL := map[string]bool{}
	for r := range in.Correlations {
		r = strings.ReplaceAll(r, "$", "")
		if _, ok := tokL[r]; ok { //Skip as added
			continue
		}
		tok = append(tok, r)
	}

	return tok
}

func (svc *AnalyserService) RegisterTokenListener(token string, callback SentimentCallback) error {
	t := strings.ToUpper(token)
	log.Printf("Registering listener for token: %s", t)

	if _, ok := svc.tokenListeners[t]; !ok {
		svc.tokenListeners[t] = []SentimentCallback{}
	}

	svc.tokenListeners[t] = append(svc.tokenListeners[t], callback)

	return nil
}

func (svc *AnalyserService) RegisterUserListener(user string, callback SentimentCallback) error {
	log.Printf("Registering listener for user: %s", user)

	if _, ok := svc.userListeners[user]; !ok {
		svc.userListeners[user] = []SentimentCallback{}
	}

	svc.userListeners[user] = append(svc.userListeners[user], callback)

	return nil
}

func (svc *AnalyserService) Notify(in *market_sentiment.MarketInput) error {
	_ = svc.notifyUserListeners(in.User, in)

	for corr := range in.Correlations {
		token := strings.ToUpper(corr)
		_ = svc.notifyTokenListeners(token, in)
	}

	return nil
}

func (svc *AnalyserService) notifyTokenListeners(token string, in *market_sentiment.MarketInput) error {
	if _, ok := svc.tokenListeners[token]; !ok {
		return nil //No listeners registered
	}

	log.Printf("Notifying token listeners: %s (%v)", token, len(svc.tokenListeners[token]))
	for _, lis := range svc.tokenListeners[token] {
		_ = lis(in)
	}

	return nil
}

func (svc *AnalyserService) notifyUserListeners(user string, in *market_sentiment.MarketInput) error {
	if _, ok := svc.userListeners[user]; !ok {
		return nil //No listeners registered
	}

	log.Printf("Notifying user listeners: %s (%v)", user, len(svc.userListeners[user]))
	for _, lis := range svc.userListeners[user] {
		_ = lis(in)
	}

	return nil
}
