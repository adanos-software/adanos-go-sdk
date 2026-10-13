package adanos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// Version is the SDK release version.
	Version = "1.0.0"
	// APIVersion is the API contract used to generate the SDK models.
	APIVersion = "1.54.0"
	// DefaultBaseURL is the public Adanos API endpoint.
	DefaultBaseURL = "https://api.adanos.com"
)

// Config controls authentication, the API origin, and request deadlines.
type Config struct {
	APIKey  string
	BaseURL string
	// Timeout bounds the entire operation. Zero selects 30 seconds.
	Timeout time.Duration
	// HTTPClient supplies an optional transport and redirect policy. It is copied.
	HTTPClient *http.Client
}

// Client can be reused concurrently. Caller-owned transports and options must
// also support concurrent access without mutation.
type Client struct {
	baseURL    string
	apiKey     string
	timeout    time.Duration
	httpClient *http.Client
	reddit     RedditService
	crypto     CryptoService
	x          XService
	news       NewsService
	polymarket PolymarketService
	sentiment  SentimentService
}

// NewClient validates configuration without performing a network request.
// An empty API key permits public health calls. Protected calls return APIError.
func NewClient(config Config) (*Client, error) {
	if strings.ContainsAny(config.APIKey, "\r\n") {
		return nil, fmt.Errorf("adanos: API key contains a newline")
	}
	if config.Timeout < 0 {
		return nil, fmt.Errorf("adanos: timeout must not be negative")
	}
	base := config.BaseURL
	if base == "" {
		base = DefaultBaseURL
	}
	u, err := url.Parse(base)
	if err != nil {
		return nil, fmt.Errorf("adanos: invalid base URL: %w", err)
	}
	if (u.Scheme != "http" && u.Scheme != "https") || u.Hostname() == "" || u.User != nil || u.RawQuery != "" || u.ForceQuery || u.Fragment != "" {
		return nil, fmt.Errorf("adanos: base URL must be an absolute HTTP(S) URL without credentials, query, or fragment")
	}
	base = strings.TrimRight(u.String(), "/")
	timeout := config.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}
	hc := http.Client{}
	if config.HTTPClient != nil {
		hc = *config.HTTPClient
	}
	originalRedirect := hc.CheckRedirect
	origin := *u
	hc.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) == 0 {
			return fmt.Errorf("adanos: redirect has no origin")
		}
		if !sameOrigin(&origin, req) {
			return http.ErrUseLastResponse
		}
		if originalRedirect != nil {
			if err := originalRedirect(req, via); err != nil {
				return err
			}
		} else if len(via) >= 10 {
			return fmt.Errorf("stopped after 10 redirects")
		}
		if !sameOrigin(&origin, req) {
			return http.ErrUseLastResponse
		}
		return nil
	}
	c := &Client{baseURL: base, apiKey: config.APIKey, timeout: timeout, httpClient: &hc}
	c.reddit.client = c
	c.crypto.client = c
	c.x.client = c
	c.news.client = c
	c.polymarket.client = c
	c.sentiment.client = c
	return c, nil
}

func sameOrigin(origin *url.URL, req *http.Request) bool {
	if req.URL == nil || !strings.EqualFold(origin.Scheme, req.URL.Scheme) || !strings.EqualFold(origin.Host, req.URL.Host) ||
		(req.Host != "" && !strings.EqualFold(origin.Host, req.Host)) {
		return false
	}
	if req.URL.Opaque == "" {
		return true
	}
	wireURL, err := url.ParseRequestURI(req.URL.RequestURI())
	return err == nil && wireURL.Opaque == "" &&
		((wireURL.Scheme == "" && wireURL.Host == "") ||
			(strings.EqualFold(origin.Scheme, wireURL.Scheme) && strings.EqualFold(origin.Host, wireURL.Host)))
}

// Reddit returns the Reddit Stocks client.
func (c *Client) Reddit() *RedditService { return &c.reddit }

// Crypto returns the Reddit Crypto client.
func (c *Client) Crypto() *CryptoService { return &c.crypto }

// RedditCrypto is an alias for Crypto.
func (c *Client) RedditCrypto() *CryptoService { return &c.crypto }

// X returns the X Stocks client.
func (c *Client) X() *XService { return &c.x }

// News returns the News Stocks client.
func (c *Client) News() *NewsService { return &c.news }

// Polymarket returns the Polymarket Stocks client.
func (c *Client) Polymarket() *PolymarketService { return &c.polymarket }

// Sentiment returns the direct text analysis client.
func (c *Client) Sentiment() *SentimentService { return &c.sentiment }

// Health returns the root API health report.
func (c *Client) Health(ctx context.Context) (*RootHealthResponse, error) {
	return execute[*RootHealthResponse](c, ctx, http.MethodGet, "/health", nil, nil)
}
