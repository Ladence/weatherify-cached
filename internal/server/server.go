package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
	"github.com/sirupsen/logrus"
)

type Server struct {
	log        *logrus.Logger
	redisCache *cache.Cache
}

func NewServer(log *logrus.Logger, useRedis bool) *Server {
	s := &Server{
		log: log,
	}

	if useRedis {
		s.redisCache = newRedisCache()
	}
	return s
}

func (s *Server) Run(port string) error {
	r := gin.Default()
	r.Use(s.cacheMw)

	r.GET("/weather", s.cacheMw, func(context *gin.Context) {

	})

	return r.Run("localhost:" + port)
}
