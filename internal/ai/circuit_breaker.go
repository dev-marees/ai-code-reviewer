package ai

import (
	"context"

	"github.com/sony/gobreaker"
)

type CircuitBreakerProvider struct {
	provider Provider
	cb       *gobreaker.CircuitBreaker
}

func NewCircuitBreaker(p Provider) *CircuitBreakerProvider {

	settings := gobreaker.Settings{
		Name:        "ai-provider",
		MaxRequests: 3,
		Interval:    0,
		Timeout:     30,
	}

	return &CircuitBreakerProvider{
		provider: p,
		cb:       gobreaker.NewCircuitBreaker(settings),
	}
}

func (c *CircuitBreakerProvider) Review(
	ctx context.Context,
	r ReviewRequest,
) (string, error) {

	out, err := c.cb.Execute(func() (interface{}, error) {
		return c.provider.Review(ctx, r)
	})

	if err != nil {
		return "", err
	}

	return out.(string), nil
}
