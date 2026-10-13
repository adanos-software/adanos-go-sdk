# Adanos Go SDK

Go client for the [Adanos Market Sentiment API](https://api.adanos.com/docs), generated from API **1.54.0**.
Version **1.0.0** covers all 53 operations and 113 schemas exposed by the API contract.
The SDK requires Go 1.26 or newer and uses only the Go standard library.

## Install

```sh
go get github.com/adanos-software/adanos-go-sdk@v1.0.0
```

## Get stock sentiment

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	adanos "github.com/adanos-software/adanos-go-sdk"
)

func main() {
	client, err := adanos.NewClient(adanos.Config{
		APIKey: os.Getenv("ADANOS_API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stocks, err := client.Reddit().Trending(ctx, &adanos.TrendingOptions{Limit: new(3)})
	if err != nil {
		log.Fatal(err)
	}
	for _, stock := range stocks {
		fmt.Println(stock.Ticker, stock.BuzzScore, stock.SentimentScore)
	}
}
```

Get an API key from [Adanos](https://adanos.com). An empty key can call public health endpoints;
protected endpoints return the server's authentication error.

## Clients and operations

Each method accepts `context.Context` as its first argument. Object responses return a pointer;
list responses return a slice. A failed call returns a nil result and an error.

| Client | Operations |
| --- | --- |
| `client.Reddit()` | Trending, TrendingSectors, TrendingCountries, Stock, Mentions, Explain, Search, Compare, MarketSentiment, Stats, Health |
| `client.Crypto()` | Trending, Token, Mentions, Search, Compare, MarketSentiment, Stats, Health |
| `client.X()` | Trending, TrendingSectors, TrendingCountries, Stock, Mentions, Explain, Search, Compare, MarketSentiment, Stats, Health |
| `client.News()` | Trending, TrendingSectors, TrendingCountries, Stock, Mentions, Explain, Search, Compare, MarketSentiment, Stats, Health |
| `client.Polymarket()` | Trending, TrendingSectors, TrendingCountries, Stock, Mentions, Search, Compare, MarketSentiment, Stats, Health |
| `client.Sentiment()` | Analyze |
| `client.Health(ctx)` | Root health report |

`client.RedditCrypto()` is an alias for `client.Crypto()`.
The [package documentation](https://pkg.go.dev/github.com/adanos-software/adanos-go-sdk) lists all method signatures and response fields.

## Select a date window and filters

```go
news, err := client.News().Stock(ctx, "NU.V", &adanos.PeriodOptions{
	From: new("2026-10-01"),
	To:   new("2026-10-13"),
})

mentions, err := client.Crypto().Mentions(ctx, "BTC", &adanos.RedditMentionsOptions{
	Days:             new(7),
	Offset:           new(0),
	IncludeInherited: new(false),
})

trending, err := client.News().Trending(ctx, &adanos.NewsTrendingOptions{
	Source: new("Reuters"),
	Limit:  new(10),
})

comparison, err := client.X().Compare(ctx, []string{"AAPL", "MSFT"}, nil)
```

Nil options and nil fields omit query parameters. Pointers preserve explicit zero, false,
and empty strings. The server applies defaults and validates date windows and limits.
`Days` is the legacy parameter retained for compatibility with the other Adanos SDKs.
`Source` applies only to News trending endpoints. `IncludeInherited` applies only to
Reddit Stocks and Reddit Crypto mentions.

Ticker and symbol spelling remains unchanged, including News exchange suffixes such as `NU.V` and `BRK.A`.
Comparison lists are sent as one comma-separated query value.

## Analyze text

```go
analysis, err := client.Sentiment().Analyze(ctx,
	"Goldman Sachs cuts Apple price target",
	&adanos.SentimentAnalyzeOptions{Target: new("Apple")},
)
```

Pass nil options to omit the target. `Target: new("")` sends an explicit empty target.
Nullable component scores use pointers, so an unavailable model score differs from a score of zero.

## Inspect errors

```go
var apiErr *adanos.APIError
if errors.As(err, &apiErr) {
	fmt.Println(apiErr.StatusCode, apiErr.Detail)
	fmt.Println(apiErr.Headers.Get("Retry-After"))
	fmt.Println(apiErr.Headers.Get("X-Request-ID"))
}
if errors.Is(err, context.DeadlineExceeded) {
	fmt.Println("request deadline exceeded")
}
```

`APIError` retains the original response bytes in `Payload` and copies the response headers.
It also unwraps a response-body read error when one occurs. Invalid successful JSON, transport
failures, and context cancellation remain ordinary wrapped errors. A non-successful root health
response returns `APIError`; its health payload remains available in `Payload`.
The SDK does not retry requests automatically.

## Configure HTTP

```go
client, err := adanos.NewClient(adanos.Config{
	APIKey:     os.Getenv("ADANOS_API_KEY"),
	BaseURL:    "https://api.adanos.org",
	Timeout:    10 * time.Second,
	HTTPClient: &http.Client{Transport: http.DefaultTransport},
})
```

The default base URL is `https://api.adanos.com`. Explicit `.org` URLs and custom mount paths
remain effective. URLs must use HTTP or HTTPS and must not contain credentials, a query, or a fragment.
The default operation timeout is 30 seconds. A caller's context or injected HTTP client's shorter
timeout takes precedence. The injected client is copied and remains unchanged.
Cross-origin redirects are blocked before sending credentials, including when an injected redirect
callback changes the destination. Same-origin redirects retain the supplied callback.

Reuse the client concurrently. Custom transports and callbacks must support concurrent use;
do not change options or their pointed-to values while a call uses them.

## Response representation

Integers are signed `int64`. Optional and nullable scalars use pointers.
Missing and JSON null share nil. Nil slices represent missing or null arrays;
JSON `[]` produces a nonnil empty slice. Dates and timestamps remain strings with their original precision.
Enums remain strings and accept future values.

The three mixed error unions use `json.RawMessage`. Dynamic values inside root health services
also use `json.RawMessage`, which preserves large integers. Successful responses accept unknown fields
for forward compatibility. The decoder checks JSON syntax and declared scalar types, without validating
every required field or business constraint.

## Develop and regenerate

```sh
go generate ./...
go vet ./...
go test -count=1 ./...
go test -race -count=1 ./...
git diff --exit-code -- models.gen.go
```

The generator reads the frozen `api/openapi.json` and verifies its hash against `api/source.json`.
It sorts schemas and fields and fails on unsupported shapes or naming collisions.
The snapshot replaces the concrete key example at `/components/securitySchemes/ApiKeyAuth/description`
with `ADANOS_API_KEY`. The metadata records the original `source_sha256`, the stored `sha256`,
and that transformation. All paths and response schemas remain identical to the source contract.
Tests send every operation to an isolated HTTP server and audit every generated model property.
Set `ADANOS_TEST_ARTIFACT_DIR` to retain per-test JSON receipts. Without it, receipts go to
`adanos-go-sdk-tests` under the operating system's temporary directory.
CI runs Go 1.26 and 1.27, including race detection, model regeneration, and release-tag checks.

## License

[MIT](LICENSE).
