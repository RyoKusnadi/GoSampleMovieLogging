package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"

	"scalable-go-movie/adapter"
	"scalable-go-movie/config"
	"scalable-go-movie/infrastructure"
	"scalable-go-movie/middleware/logger"
	ginRequestRecorder "scalable-go-movie/middleware/logger/gin"
	"scalable-go-movie/usecase"
)

func main() {
	sysConfig := config.Get()

	headers := map[string]string{
		"Authorization": "Bearer " + sysConfig.MovieApi.Token,
		"accept":        "application/json",
	}
	httpClient := infrastructure.NewHTTPClient()
	httpClient.SetRequestHeaders(headers)
	movieUsecase := usecase.NewMovieUsecase(httpClient)
	httpAdapter := adapter.NewHTTPAdapter(movieUsecase)

	router := gin.New()

	logOptions := make([]logger.Option, 0)
	loggerConfig := config.Get().Logger
	if !loggerConfig.WithRequestID {
		logOptions = append(logOptions, logger.WithRequestID())
	}
	if loggerConfig.LogFilePath != nil {
		logFilePath := make(map[slog.Level]string)
		for level, path := range loggerConfig.LogFilePath {
			slogLevel := slog.Level(level)
			logFilePath[slogLevel] = path
		}
		logOptions = append(logOptions, logger.WithLogFilePaths(logFilePath))
	}
	if loggerConfig.EnableWriteTxtLog != false {
		logOptions = append(logOptions, logger.WithEnableWriteTxtLog())
	}
	if loggerConfig.CustomLogLevels != nil {
		customLevelLogs := make([]slog.Level, 0)
		for _, customLogLevel := range loggerConfig.CustomLogLevels {
			slogLevel := slog.Level(customLogLevel)
			customLevelLogs = append(customLevelLogs, slogLevel)
		}
		logOptions = append(logOptions, logger.WithCustomLevelLogs(customLevelLogs))
	}
	if loggerConfig.GrayScale != nil {
		grayScaleConfig := logger.GrayScaleConfig{
			Enabled:       loggerConfig.GrayScale.Enabled,
			Threshold:     loggerConfig.GrayScale.Threshold,
			Percentage:    loggerConfig.GrayScale.Percentage,
			TotalRequests: loggerConfig.GrayScale.TotalRequests,
		}
		logOptions = append(logOptions, logger.WithGrayScale(&grayScaleConfig))
	}

	requestLogger, err := logger.NewRequestLogger(
		slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		logOptions...,
	)
	if err != nil {
		panic(err)
	}

	ginRequestRecorderMiddleware := ginRequestRecorder.NewGinRequestLogger(requestLogger).Middleware()
	router.Use(ginRequestRecorderMiddleware)

	httpAdapter.RegisterRoutes(router)
	router.Run(":8080")
}
