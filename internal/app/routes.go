package app

import (
	"context"
	"net/http"

	"ai-code-reviewer/internal/ai"
	"ai-code-reviewer/internal/dedup"
	"ai-code-reviewer/internal/github"
	"ai-code-reviewer/internal/observability"
	"ai-code-reviewer/internal/ratelimit"
	"ai-code-reviewer/internal/worker"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (s *Server) routes() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", s.health)
	// create core queue

	// create queue based on config
	queue := worker.NewQueue(s.cfg)

	// adapter so github pkg doesn't know worker
	adapter := worker.NewAdapter(queue)

	// github client
	ghClient := github.NewClient(s.cfg, s.logger)

	// webhook
	gh := github.NewWebhookHandler(
		s.cfg,
		s.logger,
		ghClient,
		adapter,
	)

	commenter := github.NewCommentService(
		s.cfg.GitHubToken,
	)

	primary := ai.NewProvider(s.cfg)

	fallback := ai.NewFallback(primary, ai.NewOpenAI(s.cfg.OpenAIKey, s.cfg.OpenAIModel))

	dedup := dedup.NewMemory()

	//ratelimiter
	rateLimiter := ratelimit.New(s.cfg.RateLimitRPS, s.cfg.RateLimitBurst)

	// background processor
	processor := worker.NewProcessor(
		queue,
		ghClient,
		commenter,
		dedup,
		s.logger,
		fallback,
		rateLimiter,
	)

	// init metrics
	observability.InitMetrics()

	mux.HandleFunc("/webhook/github", gh.Handle)
	mux.Handle("/metrics", promhttp.Handler())

	processor.Start(context.Background())

	s.http.Handler = mux
}
