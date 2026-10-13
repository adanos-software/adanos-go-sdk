package adanos

import (
	"context"
	"net/http"
	"strings"
)

// RedditService provides Reddit API operations.
type RedditService struct{ client *Client }

// Trending calls the /reddit/stocks/v1 Trending endpoint.
func (s *RedditService) Trending(ctx context.Context, o *TrendingOptions) ([]TrendingStock, error) {
	q := o.query()
	return execute[[]TrendingStock](s.client, ctx, http.MethodGet, "/reddit/stocks/v1/trending", q, nil)
}

// TrendingSectors calls the /reddit/stocks/v1 TrendingSectors endpoint.
func (s *RedditService) TrendingSectors(ctx context.Context, o *TrendingGroupOptions) ([]TrendingSector, error) {
	q := o.query()
	return execute[[]TrendingSector](s.client, ctx, http.MethodGet, "/reddit/stocks/v1/trending/sectors", q, nil)
}

// TrendingCountries calls the /reddit/stocks/v1 TrendingCountries endpoint.
func (s *RedditService) TrendingCountries(ctx context.Context, o *TrendingGroupOptions) ([]TrendingCountry, error) {
	q := o.query()
	return execute[[]TrendingCountry](s.client, ctx, http.MethodGet, "/reddit/stocks/v1/trending/countries", q, nil)
}

// Stock calls the /reddit/stocks/v1 Stock endpoint.
func (s *RedditService) Stock(ctx context.Context, ticker string, o *PeriodOptions) (*StockSentiment, error) {
	q := o.query()
	return execute[*StockSentiment](s.client, ctx, http.MethodGet, "/reddit/stocks/v1/stock/"+segment(ticker), q, nil)
}

// Mentions calls the /reddit/stocks/v1 Mentions endpoint.
func (s *RedditService) Mentions(ctx context.Context, ticker string, o *RedditMentionsOptions) (*RedditRawMentionsResponse, error) {
	q := o.query()
	return execute[*RedditRawMentionsResponse](s.client, ctx, http.MethodGet, "/reddit/stocks/v1/stock/"+segment(ticker)+"/mentions", q, nil)
}

// Explain calls the /reddit/stocks/v1 Explain endpoint.
func (s *RedditService) Explain(ctx context.Context, ticker string) (*StockExplanationResponse, error) {
	return execute[*StockExplanationResponse](s.client, ctx, http.MethodGet, "/reddit/stocks/v1/stock/"+segment(ticker)+"/explain", nil, nil)
}

// Search calls the /reddit/stocks/v1 Search endpoint.
func (s *RedditService) Search(ctx context.Context, query string, o *SearchOptions) (*SearchResponse, error) {
	q := o.query()
	q.Set("q", query)
	return execute[*SearchResponse](s.client, ctx, http.MethodGet, "/reddit/stocks/v1/search", q, nil)
}

// Compare calls the /reddit/stocks/v1 Compare endpoint.
func (s *RedditService) Compare(ctx context.Context, tickers []string, o *PeriodOptions) (*CompareResponse, error) {
	q := o.query()
	q.Set("tickers", strings.Join(tickers, ","))
	return execute[*CompareResponse](s.client, ctx, http.MethodGet, "/reddit/stocks/v1/compare", q, nil)
}

// MarketSentiment calls the /reddit/stocks/v1 MarketSentiment endpoint.
func (s *RedditService) MarketSentiment(ctx context.Context, o *PeriodOptions) (*RedditMarketSentimentResponse, error) {
	q := o.query()
	return execute[*RedditMarketSentimentResponse](s.client, ctx, http.MethodGet, "/reddit/stocks/v1/market-sentiment", q, nil)
}

// Stats calls the /reddit/stocks/v1 Stats endpoint.
func (s *RedditService) Stats(ctx context.Context) (*RedditStatsResponse, error) {
	return execute[*RedditStatsResponse](s.client, ctx, http.MethodGet, "/reddit/stocks/v1/stats", nil, nil)
}

// Health calls the /reddit/stocks/v1 Health endpoint.
func (s *RedditService) Health(ctx context.Context) (*HealthResponse, error) {
	return execute[*HealthResponse](s.client, ctx, http.MethodGet, "/reddit/stocks/v1/health", nil, nil)
}
