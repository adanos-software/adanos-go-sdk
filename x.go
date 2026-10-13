package adanos

import (
	"context"
	"net/http"
	"strings"
)

// XService provides X API operations.
type XService struct{ client *Client }

// Trending calls the /x/stocks/v1 Trending endpoint.
func (s *XService) Trending(ctx context.Context, o *TrendingOptions) ([]XTrendingStock, error) {
	q := o.query()
	return execute[[]XTrendingStock](s.client, ctx, http.MethodGet, "/x/stocks/v1/trending", q, nil)
}

// TrendingSectors calls the /x/stocks/v1 TrendingSectors endpoint.
func (s *XService) TrendingSectors(ctx context.Context, o *TrendingGroupOptions) ([]XTrendingSector, error) {
	q := o.query()
	return execute[[]XTrendingSector](s.client, ctx, http.MethodGet, "/x/stocks/v1/trending/sectors", q, nil)
}

// TrendingCountries calls the /x/stocks/v1 TrendingCountries endpoint.
func (s *XService) TrendingCountries(ctx context.Context, o *TrendingGroupOptions) ([]XTrendingCountry, error) {
	q := o.query()
	return execute[[]XTrendingCountry](s.client, ctx, http.MethodGet, "/x/stocks/v1/trending/countries", q, nil)
}

// Stock calls the /x/stocks/v1 Stock endpoint.
func (s *XService) Stock(ctx context.Context, ticker string, o *PeriodOptions) (*XStockDetailResponse, error) {
	q := o.query()
	return execute[*XStockDetailResponse](s.client, ctx, http.MethodGet, "/x/stocks/v1/stock/"+segment(ticker), q, nil)
}

// Mentions calls the /x/stocks/v1 Mentions endpoint.
func (s *XService) Mentions(ctx context.Context, ticker string, o *MentionsOptions) (*XRawMentionsResponse, error) {
	q := o.query()
	return execute[*XRawMentionsResponse](s.client, ctx, http.MethodGet, "/x/stocks/v1/stock/"+segment(ticker)+"/mentions", q, nil)
}

// Explain calls the /x/stocks/v1 Explain endpoint.
func (s *XService) Explain(ctx context.Context, ticker string) (*StockExplanationResponse, error) {
	return execute[*StockExplanationResponse](s.client, ctx, http.MethodGet, "/x/stocks/v1/stock/"+segment(ticker)+"/explain", nil, nil)
}

// Search calls the /x/stocks/v1 Search endpoint.
func (s *XService) Search(ctx context.Context, query string, o *SearchOptions) (*XSearchResponse, error) {
	q := o.query()
	q.Set("q", query)
	return execute[*XSearchResponse](s.client, ctx, http.MethodGet, "/x/stocks/v1/search", q, nil)
}

// Compare calls the /x/stocks/v1 Compare endpoint.
func (s *XService) Compare(ctx context.Context, tickers []string, o *PeriodOptions) (*XCompareResponse, error) {
	q := o.query()
	q.Set("tickers", strings.Join(tickers, ","))
	return execute[*XCompareResponse](s.client, ctx, http.MethodGet, "/x/stocks/v1/compare", q, nil)
}

// MarketSentiment calls the /x/stocks/v1 MarketSentiment endpoint.
func (s *XService) MarketSentiment(ctx context.Context, o *PeriodOptions) (*XMarketSentimentResponse, error) {
	q := o.query()
	return execute[*XMarketSentimentResponse](s.client, ctx, http.MethodGet, "/x/stocks/v1/market-sentiment", q, nil)
}

// Stats calls the /x/stocks/v1 Stats endpoint.
func (s *XService) Stats(ctx context.Context) (*XStatsResponse, error) {
	return execute[*XStatsResponse](s.client, ctx, http.MethodGet, "/x/stocks/v1/stats", nil, nil)
}

// Health calls the /x/stocks/v1 Health endpoint.
func (s *XService) Health(ctx context.Context) (*XHealthResponse, error) {
	return execute[*XHealthResponse](s.client, ctx, http.MethodGet, "/x/stocks/v1/health", nil, nil)
}
