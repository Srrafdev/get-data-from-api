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
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return 
		}
		next.ServeHTTP(w, r)
	}
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// type RateLimiter struct {
// 	rate       int
// 	interval   time.Duration
// 	mutex      sync.Mutex
// 	requests   int
// 	lastReset  time.Time
// }

// func NewRateLimiter(rate int, interval time.Duration) *RateLimiter {
// 	return &RateLimiter{
// 		rate:      rate,
// 		interval:  interval,
// 		lastReset: time.Now(),
// 	}
// }

// func (rl *RateLimiter) Allow() bool {
// 	rl.mutex.Lock()
// 	defer rl.mutex.Unlock()

// 	now := time.Now()
// 	if now.Sub(rl.lastReset) > rl.interval {
// 		rl.requests = 0
// 		rl.lastReset = now
// 	}

// 	if rl.requests >= rl.rate {
// 		return false
// 	}

// 	rl.requests++
// 	return true
// }

// func rateLimitMiddleware(next http.HandlerFunc, rl *RateLimiter) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if !rl.Allow() {
// 			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	}
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, World!")
// }

// func main() {
// 	limiter := NewRateLimiter(10, time.Minute)
// 	http.HandleFunc("/", rateLimitMiddleware(helloHandler, limiter))
// 	http.ListenAndServe(":8080", nil)
// }
