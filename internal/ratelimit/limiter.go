package ratelimit

import (
	"sync"

	"golang.org/x/time/rate"
)

type Limiter struct {
	mu       sync.Mutex
	Limiters map[string]*rate.Limiter
	Rps      rate.Limit
	Burst    int
}

func New(rps int, burst int) *Limiter {
	return &Limiter{
		Limiters: make(map[string]*rate.Limiter),
		Rps:      rate.Limit(rps),
		Burst:    burst,
	}
}

func (l *Limiter) Get(repo string) *rate.Limiter {
	l.mu.Lock()
	defer l.mu.Unlock()

	if limiter, ok := l.Limiters[repo]; ok {
		return limiter
	}

	limiter := rate.NewLimiter(l.Rps, l.Burst)
	l.Limiters[repo] = limiter
	return limiter
}
