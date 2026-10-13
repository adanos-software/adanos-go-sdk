package adanos_test

import (
	"context"
	adanos "github.com/adanos-software/adanos-go-sdk"
	"net/url"
)

func operationCases() []operationCase {
	return []operationCase{
		{name: "Reddit.Trending", method: "GET", template: "/reddit/stocks/v1/trending", path: "/reddit/stocks/v1/trending", fixture: `[{"ticker":"NU.V"}]`, field: "ticker", want: "NU.V", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}, "type": {""}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Reddit().Trending(ctx, &adanos.TrendingOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0), Type: new("")})
		}},
		{name: "Reddit.TrendingSectors", method: "GET", template: "/reddit/stocks/v1/trending/sectors", path: "/reddit/stocks/v1/trending/sectors", fixture: `[{"sector":"Technology"}]`, field: "sector", want: "Technology", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Reddit().TrendingSectors(ctx, &adanos.TrendingGroupOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0)})
		}},
		{name: "Reddit.TrendingCountries", method: "GET", template: "/reddit/stocks/v1/trending/countries", path: "/reddit/stocks/v1/trending/countries", fixture: `[{"country":"CA"}]`, field: "country", want: "CA", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Reddit().TrendingCountries(ctx, &adanos.TrendingGroupOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0)})
		}},
		{name: "Reddit.Stock", method: "GET", template: "/reddit/stocks/v1/stock/{ticker}", path: "/reddit/stocks/v1/stock/NU.V", fixture: `{"ticker":"NU.V"}`, field: "ticker", want: "NU.V", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Reddit().Stock(ctx, "NU.V", &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "Reddit.Mentions", method: "GET", template: "/reddit/stocks/v1/stock/{ticker}/mentions", path: "/reddit/stocks/v1/stock/NU.V/mentions", fixture: `{"ticker":"NU.V"}`, field: "ticker", want: "NU.V", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}, "include_inherited": {"false"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Reddit().Mentions(ctx, "NU.V", &adanos.RedditMentionsOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0), IncludeInherited: new(false)})
		}},
		{name: "Reddit.Explain", method: "GET", template: "/reddit/stocks/v1/stock/{ticker}/explain", path: "/reddit/stocks/v1/stock/NU.V/explain", fixture: `{"ticker":"NU.V"}`, field: "ticker", want: "NU.V", query: url.Values{}, call: func(ctx context.Context, c *adanos.Client) (any, error) { return c.Reddit().Explain(ctx, "NU.V") }},
		{name: "Reddit.Search", method: "GET", template: "/reddit/stocks/v1/search", path: "/reddit/stocks/v1/search", fixture: `{"query":"EV"}`, field: "query", want: "EV", query: url.Values{"q": {"EV"}, "limit": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Reddit().Search(ctx, "EV", &adanos.SearchOptions{Limit: new(0)})
		}},
		{name: "Reddit.Compare", method: "GET", template: "/reddit/stocks/v1/compare", path: "/reddit/stocks/v1/compare", fixture: `{"period_days":23}`, field: "period_days", want: float64(23), query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "tickers": {"NU.V,brk.a"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Reddit().Compare(ctx, []string{"NU.V", "brk.a"}, &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "Reddit.MarketSentiment", method: "GET", template: "/reddit/stocks/v1/market-sentiment", path: "/reddit/stocks/v1/market-sentiment", fixture: `{"mentions":23}`, field: "mentions", want: float64(23), query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Reddit().MarketSentiment(ctx, &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "Reddit.Stats", method: "GET", template: "/reddit/stocks/v1/stats", path: "/reddit/stocks/v1/stats", fixture: `{"total_mentions":23}`, field: "total_mentions", want: float64(23), query: url.Values{}, call: func(ctx context.Context, c *adanos.Client) (any, error) { return c.Reddit().Stats(ctx) }},
		{name: "Reddit.Health", method: "GET", template: "/reddit/stocks/v1/health", path: "/reddit/stocks/v1/health", fixture: `{"status":"healthy"}`, field: "status", want: "healthy", query: url.Values{}, call: func(ctx context.Context, c *adanos.Client) (any, error) { return c.Reddit().Health(ctx) }},
		{name: "Crypto.Trending", method: "GET", template: "/reddit/crypto/v1/trending", path: "/reddit/crypto/v1/trending", fixture: `[{"symbol":"BTC"}]`, field: "symbol", want: "BTC", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Crypto().Trending(ctx, &adanos.TrendingGroupOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0)})
		}},
		{name: "Crypto.Token", method: "GET", template: "/reddit/crypto/v1/token/{symbol}", path: "/reddit/crypto/v1/token/BTC", fixture: `{"symbol":"BTC"}`, field: "symbol", want: "BTC", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Crypto().Token(ctx, "BTC", &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "Crypto.Mentions", method: "GET", template: "/reddit/crypto/v1/token/{symbol}/mentions", path: "/reddit/crypto/v1/token/BTC/mentions", fixture: `{"symbol":"BTC"}`, field: "symbol", want: "BTC", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}, "include_inherited": {"false"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Crypto().Mentions(ctx, "BTC", &adanos.RedditMentionsOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0), IncludeInherited: new(false)})
		}},
		{name: "Crypto.Search", method: "GET", template: "/reddit/crypto/v1/search", path: "/reddit/crypto/v1/search", fixture: `{"query":"EV"}`, field: "query", want: "EV", query: url.Values{"q": {"EV"}, "limit": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Crypto().Search(ctx, "EV", &adanos.SearchOptions{Limit: new(0)})
		}},
		{name: "Crypto.Compare", method: "GET", template: "/reddit/crypto/v1/compare", path: "/reddit/crypto/v1/compare", fixture: `{"period_days":23}`, field: "period_days", want: float64(23), query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "symbols": {"btc,ETH"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Crypto().Compare(ctx, []string{"btc", "ETH"}, &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "Crypto.MarketSentiment", method: "GET", template: "/reddit/crypto/v1/market-sentiment", path: "/reddit/crypto/v1/market-sentiment", fixture: `{"mentions":23}`, field: "mentions", want: float64(23), query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Crypto().MarketSentiment(ctx, &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "Crypto.Stats", method: "GET", template: "/reddit/crypto/v1/stats", path: "/reddit/crypto/v1/stats", fixture: `{"total_mentions":23}`, field: "total_mentions", want: float64(23), query: url.Values{}, call: func(ctx context.Context, c *adanos.Client) (any, error) { return c.Crypto().Stats(ctx) }},
		{name: "Crypto.Health", method: "GET", template: "/reddit/crypto/v1/health", path: "/reddit/crypto/v1/health", fixture: `{"status":"healthy"}`, field: "status", want: "healthy", query: url.Values{}, call: func(ctx context.Context, c *adanos.Client) (any, error) { return c.Crypto().Health(ctx) }},
		{name: "X.Trending", method: "GET", template: "/x/stocks/v1/trending", path: "/x/stocks/v1/trending", fixture: `[{"ticker":"NU.V"}]`, field: "ticker", want: "NU.V", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}, "type": {""}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.X().Trending(ctx, &adanos.TrendingOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0), Type: new("")})
		}},
		{name: "X.TrendingSectors", method: "GET", template: "/x/stocks/v1/trending/sectors", path: "/x/stocks/v1/trending/sectors", fixture: `[{"sector":"Technology"}]`, field: "sector", want: "Technology", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.X().TrendingSectors(ctx, &adanos.TrendingGroupOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0)})
		}},
		{name: "X.TrendingCountries", method: "GET", template: "/x/stocks/v1/trending/countries", path: "/x/stocks/v1/trending/countries", fixture: `[{"country":"CA"}]`, field: "country", want: "CA", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.X().TrendingCountries(ctx, &adanos.TrendingGroupOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0)})
		}},
		{name: "X.Stock", method: "GET", template: "/x/stocks/v1/stock/{ticker}", path: "/x/stocks/v1/stock/NU.V", fixture: `{"ticker":"NU.V"}`, field: "ticker", want: "NU.V", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.X().Stock(ctx, "NU.V", &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "X.Mentions", method: "GET", template: "/x/stocks/v1/stock/{ticker}/mentions", path: "/x/stocks/v1/stock/NU.V/mentions", fixture: `{"ticker":"NU.V"}`, field: "ticker", want: "NU.V", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.X().Mentions(ctx, "NU.V", &adanos.MentionsOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0)})
		}},
		{name: "X.Explain", method: "GET", template: "/x/stocks/v1/stock/{ticker}/explain", path: "/x/stocks/v1/stock/NU.V/explain", fixture: `{"ticker":"NU.V"}`, field: "ticker", want: "NU.V", query: url.Values{}, call: func(ctx context.Context, c *adanos.Client) (any, error) { return c.X().Explain(ctx, "NU.V") }},
		{name: "X.Search", method: "GET", template: "/x/stocks/v1/search", path: "/x/stocks/v1/search", fixture: `{"query":"EV"}`, field: "query", want: "EV", query: url.Values{"q": {"EV"}, "limit": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.X().Search(ctx, "EV", &adanos.SearchOptions{Limit: new(0)})
		}},
		{name: "X.Compare", method: "GET", template: "/x/stocks/v1/compare", path: "/x/stocks/v1/compare", fixture: `{"period_days":23}`, field: "period_days", want: float64(23), query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "tickers": {"NU.V,brk.a"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.X().Compare(ctx, []string{"NU.V", "brk.a"}, &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "X.MarketSentiment", method: "GET", template: "/x/stocks/v1/market-sentiment", path: "/x/stocks/v1/market-sentiment", fixture: `{"mentions":23}`, field: "mentions", want: float64(23), query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.X().MarketSentiment(ctx, &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "X.Stats", method: "GET", template: "/x/stocks/v1/stats", path: "/x/stocks/v1/stats", fixture: `{"total_mentions":23}`, field: "total_mentions", want: float64(23), query: url.Values{}, call: func(ctx context.Context, c *adanos.Client) (any, error) { return c.X().Stats(ctx) }},
		{name: "X.Health", method: "GET", template: "/x/stocks/v1/health", path: "/x/stocks/v1/health", fixture: `{"status":"healthy"}`, field: "status", want: "healthy", query: url.Values{}, call: func(ctx context.Context, c *adanos.Client) (any, error) { return c.X().Health(ctx) }},
		{name: "News.Trending", method: "GET", template: "/news/stocks/v1/trending", path: "/news/stocks/v1/trending", fixture: `[{"ticker":"NU.V"}]`, field: "ticker", want: "NU.V", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}, "type": {""}, "source": {"Reuters"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.News().Trending(ctx, &adanos.NewsTrendingOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0), Type: new(""), Source: new("Reuters")})
		}},
		{name: "News.TrendingSectors", method: "GET", template: "/news/stocks/v1/trending/sectors", path: "/news/stocks/v1/trending/sectors", fixture: `[{"sector":"Technology"}]`, field: "sector", want: "Technology", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}, "source": {"Reuters"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.News().TrendingSectors(ctx, &adanos.NewsTrendingGroupOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0), Source: new("Reuters")})
		}},
		{name: "News.TrendingCountries", method: "GET", template: "/news/stocks/v1/trending/countries", path: "/news/stocks/v1/trending/countries", fixture: `[{"country":"CA"}]`, field: "country", want: "CA", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}, "source": {"Reuters"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.News().TrendingCountries(ctx, &adanos.NewsTrendingGroupOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0), Source: new("Reuters")})
		}},
		{name: "News.Stock", method: "GET", template: "/news/stocks/v1/stock/{ticker}", path: "/news/stocks/v1/stock/NU.V", fixture: `{"ticker":"NU.V"}`, field: "ticker", want: "NU.V", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.News().Stock(ctx, "NU.V", &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "News.Mentions", method: "GET", template: "/news/stocks/v1/stock/{ticker}/mentions", path: "/news/stocks/v1/stock/NU.V/mentions", fixture: `{"ticker":"NU.V"}`, field: "ticker", want: "NU.V", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.News().Mentions(ctx, "NU.V", &adanos.MentionsOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0)})
		}},
		{name: "News.Explain", method: "GET", template: "/news/stocks/v1/stock/{ticker}/explain", path: "/news/stocks/v1/stock/NU.V/explain", fixture: `{"ticker":"NU.V"}`, field: "ticker", want: "NU.V", query: url.Values{}, call: func(ctx context.Context, c *adanos.Client) (any, error) { return c.News().Explain(ctx, "NU.V") }},
		{name: "News.Search", method: "GET", template: "/news/stocks/v1/search", path: "/news/stocks/v1/search", fixture: `{"query":"EV"}`, field: "query", want: "EV", query: url.Values{"q": {"EV"}, "limit": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.News().Search(ctx, "EV", &adanos.SearchOptions{Limit: new(0)})
		}},
		{name: "News.Compare", method: "GET", template: "/news/stocks/v1/compare", path: "/news/stocks/v1/compare", fixture: `{"period_days":23}`, field: "period_days", want: float64(23), query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "tickers": {"NU.V,brk.a"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.News().Compare(ctx, []string{"NU.V", "brk.a"}, &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "News.MarketSentiment", method: "GET", template: "/news/stocks/v1/market-sentiment", path: "/news/stocks/v1/market-sentiment", fixture: `{"mentions":23}`, field: "mentions", want: float64(23), query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.News().MarketSentiment(ctx, &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "News.Stats", method: "GET", template: "/news/stocks/v1/stats", path: "/news/stocks/v1/stats", fixture: `{"total_mentions":23}`, field: "total_mentions", want: float64(23), query: url.Values{}, call: func(ctx context.Context, c *adanos.Client) (any, error) { return c.News().Stats(ctx) }},
		{name: "News.Health", method: "GET", template: "/news/stocks/v1/health", path: "/news/stocks/v1/health", fixture: `{"status":"healthy"}`, field: "status", want: "healthy", query: url.Values{}, call: func(ctx context.Context, c *adanos.Client) (any, error) { return c.News().Health(ctx) }},
		{name: "Polymarket.Trending", method: "GET", template: "/polymarket/stocks/v1/trending", path: "/polymarket/stocks/v1/trending", fixture: `[{"ticker":"NU.V"}]`, field: "ticker", want: "NU.V", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}, "type": {""}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Polymarket().Trending(ctx, &adanos.TrendingOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0), Type: new("")})
		}},
		{name: "Polymarket.TrendingSectors", method: "GET", template: "/polymarket/stocks/v1/trending/sectors", path: "/polymarket/stocks/v1/trending/sectors", fixture: `[{"sector":"Technology"}]`, field: "sector", want: "Technology", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Polymarket().TrendingSectors(ctx, &adanos.TrendingGroupOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0)})
		}},
		{name: "Polymarket.TrendingCountries", method: "GET", template: "/polymarket/stocks/v1/trending/countries", path: "/polymarket/stocks/v1/trending/countries", fixture: `[{"country":"CA"}]`, field: "country", want: "CA", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Polymarket().TrendingCountries(ctx, &adanos.TrendingGroupOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0)})
		}},
		{name: "Polymarket.Stock", method: "GET", template: "/polymarket/stocks/v1/stock/{ticker}", path: "/polymarket/stocks/v1/stock/NU.V", fixture: `{"ticker":"NU.V"}`, field: "ticker", want: "NU.V", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Polymarket().Stock(ctx, "NU.V", &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "Polymarket.Mentions", method: "GET", template: "/polymarket/stocks/v1/stock/{ticker}/mentions", path: "/polymarket/stocks/v1/stock/NU.V/mentions", fixture: `{"ticker":"NU.V"}`, field: "ticker", want: "NU.V", query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "limit": {"0"}, "offset": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Polymarket().Mentions(ctx, "NU.V", &adanos.MentionsOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7), Limit: new(0), Offset: new(0)})
		}},
		{name: "Polymarket.Search", method: "GET", template: "/polymarket/stocks/v1/search", path: "/polymarket/stocks/v1/search", fixture: `{"query":"EV"}`, field: "query", want: "EV", query: url.Values{"q": {"EV"}, "limit": {"0"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Polymarket().Search(ctx, "EV", &adanos.SearchOptions{Limit: new(0)})
		}},
		{name: "Polymarket.Compare", method: "GET", template: "/polymarket/stocks/v1/compare", path: "/polymarket/stocks/v1/compare", fixture: `{"period_days":23}`, field: "period_days", want: float64(23), query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}, "tickers": {"NU.V,brk.a"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Polymarket().Compare(ctx, []string{"NU.V", "brk.a"}, &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "Polymarket.MarketSentiment", method: "GET", template: "/polymarket/stocks/v1/market-sentiment", path: "/polymarket/stocks/v1/market-sentiment", fixture: `{"trade_count":23}`, field: "trade_count", want: float64(23), query: url.Values{"from": {"2026-10-01"}, "to": {"2026-10-13"}, "days": {"7"}}, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Polymarket().MarketSentiment(ctx, &adanos.PeriodOptions{From: new("2026-10-01"), To: new("2026-10-13"), Days: new(7)})
		}},
		{name: "Polymarket.Stats", method: "GET", template: "/polymarket/stocks/v1/stats", path: "/polymarket/stocks/v1/stats", fixture: `{"total_trades":23}`, field: "total_trades", want: float64(23), query: url.Values{}, call: func(ctx context.Context, c *adanos.Client) (any, error) { return c.Polymarket().Stats(ctx) }},
		{name: "Polymarket.Health", method: "GET", template: "/polymarket/stocks/v1/health", path: "/polymarket/stocks/v1/health", fixture: `{"status":"healthy"}`, field: "status", want: "healthy", query: url.Values{}, call: func(ctx context.Context, c *adanos.Client) (any, error) { return c.Polymarket().Health(ctx) }},
		{name: "Root.Health", method: "GET", template: "/health", path: "/health", fixture: `{"status":"healthy"}`, field: "status", want: "healthy", query: url.Values{}, call: func(ctx context.Context, c *adanos.Client) (any, error) { return c.Health(ctx) }},
		{name: "Sentiment.Analyze", method: "POST", template: "/sentiment/v1/analyze", path: "/sentiment/v1/analyze", fixture: `{"text":"sample","sentiment_score":0.25}`, field: "sentiment_score", want: float64(0.25), query: url.Values{}, body: `{"text":"sample","target":"Apple"}`, call: func(ctx context.Context, c *adanos.Client) (any, error) {
			return c.Sentiment().Analyze(ctx, "sample", &adanos.SentimentAnalyzeOptions{Target: new("Apple")})
		}},
	}
}
