package ratelimit

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type Limiter struct {
	mu       sync.Mutex
	visitors map[string]*visitor
	r        rate.Limit
	burst    int
}

func New(r rate.Limit, burst int) *Limiter {
	l := &Limiter{
		visitors: make(map[string]*visitor),
		r:        r,
		burst:    burst,
	}

	go l.cleanupLoop()

	return l
}

func (l *Limiter) allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	v, ok := l.visitors[key]
	if !ok {
		v = &visitor{limiter: rate.NewLimiter(l.r, l.burst)}
		l.visitors[key] = v
	}

	v.lastSeen = time.Now()

	return v.limiter.Allow()
}

func (l *Limiter) cleanupLoop() {
	for range time.Tick(5 * time.Minute) {
		l.mu.Lock()

		for key, v := range l.visitors {
			if time.Since(v.lastSeen) > 10*time.Minute {
				delete(l.visitors, key)
			}
		}

		l.mu.Unlock()
	}
}

func (l *Limiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !l.allow(c.ClientIP()) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests, please slow down",
			})
			return
		}

		c.Next()
	}
}
