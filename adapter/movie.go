package adapter

import (
	"net/http"
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
	movieList, err := a.MovieUsecase.GetMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"movie_data": movieList})
}
