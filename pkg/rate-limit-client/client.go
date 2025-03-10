package rateLimitClient

import (
	"context"
	"io"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type RateLimitedClient struct {
	client  *http.Client
	limiter *rate.Limiter
}

// @params rps: request per second
// @params burst: burst limit
// @params client: http client
func New(rps int, burst int, client *http.Client) *RateLimitedClient {
	if client == nil {
		client = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	return &RateLimitedClient{
		client:  client,
		limiter: rate.NewLimiter(rate.Limit(rps), burst),
	}
}

func (rlc *RateLimitedClient) Get(url string) (*http.Response, error) {
	if err := rlc.limiter.Wait(context.Background()); err != nil {
		return nil, err
	}
	return rlc.client.Get(url)
}

func (rlc *RateLimitedClient) Post(url string, body io.Reader) (*http.Response, error) {
	if err := rlc.limiter.Wait(context.Background()); err != nil {
		return nil, err
	}
	return rlc.client.Post(url, "application/json", body)
}
