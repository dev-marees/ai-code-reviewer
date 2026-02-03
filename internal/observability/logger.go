package observability

import (
	"log/slog"
	"os"

	"ai-code-reviewer/internal/config"
)

type Logger struct {
	*slog.Logger
}

func NewLogger(cfg *config.Config) *Logger {
	level := slog.LevelDebug
	if cfg.LogLevel == "info" {
		level = slog.LevelInfo
	}

	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	return &Logger{
		Logger: slog.New(h),
	}
}

func (l *Logger) Error(msg string, args ...any) {
	l.Logger.Error(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.Logger.Debug(msg, args...)
}
