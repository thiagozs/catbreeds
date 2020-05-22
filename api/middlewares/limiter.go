package middlewares

import (
	"catbreeds/api/models"
	"fmt"
	"net/http"
	"time"

	"github.com/codebear4/ttlcache"
	"github.com/gin-gonic/gin"
)

// Limiter monitoring login call url
func Limiter() gin.HandlerFunc {
	cache := ttlcache.NewCache()
	cache.SetExpirationCallback(func(key string, value interface{}) {
		fmt.Printf("This session keys=%s has restored\n", key)
		cache.Remove(key)
	})
	return func(c *gin.Context) {

		user, ok := c.Get("uuid")
		if !ok {
			c.Next()
		}
		uuid := user.(*models.Login).UserName
		hash := fmt.Sprintf("%s@%s", c.Request.URL, uuid)

		vl, ok := cache.Get(hash)
		if !ok {
			cache.SetWithTTL(hash, 1, time.Duration(120)*time.Second)
		} else {
			val := vl.(int)
			if val > 3 {
				c.AbortWithStatusJSON(http.StatusUnauthorized,
					map[string]string{
						"code":    "401",
						"message": "session limit reached",
					})

				return
			}
			val += 1
			cache.SetWithTTL(hash, val, time.Duration(60)*time.Second)
		}
		c.Next()
	}
}
