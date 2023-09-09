package logger

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	StatusOK       = http.StatusOK
	StatusCreated  = http.StatusCreated
	StatusAccepted = http.StatusAccepted
	DefaultLogPath = "log.txt"
)

func (rl *RequestLogger) DetermineHttpLogLevel(status int) slog.Level {
	switch {
	case status == StatusOK ||
		status == StatusCreated ||
		status == StatusAccepted:
		return slog.LevelInfo
	default:
		return slog.LevelError
	}
}

func (rl *RequestLogger) LogWithContext(c *gin.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	if rl.Config.EnableWriteTxtLog {
		logFileDir, ok := rl.Config.LogFilePaths[level]
		if !ok || logFileDir == "" {
			logFileDir = DefaultLogPath
		}

		file, err := os.OpenFile(logFileDir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			defer file.Close()
			logLine := fmt.Sprintf("[%s] %s", level, msg)
			for _, attr := range attrs {
				logLine += " " + fmt.Sprint(attr)
			}
			logLine += "\n"

			file.WriteString(logLine)
		}
	}

	rl.Logger.LogAttrs(c, level, msg, attrs...)
}
