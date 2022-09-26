package dto

import "time"

//TODO This should be on chain eventually

type SiteDomainResponse struct {
	CID          string   `json:"cid"`
	BlokDomains  []string `json:"blok_domains"`
	CustomDomain string   `json:"custom_domain"`
}

type SiteSSLResponse struct {
	Provider  string     `json:"provider"`
	Domain    string     `json:"domain"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type SiteStatsResponse struct {
	Stats       SiteEpochDetail `json:"stats"`
	BlockHeight uint64          `json:"block_height"`
}

// Stats per epoch/block on the network
type SiteEpochDetail struct {
	Requests     uint64 `json:"requests"`
	RawRequests  uint64 `json:"raw_requests"` //Requests that did NOT hit caching mechanisms
	Bandwidth    uint64 `json:"bandwidth"`
	BuildMinutes uint8  `json:"build_minutes"`
}

type Epoch struct {
	Blocks []*Block
	Hash   string
}

type Block struct {
	Requests []*Request
	Hash     string
}

type Request struct {
	DriveID    string //solana.PublicKey
	NodeID     string //solana.PublicKey
	Path       string
	Nonce      string //Used to verify sha hash
	Hash       string
	CacheHit   bool
	ProviderID string //solana.PublicKey - Node who provided the data
}
