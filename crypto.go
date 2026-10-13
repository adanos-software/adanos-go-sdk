package adanos

import (
	"context"
	"net/http"
	"strings"
)

// CryptoService provides Crypto API operations.
type CryptoService struct{ client *Client }

// Trending calls the /reddit/crypto/v1 Trending endpoint.
func (s *CryptoService) Trending(ctx context.Context, o *TrendingGroupOptions) ([]CryptoTrendingToken, error) {
	q := o.query()
	return execute[[]CryptoTrendingToken](s.client, ctx, http.MethodGet, "/reddit/crypto/v1/trending", q, nil)
}

// Token calls the /reddit/crypto/v1 Token endpoint.
func (s *CryptoService) Token(ctx context.Context, symbol string, o *PeriodOptions) (*CryptoTokenSentiment, error) {
	q := o.query()
	return execute[*CryptoTokenSentiment](s.client, ctx, http.MethodGet, "/reddit/crypto/v1/token/"+segment(symbol), q, nil)
}

// Mentions calls the /reddit/crypto/v1 Mentions endpoint.
func (s *CryptoService) Mentions(ctx context.Context, symbol string, o *RedditMentionsOptions) (*CryptoRawMentionsResponse, error) {
	q := o.query()
	return execute[*CryptoRawMentionsResponse](s.client, ctx, http.MethodGet, "/reddit/crypto/v1/token/"+segment(symbol)+"/mentions", q, nil)
}

// Search calls the /reddit/crypto/v1 Search endpoint.
func (s *CryptoService) Search(ctx context.Context, query string, o *SearchOptions) (*CryptoSearchResponse, error) {
	q := o.query()
	q.Set("q", query)
	return execute[*CryptoSearchResponse](s.client, ctx, http.MethodGet, "/reddit/crypto/v1/search", q, nil)
}

// Compare calls the /reddit/crypto/v1 Compare endpoint.
func (s *CryptoService) Compare(ctx context.Context, symbols []string, o *PeriodOptions) (*CryptoCompareResponse, error) {
	q := o.query()
	q.Set("symbols", strings.Join(symbols, ","))
	return execute[*CryptoCompareResponse](s.client, ctx, http.MethodGet, "/reddit/crypto/v1/compare", q, nil)
}

// MarketSentiment calls the /reddit/crypto/v1 MarketSentiment endpoint.
func (s *CryptoService) MarketSentiment(ctx context.Context, o *PeriodOptions) (*CryptoMarketSentimentResponse, error) {
	q := o.query()
	return execute[*CryptoMarketSentimentResponse](s.client, ctx, http.MethodGet, "/reddit/crypto/v1/market-sentiment", q, nil)
}

// Stats calls the /reddit/crypto/v1 Stats endpoint.
func (s *CryptoService) Stats(ctx context.Context) (*CryptoStatsResponse, error) {
	return execute[*CryptoStatsResponse](s.client, ctx, http.MethodGet, "/reddit/crypto/v1/stats", nil, nil)
}

// Health calls the /reddit/crypto/v1 Health endpoint.
func (s *CryptoService) Health(ctx context.Context) (*CryptoHealthResponse, error) {
	return execute[*CryptoHealthResponse](s.client, ctx, http.MethodGet, "/reddit/crypto/v1/health", nil, nil)
}
