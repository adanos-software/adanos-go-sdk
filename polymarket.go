package adanos

import (
	"context"
	"net/http"
	"strings"
)

// PolymarketService provides Polymarket API operations.
type PolymarketService struct{ client *Client }

// Trending calls the /polymarket/stocks/v1 Trending endpoint.
func (s *PolymarketService) Trending(ctx context.Context, o *TrendingOptions) ([]PolymarketTrendingStock, error) {
	q := o.query()
	return execute[[]PolymarketTrendingStock](s.client, ctx, http.MethodGet, "/polymarket/stocks/v1/trending", q, nil)
}

// TrendingSectors calls the /polymarket/stocks/v1 TrendingSectors endpoint.
func (s *PolymarketService) TrendingSectors(ctx context.Context, o *TrendingGroupOptions) ([]PolymarketTrendingSector, error) {
	q := o.query()
	return execute[[]PolymarketTrendingSector](s.client, ctx, http.MethodGet, "/polymarket/stocks/v1/trending/sectors", q, nil)
}

// TrendingCountries calls the /polymarket/stocks/v1 TrendingCountries endpoint.
func (s *PolymarketService) TrendingCountries(ctx context.Context, o *TrendingGroupOptions) ([]PolymarketTrendingCountry, error) {
	q := o.query()
	return execute[[]PolymarketTrendingCountry](s.client, ctx, http.MethodGet, "/polymarket/stocks/v1/trending/countries", q, nil)
}

// Stock calls the /polymarket/stocks/v1 Stock endpoint.
func (s *PolymarketService) Stock(ctx context.Context, ticker string, o *PeriodOptions) (*PolymarketStockDetailResponse, error) {
	q := o.query()
	return execute[*PolymarketStockDetailResponse](s.client, ctx, http.MethodGet, "/polymarket/stocks/v1/stock/"+segment(ticker), q, nil)
}

// Mentions calls the /polymarket/stocks/v1 Mentions endpoint.
func (s *PolymarketService) Mentions(ctx context.Context, ticker string, o *MentionsOptions) (*PolymarketRawMentionsResponse, error) {
	q := o.query()
	return execute[*PolymarketRawMentionsResponse](s.client, ctx, http.MethodGet, "/polymarket/stocks/v1/stock/"+segment(ticker)+"/mentions", q, nil)
}

// Search calls the /polymarket/stocks/v1 Search endpoint.
func (s *PolymarketService) Search(ctx context.Context, query string, o *SearchOptions) (*PolymarketSearchResponse, error) {
	q := o.query()
	q.Set("q", query)
	return execute[*PolymarketSearchResponse](s.client, ctx, http.MethodGet, "/polymarket/stocks/v1/search", q, nil)
}

// Compare calls the /polymarket/stocks/v1 Compare endpoint.
func (s *PolymarketService) Compare(ctx context.Context, tickers []string, o *PeriodOptions) (*PolymarketCompareResponse, error) {
	q := o.query()
	q.Set("tickers", strings.Join(tickers, ","))
	return execute[*PolymarketCompareResponse](s.client, ctx, http.MethodGet, "/polymarket/stocks/v1/compare", q, nil)
}

// MarketSentiment calls the /polymarket/stocks/v1 MarketSentiment endpoint.
func (s *PolymarketService) MarketSentiment(ctx context.Context, o *PeriodOptions) (*PolymarketMarketSentimentResponse, error) {
	q := o.query()
	return execute[*PolymarketMarketSentimentResponse](s.client, ctx, http.MethodGet, "/polymarket/stocks/v1/market-sentiment", q, nil)
}

// Stats calls the /polymarket/stocks/v1 Stats endpoint.
func (s *PolymarketService) Stats(ctx context.Context) (*PolymarketStatsResponse, error) {
	return execute[*PolymarketStatsResponse](s.client, ctx, http.MethodGet, "/polymarket/stocks/v1/stats", nil, nil)
}

// Health calls the /polymarket/stocks/v1 Health endpoint.
func (s *PolymarketService) Health(ctx context.Context) (*PolymarketHealthResponse, error) {
	return execute[*PolymarketHealthResponse](s.client, ctx, http.MethodGet, "/polymarket/stocks/v1/health", nil, nil)
}
