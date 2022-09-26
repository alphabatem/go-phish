package services

import (
	"github.com/alphabatem/go-phish/providers"
	"github.com/gagliardetto/solana-go"
	"os"
	"time"
)

type SolanaService struct {
	rpcUrl       string
	pollInterval time.Duration
	programID    string

	sol *providers.SolanaProvider
}

type SignatureLookup struct {
	Signature solana.Signature
	BlokTime  time.Time
	Attempts  int
}

const SOLANA_SVC = "solana_svc"

func (SolanaService) Id() string {
	return SOLANA_SVC
}

func (svc *SolanaService) Configure(svcs map[string]Service) error {
	svc.rpcUrl = os.Getenv("RPC_ENDPOINT")
	svc.programID = os.Getenv("PROGRAM_ID")
	svc.pollInterval = 10 * time.Second

	return nil
}

func (svc *SolanaService) Start() error {
	svc.sol = providers.NewSolanaProvider(svc.rpcUrl)

	return nil
}

/**
 * Note: So we could do WS here (and can always swap) but seems pointless for cid registration when we can poll
 * 	     and consume less resources rather than having the open TCP socket - something to think on.
 */

// PollWs using websocket listener
// func (svc *SolanaService) PollWs(onInstruction func(acc rpc.KeyedAccount)) error {
func (svc *SolanaService) PollWs(onLogs func(signature *SignatureLookup, logs []string)) error {
	sub, err := svc.sol.LogSubscribe(svc.programID)
	if err != nil {
		return err
	}

	go svc.sol.ListenLogs(sub, func(signature solana.Signature, logs []string) {
		onLogs(&SignatureLookup{
			Signature: signature,
		}, nil)
	})

	return nil
}

func (svc *SolanaService) Stop() {
	//
}
