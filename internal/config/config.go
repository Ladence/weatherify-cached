package config

import "github.com/Ladence/weatherify-cached/internal/gateway/weatherstack"

type Config struct {
	Redis        *Redis              `json:"redis"`
	Port         string              `json:"port"`
	Weatherstack weatherstack.Config `json:"weatherstack"`
}

type Redis struct {
	Address string `json:"address"`
}
