package main

import (
	"encoding/json"
	"flag"
	"github.com/Ladence/weatherify-cached/internal/config"
	"github.com/Ladence/weatherify-cached/internal/gateway/weatherstack"
	"github.com/Ladence/weatherify-cached/internal/server"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var cfgPath string

func initOpts() {
	flag.StringVar(&cfgPath, "config", "config.json", "path to a config file")
	flag.Parse()
}

func readConfig() (*config.Config, error) {
	bytes, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}
	cfg := &config.Config{}
	if err = json.Unmarshal(bytes, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func main() {
	log := logrus.New()
	weatherstackClient := weatherstack.NewClient(log)
	initOpts()
	cfg, err := readConfig()
	if err != nil {
		log.Fatalf("Error on reading config. Error: %v", err)
	}
	s, err := server.NewServer(log, cfg, weatherstackClient)
	if err != nil {
		log.Fatalf("Couldn't create server! Error: %v", err)
	}
	err = s.Run()
	if err != nil {
		log.Fatal("Error on running server")
	}
}
