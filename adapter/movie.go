package adapter

import (
	"scalable-go-movie/adapter/utils"
	"scalable-go-movie/usecase"

	"github.com/gin-gonic/gin"
)

type HTTPAdapter struct {
	MovieUsecase *usecase.MovieUsecase
}

func NewHTTPAdapter(movieUsecase *usecase.MovieUsecase) *HTTPAdapter {
	return &HTTPAdapter{
		MovieUsecase: movieUsecase,
	}
}

func (a *HTTPAdapter) RegisterRoutes(router *gin.Engine) {
	router.GET("/movies", a.GetMovieHandler)
}

func (a *HTTPAdapter) GetMovieHandler(c *gin.Context) {
	utils.MarshalAndSendResponse(c, a.MovieUsecase.GetMovies())
}
