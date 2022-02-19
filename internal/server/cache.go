package server

import (
	ctx "context"
	"github.com/Ladence/weatherify-cached/internal/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"

	"github.com/Ladence/weatherify-cached/internal/domain"
)

func newRedisCache(config *config.Redis) *cache.Cache {
	r := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": config.Address,
		},
	})
	return cache.New(&cache.Options{
		Redis:      r,
		LocalCache: cache.NewTinyLFU(10, time.Minute),
	})
}

func (s *Server) cacheMw(context *gin.Context) {
	if s.redisCache == nil {
		s.log.Infof("Redis caching is turned off. Going to true request")
		return
	}

	wanted := &domain.Weather{}
	cacheKey := context.Query("city")
	err := s.redisCache.Get(ctx.Background(), cacheKey, wanted)
	if err != nil {
		s.log.Debugf("Error on redisCache.Get. %v", err)
		return
	}
	s.log.Infof("Pulled from Redis cache. Weather: %+v", wanted)
	context.JSON(200, *wanted)
	context.Abort()
}
