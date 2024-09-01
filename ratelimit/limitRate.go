package box

import (
	"net/http"
	"sync"
	"time"
)

type SimpleLimiter struct {
	mu          sync.Mutex
	lastRequest time.Time
	minInterval time.Duration
}
// time bitwen 
func NewLimiter(requestsPerSecond int) *SimpleLimiter {
	return &SimpleLimiter{
		minInterval: time.Second / time.Duration(requestsPerSecond),
	}
}

func (sl *SimpleLimiter) Allow() bool {
	sl.mu.Lock()
	defer sl.mu.Unlock()

	now := time.Now()
	if now.Sub(sl.lastRequest) < sl.minInterval {
		return false
	}

	sl.lastRequest = now
	return true
}

func RateLimitMiddleware(next http.HandlerFunc, limiter *SimpleLimiter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			return 
		}
		next.ServeHTTP(w, r)
	}
}
