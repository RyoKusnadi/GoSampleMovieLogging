package usecase

import (
	"net/http"

	"scalable-go-movie/domain"
	"scalable-go-movie/infrastructure"
	"scalable-go-movie/usecase/utils"
)

const (
	TMDB_BASE_URL    = "https://api.themoviedb.org"
	TMDB_MOVIE_URL   = "/3/discover/movie?include_adult=false&include_video=false&language=en-US&page=1&sort_by=popularity.desc"
	TMDB_ERROR_FIELD = "status_message"
)

type MovieUsecase struct {
	HTTPClient *infrastructure.HTTPClient
}

func (uc *MovieUsecase) GetMovies() domain.ApiResponse {
	getMoviesEndPoint := TMDB_BASE_URL + TMDB_MOVIE_URL
	statusCode, response, err := uc.HTTPClient.HTTPRequest(http.MethodGet, getMoviesEndPoint, "")
	if err != nil {
		return utils.CreateDefaultErrorApiResponse(http.StatusBadRequest, err.Error())
	}

	successResponse := domain.TMDBSuccessCommonResponse{}
	errorResponse := domain.TMDBErrorCommonResponse{}
	successExcludeFields := map[string]bool{
		"page":          true,
		"total_pages":   true,
		"total_results": true,
	}

	return utils.DecodeResponse(statusCode, response, &successResponse, successExcludeFields, &errorResponse, TMDB_ERROR_FIELD)
}

func NewMovieUsecase(httpClient *infrastructure.HTTPClient) *MovieUsecase {
	return &MovieUsecase{
		HTTPClient: httpClient,
	}
}
