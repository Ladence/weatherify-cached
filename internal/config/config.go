package config

import "github.com/Ladence/weatherify-cached/internal/gateway/weatherstack"

type Config struct {
	UseRedis     bool                `json:"use_redis"`
	Port         string              `json:"port"`
	Weatherstack weatherstack.Config `json:"weatherstack"`
}
