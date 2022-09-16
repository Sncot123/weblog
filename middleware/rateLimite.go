package middleware

import (
	"net/http"
	"time"

	"github.com/juju/ratelimit"

	"github.com/gin-gonic/gin"
)

/*
	限流策略：1、漏桶策略   go get github.com/uber-go/ratelimit
			2、令牌桶策略  go get github.com/juju/ratelimit

*/

// RateLimitMiddleWare 令牌桶限流策略
func RateLimitMiddleWare(fileInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fileInterval, cap)
	return func(c *gin.Context) {
		//如果取不到令牌就中断本次请求返回 rate limit
		if bucket.TakeAvailable(1) < 1 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}
