package main

import (
	"fmt"
	"log"
	config "tlkm-api/configs"
)

func main() {
	log.Println("starting: telkom-api via http.. ")

	err := config.Init(
		config.WithConfigFile("config"),
		config.WithConfigType("yaml"),
	)
	if err != nil {
		log.Panicf("failed to initialize config: %v", err)
	}

	fmt.Println("HELLO WORLD")

	srv := NewHTTP()
	NewDelivery(srv)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
		return
	}
}
