package go_phish

import "time"

type FeedInput struct {
	ID          uint       `json:"-"`
	ContextText string     `json:"context_text"` //The context of the text (reply or a subject etc)
	Text        string     `json:"text"`
	User        string     `json:"user"`
	Platform    string     `json:"platform"`
	Media       []string   `json:"media"`
	CreatedAt   time.Time  `json:"created_at"`
	ProcessedAt *time.Time `json:"processed_at"`
}
