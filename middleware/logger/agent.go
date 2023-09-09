package logger

import "log/slog"

const (
	LevelSamplingDebug = slog.Level(-8)
)

func (rl *RequestLogger) Log(level slog.Level, msg string, attrs ...slog.Attr) {
	if level != LevelSamplingDebug {
		rl.WriteLogIntoTxt(level, msg, attrs...)
	} else if rl.shouldLog(level) {
		rl.WriteLogIntoTxt(level, msg, attrs...)
	}
}

func (rl *RequestLogger) shouldLog(level slog.Level) bool {
	for _, customLevel := range rl.Config.CustomLogLevels {
		if level == customLevel {
			if rl.Config.Grayscale.Enabled {
				if rl.isGrayScaleRequest() {
					return true
				}
				return false
			}
			return true
		}
	}
	return false
}

func (rl *RequestLogger) isGrayScaleRequest() bool {
	rl.Config.Grayscale.TotalRequestsMux.Lock()
	defer rl.Config.Grayscale.TotalRequestsMux.Unlock()

	totalRequests := rl.Config.Grayscale.TotalRequests
	grayscaleThreshold := rl.Config.Grayscale.Threshold
	grayscalePercentage := rl.Config.Grayscale.Percentage

	if totalRequests > 0 && grayscaleThreshold > 0 {
		percentage := (totalRequests * 100) / grayscaleThreshold
		return percentage <= grayscalePercentage
	}
	return false
}
