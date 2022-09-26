package services

type Service interface {
	Id() string
	Configure(map[string]Service) error
	Start() error
	Stop()
}
