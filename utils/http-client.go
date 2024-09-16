package utils

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type HttpClient struct {
	client *resty.Client
}

func NewHttpClient(baseURL string) *HttpClient {
	client := resty.New().
		SetBaseURL(baseURL).
		SetRetryCount(3)
	return &HttpClient{client: client}
}

func (h *HttpClient) Get(url string, result interface{}) error {
	resp, err := h.client.R().
		SetResult(result).
		Get(url)

	if err != nil {
		return fmt.Errorf("failed to fetch data: %v", err)
	}

	if resp.IsError() {
		return fmt.Errorf("HTTP error: %s", resp.Status())
	}

	return nil
}
