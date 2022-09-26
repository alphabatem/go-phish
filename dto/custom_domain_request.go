package dto

import (
	"github.com/gagliardetto/solana-go"
)

type CustomDomainRequest struct {
	DriveID   solana.PublicKey `json:"drive_id"` //solana.PublicKey
	Domain    string           `json:"domain"`
	Signature SignedMessage    `json:"signature"`
	Timestamp int64            `json:"timestamp"`
}
