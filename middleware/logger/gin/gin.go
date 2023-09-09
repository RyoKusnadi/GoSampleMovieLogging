package gin

import (
	"time"

	"log/slog"

	"scalable-go-movie/middleware/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GinRequestLogger struct {
	*logger.RequestLogger
}

func NewGinRequestLogger(rl *logger.RequestLogger) *GinRequestLogger {
	return &GinRequestLogger{
		RequestLogger: rl,
	}
}

func (grl *GinRequestLogger) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		requestID, _ := grl.generateRequestID(c)
		grl.setHeader(c, requestID)

		defer func() {
			end := time.Now()
			latency := end.Sub(start)

			level := grl.RequestLogger.DetermineHttpLogLevel(c.Writer.Status())
			attrs := grl.buildAttributes(c, path, requestID, latency, level)
			grl.RequestLogger.LogWithContext(c, level, "request:", attrs...)
		}()

		c.Next()
	}
}

func (rl *GinRequestLogger) generateRequestID(c *gin.Context) (string, error) {
	if rl.Config.WithRequestID {
		id, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		return id.String(), nil
	}
	return "", nil
}

func (rl *GinRequestLogger) setHeader(c *gin.Context, requestID string) {
	if rl.Config.WithRequestID {
		c.Header("X-Request-ID", requestID)
	}
}

func (rl *GinRequestLogger) buildAttributes(c *gin.Context, path, requestID string, latency time.Duration, level slog.Level) []slog.Attr {
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
