package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alphabatem/go-phish/generated/shadow_drive_user_staking"
	"github.com/gagliardetto/solana-go"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type ListenerFunc func(instruction *shadow_drive_user_staking.Instruction)

type ShadowDriveService struct {
	ssvc      *SolanaService
	programID solana.PublicKey

	httpClient *http.Client
}

type FileRequest struct {
	StorageAccount string `json:"storageAccount"`
}

type FileResponse struct {
	Keys []string `json:"keys"`
}

const SHADOW_DRIVE_SVC = "shadow_drive_svc"

func (ShadowDriveService) Id() string {
	return SHADOW_DRIVE_SVC
}

func (svc *ShadowDriveService) Configure(svcs map[string]Service) error {
	svc.ssvc = svcs[SOLANA_SVC].(*SolanaService)

	svc.httpClient = &http.Client{
		Timeout: 5 * time.Second,
	}

	pk := os.Getenv("PROGRAM_ID")

	var err error
	svc.programID, err = solana.PublicKeyFromBase58(pk)
	if err != nil {
		return err
	}

	svc.configureProgram()

	return nil
}

func (svc *ShadowDriveService) Start() error {
	return nil
}

func (svc *ShadowDriveService) configureProgram() {
	shadow_drive_user_staking.SetProgramID(svc.programID)
}

func (svc *ShadowDriveService) CheckFile(driveID solana.PublicKey) bool {
	resp, err := svc.httpClient.Head(fmt.Sprintf("https://shdw-drive.genesysgo.net/%s/index.html", driveID))
	if err != nil {
		return true //Just assume if errored
	}

	return resp.StatusCode == 200
}

func (svc *ShadowDriveService) Files(driveID solana.PublicKey) ([]string, error) {
	j, err := json.Marshal(&FileRequest{StorageAccount: driveID.String()})
	if err != nil {
		return nil, err
	}

	resp, err := svc.httpClient.Post("https://shadow-storage.genesysgo.net/list-objects", "application/json", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var fileList FileResponse
	err = json.Unmarshal(data, &fileList)
	if err != nil {
		return nil, err
	}

	return fileList.Keys, nil
}

func (svc *ShadowDriveService) Stop() {
	//
}
