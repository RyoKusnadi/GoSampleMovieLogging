package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"scalable-go-movie/adapter"
	"scalable-go-movie/infrastructure"
	"scalable-go-movie/usecase"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	apiKey := os.Getenv("API_KEY")

	headers := map[string]string{
		"Authorization": "Bearer " + apiKey,
		"accept":        "application/json",
	}
	httpClient := infrastructure.NewHTTPClient()
	httpClient.SetRequestHeaders(headers)
	movieUsecase := usecase.NewMovieUsecase(httpClient)
	httpAdapter := adapter.NewHTTPAdapter(movieUsecase)

	router := gin.Default()
	httpAdapter.RegisterRoutes(router)
	router.Run(":8080")
}
