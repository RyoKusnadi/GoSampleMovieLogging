package logger

import (
	"log/slog"
	"sync"
)

type Config struct {
	WithRequestID     bool
	LogFilePaths      map[slog.Level]string
	EnableWriteTxtLog bool
	CustomLogLevels   []slog.Level
	GrayScale         *GrayScaleConfig
	mu                sync.Mutex
}

type GrayScaleConfig struct {
	Enabled          bool
	Threshold        int
	Percentage       int
	TotalRequests    int
	TotalRequestsMux sync.Mutex
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

func WithEnableWriteTxtLog() Option {
	return func(rl *RequestLogger) {
		rl.Config.EnableWriteTxtLog = true
	}
}

func WithLogFilePaths(logFilePaths map[slog.Level]string) Option {
	return func(rl *RequestLogger) {
		rl.Config.LogFilePaths = logFilePaths
	}
}

func WithCustomLevelLogs(customLevelLog []slog.Level) Option {
	return func(rl *RequestLogger) {
		rl.Config.CustomLogLevels = customLevelLog
	}
}

func WithGrayScale(grayScaleConfig *GrayScaleConfig) Option {
	return func(rl *RequestLogger) {
		rl.Config.GrayScale = grayScaleConfig
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
