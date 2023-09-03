package utils

import (
	"encoding/json"
	"net/http"
	"scalable-go-movie/domain"

	"github.com/gin-gonic/gin"
)

func MarshalAndSendResponse(c *gin.Context, response domain.Response) {
	status := http.StatusOK

	if response.Error != nil {
		status = http.StatusInternalServerError
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal response"})
		return
	}

	c.Data(status, "application/json; charset=utf-8", jsonResponse)
}
