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

	requestLogger, err := logger.NewRequestLogger(
		slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		logger.WithRequestID(),
	)
	if err != nil {
		panic(err)
	}

	ginRequestRecorderMiddleware := ginRequestRecorder.NewGinRequestLogger(requestLogger).Middleware()
	router.Use(ginRequestRecorderMiddleware)

	httpAdapter.RegisterRoutes(router)
	router.Run(":8080")
}
