package server

import (
	ctx "context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"

	"github.com/Ladence/weatherify-cached/internal/domain"
)

func newRedisCache() *cache.Cache {
	r := redis.NewRing(&redis.RingOptions{})
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

	var wanted *domain.Weather
	cacheKey := context.Query("city")
	err := s.redisCache.Get(ctx.Background(), cacheKey, wanted)
	if err != nil {
		s.log.Errorf("Error on redisCache.Get. %v", err)
		s.redisCache.Set(&cache.Item{
			Ctx:   ctx.Background(),
			Key:   cacheKey,
			Value: nil,
			TTL:   time.Hour,
		})
	}
}
