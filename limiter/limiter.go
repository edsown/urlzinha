package limiter

import (
	"net/http"
	"sync"
	"time"

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
	b        int
}

func NewLimiter(r rate.Limit, b int) *Limiter {
	rateLimiter := &Limiter{
		visitors: make(map[string]*visitor),
		r:        r,
		b:        b,
	}

	go rateLimiter.cleanupVisitors()
	return rateLimiter
}

func (l *Limiter) getLimiter(key string) *rate.Limiter {
	l.mu.Lock()
	defer l.mu.Unlock()

	v, exists := l.visitors[key]
	if !exists {
		limiter := rate.NewLimiter(l.r, l.b)
		return limiter
	}
	v.lastSeen = time.Now()
	return v.limiter
}

func (l *Limiter) cleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		l.mu.Lock()
		for key, v := range l.visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(l.visitors, key)
			}
		}
		l.mu.Unlock()
	}
}

func (l *Limiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.RemoteAddr
		limiter := l.getLimiter(key)

		if !limiter.Allow() {
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
