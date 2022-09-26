package providers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"log"
	"net/url"
	"sync"
	"time"
)

type SolanaProvider struct {
	connection *rpc.Client
	client     *ws.Client
	rpcUrl     string

	txLock  sync.Locker
	txCache map[string]bool
}

func NewSolanaProvider(rpcUrl string) *SolanaProvider {
	return &SolanaProvider{
		rpcUrl:  rpcUrl,
		txLock:  &sync.Mutex{},
		txCache: make(map[string]bool),
	}
}

func (svc *SolanaProvider) Verify(publicKey string, nonce string, signedMessage []byte) error {

	sig := solana.SignatureFromBytes(signedMessage)
	pk, err := solana.PublicKeyFromBase58(publicKey)
	if err != nil {
		log.Printf("PublicKeyFromBase58::Error: %s", err)
		return err
	}

	ok := sig.Verify(pk, svc.Message(publicKey, nonce))
	if !ok {
		log.Printf("Message::Error: %s", err)
		return errors.New("invalid signature")
	}

	return nil
}

func (svc *SolanaProvider) Message(walletAddr string, nonce string) []byte {
	str := fmt.Sprintf("Blok Host Nonce: %s:%s", walletAddr, nonce)
	return []byte(str)
}

func (svc *SolanaProvider) AccountData(mintAddr string) (*token.Mint, error) {
	var m token.Mint

	addr := solana.MustPublicKeyFromBase58(mintAddr)

	ctx, cancel := svc.DefaultTimeoutCtx()
	defer cancel()

	err := svc.Connection().GetAccountDataInto(ctx, addr, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (svc *SolanaProvider) AccountInfo(accountAddr string) (*rpc.GetAccountInfoResult, error) {
	addr := solana.MustPublicKeyFromBase58(accountAddr)

	ctx, cancel := svc.DefaultTimeoutCtx()
	defer cancel()

	acct, err := svc.Connection().GetAccountInfo(ctx, addr)
	if err != nil {
		return nil, err
	}

	return acct, nil
}

func (svc *SolanaProvider) OwnedNFTTokens(mintAddr string, ownerAddr string) ([]string, error) {
	owner, err := solana.PublicKeyFromBase58(ownerAddr)
	if err != nil {
		return nil, err
	}

	mint, err := solana.PublicKeyFromBase58(mintAddr)
	if err != nil {
		return nil, err
	}

	resp, err := svc.GetTokenAccountsByOwner(owner, rpc.CommitmentFinalized, mint)
	if err != nil {
		return nil, err
	}

	//log.Printf("OwnedTokens::resp: %+v", resp)

	nftMints := []string{}
	for _, tok := range resp.Value {

		var ta token.Account
		_ = json.Unmarshal(tok.Account.Data.GetBinary(), &ta)

		if ta.Amount != 1 {
			continue
		}

		//log.Printf("OwnedTokens::mint: %+v", token.Account.Data.Parsed.Info.Mint)
		nftMints = append(nftMints, ta.Mint.String())
	}

	return nftMints, nil
}

func (svc *SolanaProvider) Owner(mintAddr string) (string, error) {

	mint, err := solana.PublicKeyFromBase58(mintAddr)
	if err != nil {
		return "", err
	}

	resp, err := svc.GetTokenLargestAccounts(mint)
	if err != nil {
		return "", err
	}

	if len(resp.Value) == 0 {
		return "", errors.New("no owner found")
	}

	return resp.Value[0].Address.String(), nil
}

func (svc *SolanaProvider) Connection() *rpc.Client {
	if svc.connection != nil {
		return svc.connection
	}

	uri := svc.ConnectionUri(false)
	log.Printf("Connecting to rpc: %s", uri)

	svc.connection = rpc.New(uri)
	return svc.connection
}

func (svc *SolanaProvider) ConnectionUri(ws bool) string {
	uri := svc.rpcUrl
	if uri == "" {
		uri = "https://ssc-dao.genesysgo.net"
	}
	u, _ := url.Parse(svc.rpcUrl)

	if ws {
		u.Scheme = "wss"
		return u.String()
	}

	u.Scheme = "https"
	return u.String()
}

type GetTokenAccountsResult struct {
	rpc.RPCContext
	Value []rpc.KeyedAccount
}

func (svc *SolanaProvider) DefaultTimeoutCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*5)
}

func (svc *SolanaProvider) GetTokenAccountsByOwner(publicKey solana.PublicKey, commitment rpc.CommitmentType, programID solana.PublicKey) (*rpc.GetTokenAccountsResult, error) {
	ctx, cancel := svc.DefaultTimeoutCtx()
	defer cancel()

	out, err := svc.Connection().GetTokenAccountsByOwner(ctx, publicKey, &rpc.GetTokenAccountsConfig{
		ProgramId: &programID,
	}, &rpc.GetTokenAccountsOpts{
		Commitment: commitment,
		Encoding:   "jsonParsed",
	})

	return out, err
}

func (svc *SolanaProvider) GetTokenLargestAccounts(publicKey solana.PublicKey) (out *rpc.GetTokenLargestAccountsResult, err error) {
	return svc.getTokenLargestAccounts(publicKey, rpc.CommitmentFinalized)
}

func (svc *SolanaProvider) getTokenLargestAccounts(publicKey solana.PublicKey, commitment rpc.CommitmentType) (*rpc.GetTokenLargestAccountsResult, error) {
	ctx, cancel := svc.DefaultTimeoutCtx()
	defer cancel()

	out, err := svc.Connection().GetTokenLargestAccounts(ctx, publicKey, commitment)

	return out, err
}

func (svc SolanaProvider) ProgramSubscribe(programId string) (*ws.ProgramSubscription, error) {
	var err error

	uri := svc.ConnectionUri(true)
	log.Printf("Connecting to ws: %s", uri)
	svc.client, err = ws.Connect(context.Background(), uri)
	if err != nil {
		return nil, err
	}

	program := solana.MustPublicKeyFromBase58(programId)
	sub, err := svc.client.ProgramSubscribeWithOpts(
		program,
		rpc.CommitmentFinalized,
		solana.EncodingBase64Zstd,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (svc *SolanaProvider) ListenProgram(sub *ws.ProgramSubscription, parseInput func(acc rpc.KeyedAccount)) {
	defer sub.Unsubscribe()
	for {
		got, err := sub.Recv()
		if err != nil {
			log.Printf("MarketplaceWatch:Recv::Error: %v", err)
			continue
		}
		//decodedBinary := got.Value.Account.Data.GetBinary()
		//if decodedBinary != nil {
		//	spew.Dump(decodedBinary)
		//}

		parseInput(got.Value)

		//log.Printf("%s", got.Value.Account.Data.GetBinary())
		//spew.Dump(got)
		//parseLogs(logs)
	}
}

func (svc SolanaProvider) LogSubscribe(programId string) (*ws.LogSubscription, error) {
	var err error

	uri := svc.ConnectionUri(true)
	log.Printf("Connecting to ws: %s", uri)
	svc.client, err = ws.Connect(context.Background(), uri)
	if err != nil {
		return nil, err
	}

	program := solana.MustPublicKeyFromBase58(programId)
	sub, err := svc.client.LogsSubscribeMentions(
		program,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (svc *SolanaProvider) Transaction(signature solana.Signature) (*rpc.GetTransactionResult, error) {
	return svc.Connection().GetTransaction(context.Background(), signature, &rpc.GetTransactionOpts{
		Encoding:   solana.EncodingBase64,
		Commitment: rpc.CommitmentConfirmed,
	})
}

func (svc *SolanaProvider) ListenLogs(sub *ws.LogSubscription, parseLogs func(solana.Signature, []string)) {
	defer sub.Unsubscribe()
	for {
		got, err := sub.Recv()
		if err != nil {
			log.Printf("SolanaProvider:Recv::Error: %v", err)
			continue
		}

		if got.Value.Err != nil {
			log.Printf("Error: %v", got.Value.Err)
			continue
		}

		//Dont process duplicate txns
		sig := got.Value.Signature
		if svc.AlreadyProcessed(sig.String()) {
			log.Printf("Already Processed: %s", sig)
			continue
		}

		//log.Printf("Market: %s - Log: %s", marketName, got.Value.Signature)

		//Flip instructions so we get the last instruction (rather than the first)
		logs := got.Value.Logs
		for i, j := 0, len(logs)-1; i < j; i, j = i+1, j-1 {
			logs[i], logs[j] = logs[j], logs[i]
		}

		log.Printf("%+v\n", got)
		//spew.Dump(got)
		parseLogs(sig, logs)
	}
}

func (svc *SolanaProvider) AlreadyProcessed(sig string) bool {
	//Mutex txLock
	svc.txLock.Lock()
	defer svc.txLock.Unlock()

	if svc.txCache[sig] {
		//log.Printf("Skipping duplicate txn: %s", signature)
		return true
	}

	if len(svc.txCache) > 1000 { //Make sure we dont use up too much ram with the cache
		svc.txCache = make(map[string]bool)
	}
	svc.txCache[sig] = true
	return false
}
