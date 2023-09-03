package usecase

import (
	"encoding/json"
	"errors"
	"net/http"

	"scalable-go-movie/domain"
	"scalable-go-movie/infrastructure"
)

const (
	TMDB_BASE_URL  = "https://api.themoviedb.org"
	TMDB_MOVIE_URL = "/3/discover/movie?include_adult=false&include_video=false&language=en-US&page=1&sort_by=popularity.desc"
)

type MovieUsecase struct {
	HTTPClient *infrastructure.HTTPClient
}

func (uc *MovieUsecase) GetMovies() domain.Response {
	endpoint := TMDB_BASE_URL + TMDB_MOVIE_URL
	response, err := uc.HTTPClient.HTTPRequest(http.MethodGet, endpoint, "")
	if err != nil {
		return domain.CreateErrorResponse(http.StatusBadRequest, err.Error())
	}

	var commonResp infrastructure.CommonResponse
	if err := json.Unmarshal(response, &commonResp); err != nil {
		return domain.CreateErrorResponse(http.StatusBadRequest, err.Error())
	}

	movieList, err := convertToMovies(commonResp.Results)
	if err != nil {
		return domain.CreateErrorResponse(http.StatusBadRequest, err.Error())
	}

	return domain.CreateSuccessResponse(movieList)
}

func convertToMovies(data []interface{}) (domain.MovieList, error) {
	var movies domain.MovieList

	for _, item := range data {
		movieMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, errors.New("failed to convert item to movie")
		}

		var movie domain.Movie
		if movieId, ok := movieMap["id"].(float64); ok {
			movie.ID = int(movieId)
		}

		movie.Adult, _ = movieMap["adult"].(bool)
		movie.BackdropPath, _ = movieMap["backdrop_path"].(string)
		genreIds, _ := movieMap["genre_ids"].([]interface{})
		movie.GenreIds = make([]int, len(genreIds))
		for i, id := range genreIds {
			if genreId, ok := id.(float64); ok {
				movie.GenreIds[i] = int(genreId)
			}
		}

		movie.OriginalLanguage, _ = movieMap["original_language"].(string)
		movie.OriginalTitle, _ = movieMap["original_title"].(string)
		movie.Overview, _ = movieMap["overview"].(string)
		movie.Popularity, _ = movieMap["popularity"].(float64)
		movie.PosterPath, _ = movieMap["poster_path"].(string)
		movie.ReleaseDate, _ = movieMap["release_date"].(string)
		movie.Title, _ = movieMap["title"].(string)
		movie.Video, _ = movieMap["video"].(bool)

		if voteAverage, ok := movieMap["vote_average"].(float64); ok {
			movie.VoteAverage = voteAverage
		}

		if voteCount, ok := movieMap["vote_count"].(float64); ok {
			movie.VoteCount = int(voteCount)
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func NewMovieUsecase(httpClient *infrastructure.HTTPClient) *MovieUsecase {
	return &MovieUsecase{
		HTTPClient: httpClient,
	}
}
