package alpha

import (
	"github.com/mhamm84/gofinance-alpha/internal/jsonlog"
	"net/http"
	"os"
	"time"
)

type config struct {
	token   string
	baseUrl string
	httpCfg struct {
	}
}

type Client struct {
	cfg        config
	httpClient *http.Client
	logger     *jsonlog.Logger
}

// NewClient creates a new Alpha Vantage client using a baseUrl and the token
func NewClient(baseUrl, token string) *Client {
	return &Client{
		cfg: config{
			token:   token,
			baseUrl: baseUrl,
		},
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
		logger: jsonlog.New(os.Stdout, jsonlog.LevelInfo),
	}
}

// WithBaseUrl overrides the baseUrl
func (c *Client) WithBaseUrl(url string) *Client {
	c.cfg.baseUrl = url
	return c
}

// WithHttpClient overrides the httpClient
func (c *Client) WithHttpClient(httpClient *http.Client) *Client {
	c.httpClient = httpClient
	return c
}
