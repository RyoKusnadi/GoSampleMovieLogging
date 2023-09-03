package middleware

import (
	"net/http"
	"time"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	StatusOK       = http.StatusOK
	StatusCreated  = http.StatusCreated
	StatusAccepted = http.StatusAccepted
)

type Config struct {
	WithRequestID bool
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

func NewRequestLogger(logger *slog.Logger, options ...Option) *RequestLogger {
	rl := &RequestLogger{
		Logger: logger,
	}

	for _, option := range options {
		option(rl)
	}

	return rl
}

func (rl *RequestLogger) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		requestID, _ := rl.generateRequestID(c)
		rl.setHeader(c, requestID)

		defer func() {
			end := time.Now()
			latency := end.Sub(start)

			level := rl.determineLogLevel(c.Writer.Status())
			attrs := rl.buildAttributes(c, path, requestID, latency, level)
			rl.logWithContext(c, level, "request:", attrs...)
		}()

		c.Next()
	}
}

func (rl *RequestLogger) generateRequestID(c *gin.Context) (string, error) {
	if rl.Config.WithRequestID {
		id, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		return id.String(), nil
	}
	return "", nil
}

func (rl *RequestLogger) setHeader(c *gin.Context, requestID string) {
	if rl.Config.WithRequestID {
		c.Header("X-Request-ID", requestID)
	}
}

func (rl *RequestLogger) buildAttributes(c *gin.Context, path, requestID string, latency time.Duration, level slog.Level) []slog.Attr {
	attrs := []slog.Attr{
		slog.Int("status", c.Writer.Status()),
		slog.String("method", c.Request.Method),
		slog.String("path", path),
		slog.Duration("latency", latency),
		slog.String("user-agent", c.Request.UserAgent()),
	}

	if rl.Config.WithRequestID {
		attrs = append(attrs, slog.String("request-id", requestID))
	}

	return attrs
}

func (rl *RequestLogger) determineLogLevel(status int) slog.Level {
	switch {
	case status == StatusOK ||
		status == StatusCreated ||
		status == StatusAccepted:
		return slog.LevelInfo
	default:
		return slog.LevelError
	}
}

func (rl *RequestLogger) logWithContext(c *gin.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	rl.Logger.LogAttrs(c, level, msg, attrs...)
}
