package adanos

import (
	"context"
	"net/http"
	"strings"
)

// NewsService provides News API operations.
type NewsService struct{ client *Client }

// Trending calls the /news/stocks/v1 Trending endpoint.
func (s *NewsService) Trending(ctx context.Context, o *NewsTrendingOptions) ([]NewsTrendingStock, error) {
	q := o.query()
	return execute[[]NewsTrendingStock](s.client, ctx, http.MethodGet, "/news/stocks/v1/trending", q, nil)
}

// TrendingSectors calls the /news/stocks/v1 TrendingSectors endpoint.
func (s *NewsService) TrendingSectors(ctx context.Context, o *NewsTrendingGroupOptions) ([]NewsTrendingSector, error) {
	q := o.query()
	return execute[[]NewsTrendingSector](s.client, ctx, http.MethodGet, "/news/stocks/v1/trending/sectors", q, nil)
}

// TrendingCountries calls the /news/stocks/v1 TrendingCountries endpoint.
func (s *NewsService) TrendingCountries(ctx context.Context, o *NewsTrendingGroupOptions) ([]NewsTrendingCountry, error) {
	q := o.query()
	return execute[[]NewsTrendingCountry](s.client, ctx, http.MethodGet, "/news/stocks/v1/trending/countries", q, nil)
}

// Stock calls the /news/stocks/v1 Stock endpoint.
func (s *NewsService) Stock(ctx context.Context, ticker string, o *PeriodOptions) (*NewsStockSentiment, error) {
	q := o.query()
	return execute[*NewsStockSentiment](s.client, ctx, http.MethodGet, "/news/stocks/v1/stock/"+segment(ticker), q, nil)
}

// Mentions calls the /news/stocks/v1 Mentions endpoint.
func (s *NewsService) Mentions(ctx context.Context, ticker string, o *MentionsOptions) (*NewsRawMentionsResponse, error) {
	q := o.query()
	return execute[*NewsRawMentionsResponse](s.client, ctx, http.MethodGet, "/news/stocks/v1/stock/"+segment(ticker)+"/mentions", q, nil)
}

// Explain calls the /news/stocks/v1 Explain endpoint.
func (s *NewsService) Explain(ctx context.Context, ticker string) (*StockExplanationResponse, error) {
	return execute[*StockExplanationResponse](s.client, ctx, http.MethodGet, "/news/stocks/v1/stock/"+segment(ticker)+"/explain", nil, nil)
}

// Search calls the /news/stocks/v1 Search endpoint.
func (s *NewsService) Search(ctx context.Context, query string, o *SearchOptions) (*NewsSearchResponse, error) {
	q := o.query()
	q.Set("q", query)
	return execute[*NewsSearchResponse](s.client, ctx, http.MethodGet, "/news/stocks/v1/search", q, nil)
}

// Compare calls the /news/stocks/v1 Compare endpoint.
func (s *NewsService) Compare(ctx context.Context, tickers []string, o *PeriodOptions) (*NewsCompareResponse, error) {
	q := o.query()
	q.Set("tickers", strings.Join(tickers, ","))
	return execute[*NewsCompareResponse](s.client, ctx, http.MethodGet, "/news/stocks/v1/compare", q, nil)
}

// MarketSentiment calls the /news/stocks/v1 MarketSentiment endpoint.
func (s *NewsService) MarketSentiment(ctx context.Context, o *PeriodOptions) (*NewsMarketSentimentResponse, error) {
	q := o.query()
	return execute[*NewsMarketSentimentResponse](s.client, ctx, http.MethodGet, "/news/stocks/v1/market-sentiment", q, nil)
}

// Stats calls the /news/stocks/v1 Stats endpoint.
func (s *NewsService) Stats(ctx context.Context) (*NewsStatsResponse, error) {
	return execute[*NewsStatsResponse](s.client, ctx, http.MethodGet, "/news/stocks/v1/stats", nil, nil)
}

// Health calls the /news/stocks/v1 Health endpoint.
func (s *NewsService) Health(ctx context.Context) (*HealthResponse, error) {
	return execute[*HealthResponse](s.client, ctx, http.MethodGet, "/news/stocks/v1/health", nil, nil)
}
