package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"ai-code-reviewer/internal/app"
	"ai-code-reviewer/internal/config"
	"ai-code-reviewer/internal/observability"
)

func main() {
	cfg := config.Load()

	logger := observability.NewLogger(cfg)

	srv := app.NewServer(cfg, logger)

	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := srv.Start(ctx); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
