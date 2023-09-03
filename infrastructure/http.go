package infrastructure

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	DEFAULT_TIMEOUT = 10 * time.Second
)

type CommonResponse struct {
	Page         int           `json:"page"`
	Results      []interface{} `json:"results"`
	TotalPages   int           `json:"total_pages"`
	TotalResults int           `json:"total_results"`
}

type HTTPClient struct {
	Client         *http.Client
	Timeout        time.Duration
	RequestHeaders map[string]string
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		Client:         &http.Client{},
		Timeout:        DEFAULT_TIMEOUT,
		RequestHeaders: make(map[string]string),
	}
}

func (c *HTTPClient) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}

func (c *HTTPClient) SetRequestHeaders(headers map[string]string) {
	c.RequestHeaders = headers
}

func (c *HTTPClient) HTTPRequest(method, url, body string) (statusCode int, responseBody []byte, err error) {
	if !isValidHTTPMethod(method) {
		return http.StatusBadRequest, nil, fmt.Errorf("invalid HTTP method: %s", method)
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	for key, value := range c.RequestHeaders {
		req.Header.Set(key, value)
	}

	c.Client.Timeout = c.Timeout

	response, err := c.Client.Do(req)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	defer response.Body.Close()

	responseBody, err = io.ReadAll(response.Body)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, responseBody, nil
}

func isValidHTTPMethod(method string) bool {
	return method == http.MethodGet || method == http.MethodPost ||
		method == http.MethodPut || method == http.MethodDelete ||
		method == http.MethodPatch || method == http.MethodOptions
}
