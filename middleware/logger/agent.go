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
			if rl.Config.GrayScale.Enabled {
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
	rl.Config.GrayScale.TotalRequestsMux.Lock()
	defer rl.Config.GrayScale.TotalRequestsMux.Unlock()

	totalRequests := rl.Config.GrayScale.TotalRequests
	grayScaleThreshold := rl.Config.GrayScale.Threshold
	grayScalePercentage := rl.Config.GrayScale.Percentage

	if totalRequests > 0 && grayScaleThreshold > 0 {
		percentage := (totalRequests * 100) / grayScaleThreshold
		return percentage <= grayScalePercentage
	}
	return false
}
