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

type AlphaClient struct {
	cfg        config
	httpClient *http.Client
	logger     *jsonlog.Logger
}

func NewClient(baseUrl, token string) *AlphaClient {
	return &AlphaClient{
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

func (c *AlphaClient) WithBaseUrl(url string) *AlphaClient {
	c.cfg.baseUrl = url
	return c
}

func (c *AlphaClient) WithHttpClient(httpClient *http.Client) *AlphaClient {
	c.httpClient = httpClient
	return c
}
