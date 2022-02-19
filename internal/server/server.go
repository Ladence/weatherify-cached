package server

import (
	ctx "context"
	"fmt"
	"github.com/Ladence/weatherify-cached/internal/gateway/weatherstack"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
	"github.com/sirupsen/logrus"
)

type Server struct {
	log                *logrus.Logger
	redisCache         *cache.Cache
	weatherstackClient *weatherstack.Client
}

func NewServer(log *logrus.Logger, useRedis bool, client *weatherstack.Client) (*Server, error) {
	if client == nil {
		err := fmt.Errorf("passed weatherifyClient is nil")
		return nil, err
	}
	s := &Server{
		log:                log,
		weatherstackClient: client,
	}

	if useRedis {
		s.redisCache = newRedisCache()
	}
	return s, nil
}

func (s *Server) Run(port string) error {
	r := gin.Default()
	r.GET("/weather", s.cacheMw, func(context *gin.Context) {
		s.log.Infof("Going to call WeatherStack API with city = %s", context.Query("city"))
		current, err := s.weatherstackClient.GetCurrent(ctx.TODO(), "892719c2f5b87c88ac01025c0c409dce", context.Query("city"))
		if err != nil {
			s.log.Errorf("Error when handling external request. Error :%v", err)
		}
		s.log.Infof("%+v", *current)
	})

	return r.Run("localhost:" + port)
}
