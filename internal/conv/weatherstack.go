package conv

import (
	"github.com/Ladence/weatherify-cached/internal/domain"
	"github.com/Ladence/weatherify-cached/internal/gateway/weatherstack"
)

func CurrentToWeather(current *weatherstack.Current) (*domain.Weather, error) {
	w := &domain.Weather{}
	if len(current.WeatherDescriptions) > 0 {
		w.Description = &current.WeatherDescriptions[0] // todo: just for example for now
	}
	w.Temperature = current.Temperature
	return w, nil
}
