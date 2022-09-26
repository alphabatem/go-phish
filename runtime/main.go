package main

import (
	"fmt"
	"github.com/alphabatem/go-phish/services"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Println("Runtime - Starting")

	err := godotenv.Load()
	if err != nil {
		log.Panic(fmt.Sprintf("%s", err))
		return
	}

	svcMap := map[string]services.Service{}
	svcs := []services.Service{
		&services.SolanaService{},
		&services.ShadowDriveService{},
		&services.BlockListService{},
		&services.TwitterService{},
		&services.HttpService{},
	}
	for _, service := range svcs {
		svcMap[service.Id()] = service
	}

	//Configure services
	for _, service := range svcs {
		log.Printf("Configuring: %s", service.Id())
		err := service.Configure(svcMap)
		if err != nil {
			log.Panic(fmt.Sprintf("%s", err))
			return
		}
	}
	defer shutdown(svcs)

	//Start service once configured
	for _, service := range svcs {
		log.Printf("Starting: %s", service.Id())
		err := service.Start()
		if err != nil {
			log.Panic(fmt.Sprintf("%s", err))
			return
		}
	}

	log.Println("Runtime - Complete")
}

func shutdown(svcs []services.Service) {
	log.Printf("Shutting down...")
	for _, service := range svcs {
		log.Printf("Shutting Down: %s", service.Id())
		service.Stop()
	}
}
