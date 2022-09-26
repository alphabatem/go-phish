package providers

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSolanaProvider_AccountData(t *testing.T) {
	svc := SolanaProvider{}

	resp, err := svc.AccountData("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
		return
	}

	out, _ := json.Marshal(resp)
	log.Printf("Out: %s", out)
}

func TestSolanaProvider_AccountInfo(t *testing.T) {

	//pk := "2wci94quHBAAVt1HC4T5SUerZR7699LMb8Ueh3CSVpTX"
	pk := "HXzevyLkgtNSUQVQKZHUeRyDkUCC2KaurDqMrdNe6Ude" //Sol sneks nft

	svc := SolanaProvider{}

	resp, err := svc.AccountInfo(pk)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
		return
	}

	out, _ := json.Marshal(resp)
	log.Printf("Out: %s", out)
}

func TestSolanaProvider_OwnedNFTTokens(t *testing.T) {
	pk := "2wci94quHBAAVt1HC4T5SUerZR7699LMb8Ueh3CSVpTX"
	program := "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"

	svc := SolanaProvider{}

	resp, err := svc.OwnedNFTTokens(program, pk)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
		return
	}

	log.Printf("%v", resp)

}

func TestSolanaProvider_OwnedNFTTokensCollection(t *testing.T) {
	pk := "2wci94quHBAAVt1HC4T5SUerZR7699LMb8Ueh3CSVpTX"
	program := "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"

	svc := SolanaProvider{}

	resp, err := svc.OwnedNFTTokens(program, pk)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
		return
	}

	log.Printf("%v", resp)

}

func TestSolanaProvider_Owner(t *testing.T) {
	pk := "CKVLo9GvmngP5rP1SX6d7t3tfBPeFa9eBpt3tM2hLiNe" //Sol sneks nft

	svc := SolanaProvider{}

	resp, err := svc.Owner(pk)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
		return
	}

	log.Printf("%v", resp)
}
