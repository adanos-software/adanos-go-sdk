package adanos_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	adanos "github.com/adanos-software/adanos-go-sdk"
)

func TestPolymarketGraphKeysAndNullableMetrics(t *testing.T) {
	var paths []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		switch r.URL.Path {
		case "/polymarket/stocks/v1/search":
			if r.URL.Query().Get("q") != "BK" {
				t.Errorf("search query %q", r.URL.RawQuery)
			}
			io.WriteString(w, `{"query":"BK","results":[{"ticker":"BNY"}]}`)
		case "/polymarket/stocks/v1/stock/SAP.DE":
			io.WriteString(w, `{"ticker":"SAP.DE","buzz_score":null}`)
		case "/polymarket/stocks/v1/stock/SAP.DE/mentions":
			io.WriteString(w, `{"ticker":"SAP.DE","results":[{"strike_price":null},{"strike_price":0}]}`)
		case "/polymarket/stocks/v1/compare":
			if r.URL.Query().Get("tickers") != "BNY,SAP.DE,EWG,UNKNOWN" {
				t.Errorf("compare query %q", r.URL.RawQuery)
			}
			io.WriteString(w, `{"period_days":7,"stocks":[{"ticker":"BNY","buzz_score":0},{"ticker":"UNKNOWN","buzz_score":null}]}`)
		default:
			t.Errorf("unexpected path %s", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()
	c := newTestClient(t, adanos.Config{BaseURL: srv.URL})
	search, err := c.Polymarket().Search(context.Background(), "BK", nil)
	if err != nil || len(search.Results) != 1 || search.Results[0].Ticker != "BNY" {
		t.Fatalf("search result %v error %v", search, err)
	}
	stock, err := c.Polymarket().Stock(context.Background(), "SAP.DE", nil)
	if err != nil || stock.Ticker != "SAP.DE" || stock.BuzzScore != nil {
		t.Fatalf("stock result %v error %v", stock, err)
	}
	mentions, err := c.Polymarket().Mentions(context.Background(), "SAP.DE", nil)
	if err != nil || len(mentions.Results) != 2 || mentions.Results[0].StrikePrice != nil || mentions.Results[1].StrikePrice == nil || *mentions.Results[1].StrikePrice != 0 {
		t.Fatalf("mentions result %v error %v", mentions, err)
	}
	comparison, err := c.Polymarket().Compare(context.Background(), []string{"BNY", "SAP.DE", "EWG", "UNKNOWN"}, nil)
	if err != nil || len(comparison.Stocks) != 2 || comparison.Stocks[0].BuzzScore == nil || *comparison.Stocks[0].BuzzScore != 0 || comparison.Stocks[1].BuzzScore != nil || comparison.Stocks[1].CompanyName != nil {
		t.Fatalf("compare result %v error %v", comparison, err)
	}
	want := []string{"/polymarket/stocks/v1/search", "/polymarket/stocks/v1/stock/SAP.DE", "/polymarket/stocks/v1/stock/SAP.DE/mentions", "/polymarket/stocks/v1/compare"}
	if !reflect.DeepEqual(paths, want) {
		t.Fatalf("paths %v, want %v", paths, want)
	}
	receipt(t, map[string]any{"search": "BK", "ticker": "SAP.DE", "compare": []string{"BNY", "SAP.DE", "EWG", "UNKNOWN"}}, map[string]any{"paths": paths, "search": search, "stock": stock, "mentions": mentions, "compare": comparison})
}

func TestPolymarketGraphErrors(t *testing.T) {
	for _, tc := range []struct {
		ticker string
		status int
	}{{"BK", 404}, {"BRK.B", 422}, {"BRK-B", 422}, {"$AAPL", 422}, {"BRK", 503}} {
		t.Run(fmt.Sprintf("%s/%d", tc.ticker, tc.status), func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/polymarket/stocks/v1/stock/"+tc.ticker+"/mentions" {
					t.Errorf("ticker rewritten in %s", r.URL.Path)
				}
				w.Header().Set("X-Request-ID", "graph-error")
				w.WriteHeader(tc.status)
				io.WriteString(w, `{"detail":"Graph lookup failed"}`)
			}))
			defer srv.Close()
			out, err := newTestClient(t, adanos.Config{BaseURL: srv.URL}).Polymarket().Mentions(context.Background(), tc.ticker, nil)
			var apiErr *adanos.APIError
			if out != nil || !errors.As(err, &apiErr) || apiErr.StatusCode != tc.status || apiErr.Detail != "Graph lookup failed" || apiErr.Headers.Get("X-Request-ID") != "graph-error" {
				t.Fatalf("result %v error %v", out, err)
			}
			receipt(t, map[string]any{"ticker": tc.ticker, "status": tc.status}, map[string]any{"status": apiErr.StatusCode, "detail": apiErr.Detail})
		})
	}
}
