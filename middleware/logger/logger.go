package logger

import (
	"log/slog"
	"sync"
)

type Config struct {
	WithRequestID     bool
	LogFilePaths      map[slog.Level]string
	EnableWriteTxtLog bool
	mu                sync.Mutex
}

type RequestLogger struct {
	Logger *slog.Logger
	Config Config
}

type Option func(*RequestLogger)

func WithRequestID() Option {
	return func(rl *RequestLogger) {
		rl.Config.WithRequestID = true
	}
}

func WithEnableWriteTxtLog(enable bool) Option {
	return func(rl *RequestLogger) {
		rl.Config.EnableWriteTxtLog = enable
	}
}

func WithLogFilePaths(logFilePaths map[slog.Level]string) Option {
	return func(rl *RequestLogger) {
		rl.Config.LogFilePaths = logFilePaths
	}
}

func NewRequestLogger(logger *slog.Logger, options ...Option) (*RequestLogger, error) {
	rl := &RequestLogger{
		Logger: logger,
	}

	for _, option := range options {
		option(rl)
	}

	return rl, nil
}
