package main

import (
	"github.com/Ladence/weatherify-cached/internal/gateway/weatherstack"
	"github.com/Ladence/weatherify-cached/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	weatherstackClient := weatherstack.NewClient(log)
	s, err := server.NewServer(log, false, weatherstackClient)
	if err != nil {
		log.Fatal("Couldn't create server! Error: %v", err)
	}
	err = s.Run("12337")
	if err != nil {
		log.Fatal("Error on running server")
	}
}
