package services

type BlockListService struct {
	Service
}

const BLOCKLIST_SVC = "blocklist_svc"

func (svc BlockListService) Id() string {
	return BLOCKLIST_SVC
}
