package ai

import "context"

type ReviewRequest struct {
	File    string
	Content string
}

type Provider interface {
	Review(ctx context.Context, r ReviewRequest) (string, error)
}
