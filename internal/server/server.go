package server

import (
	ctx "context"
	"fmt"
	"github.com/Ladence/weatherify-cached/internal/config"
	"github.com/Ladence/weatherify-cached/internal/conv"
	"github.com/Ladence/weatherify-cached/internal/gateway/weatherstack"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
	"github.com/sirupsen/logrus"
	"time"
)

type Server struct {
	config             *config.Config
	log                *logrus.Logger
	redisCache         *cache.Cache
	weatherstackClient *weatherstack.Client
}

func NewServer(log *logrus.Logger, config *config.Config, client *weatherstack.Client) (*Server, error) {
	if client == nil {
		err := fmt.Errorf("passed weatherifyClient is nil")
		return nil, err
	}
	s := &Server{
		log:                log,
		config:             config,
		weatherstackClient: client,
	}

	if s.config.UseRedis {
		s.redisCache = newRedisCache()
	}
	return s, nil
}

func (s *Server) Run() error {
	r := gin.Default()
	r.GET("/weather", s.cacheMw, func(context *gin.Context) {
		s.log.Infof("Going to call WeatherStack API with city = %s", context.Query("city"))
		current, err := s.weatherstackClient.GetCurrent(ctx.TODO(), s.config.Weatherstack.AccessKey, context.Query("city"))
		if err != nil {
			s.log.Errorf("Error when handling external request. Error :%v", err)
			context.Status(501)
			return
		}
		s.log.Infof("Current: %+v", *current)
		weather, err := conv.CurrentToWeather(&current.Current)
		if err != nil {
			s.log.Errorf("Error on conversion weatherstack.current -> model.weather. Error: %v", err)
			context.Status(501)
			return
		}
		context.JSON(200, *weather)
		if s.redisCache != nil {
			s.log.Debugf("Placing to Redis cache.")
			if err := s.redisCache.Set(&cache.Item{
				Ctx:   ctx.Background(),
				Key:   context.Query("city"),
				Value: nil,
				TTL:   time.Hour,
			}); err != nil {
				s.log.Debugf("Failed to place in Redis cache. Error: %v", err)
			}
		}
	})

	return r.Run("localhost:" + s.config.Port)
}
