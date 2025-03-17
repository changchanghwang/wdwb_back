package secClient

import (
	"net/http"

	rateLimitClient "github.com/changchanghwang/wdwb_back/pkg/rate-limit-client"
)

type secHttpClient struct {
	*rateLimitClient.RateLimitedClient
}

type secTransport struct {
	Transport http.RoundTripper
	UserAgent string
}

func (t *secTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.UserAgent)
	return t.Transport.RoundTrip(req)
}

// rps: 초당 요청 수
// burst: 동시 요청 수
func newSecHttpClient(rps int, burst int) *secHttpClient {
	return &secHttpClient{
		RateLimitedClient: rateLimitClient.New(rps, burst, &http.Client{
			Transport: &secTransport{
				Transport: http.DefaultTransport,
				UserAgent: "wdwb/1.0 (window95pill@gmail.com)",
			},
		}),
	}
}
