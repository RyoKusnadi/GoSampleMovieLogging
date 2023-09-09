package gin

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"

	"scalable-go-movie/middleware/logger"
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

		requestID, _ := grl.RequestLogger.GenerateRequestID()
		grl.setHeader(c, requestID)

		defer func() {
			end := time.Now()
			latency := end.Sub(start)

			level := grl.RequestLogger.DetermineHttpLogLevel(c.Writer.Status())
			attrs := grl.buildAttributes(c, path, requestID, latency, level)
			grl.LogWithContext(c, level, "request:", attrs...)
		}()

		c.Next()
	}
}

func (rl *GinRequestLogger) setHeader(c *gin.Context, requestID string) {
	if rl.Config.WithRequestID {
		c.Header("X-Request-ID", requestID)
	}
}

func (grl *GinRequestLogger) buildAttributes(c *gin.Context, path, requestID string, latency time.Duration, level slog.Level) []slog.Attr {
	attrs := []slog.Attr{
		slog.Int("status", c.Writer.Status()),
		slog.String("method", c.Request.Method),
		slog.String("path", path),
		slog.Duration("latency", latency),
		slog.String("user-agent", c.Request.UserAgent()),
	}

	if grl.Config.WithRequestID {
		attrs = append(attrs, slog.String("request-id", requestID))
	}

	return attrs
}

func (grl *GinRequestLogger) LogWithContext(c *gin.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	grl.RequestLogger.WriteLogIntoTxt(level, msg, attrs...)
	grl.Logger.LogAttrs(c, level, msg, attrs...)
}

func (rl *GinRequestLogger) GetRequestID(c *gin.Context) string {
	if rl.Config.WithRequestID {
		return c.GetHeader("X-Request-ID")
	}
	return ""
}
