package dto

type SignedMessage struct {
	PublicKey string `json:"publicKey"`
	Signature struct {
		Data []byte `json:"data"`
		Type string `json:"type"`
	} `json:"signature"`
}
