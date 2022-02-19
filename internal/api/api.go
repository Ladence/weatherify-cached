package api

import "github.com/Ladence/weatherify-cached/internal/domain"

type GetWeatherResponse struct {
	domain.Weather `json:"weather"`
}
