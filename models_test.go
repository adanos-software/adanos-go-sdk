package adanos_test

import (
	"encoding/json"
	adanos "github.com/adanos-software/adanos-go-sdk"
	"os"
	"reflect"
	"strings"
	"testing"
)

func modelRegistry() map[string]reflect.Type {
	return map[string]reflect.Type{
		"CompareLimitError":                  reflect.TypeFor[adanos.CompareLimitError](),
		"CompareLimitErrorDetail":            reflect.TypeFor[adanos.CompareLimitErrorDetail](),
		"CompareResponse":                    reflect.TypeFor[adanos.CompareResponse](),
		"CompareStockItem":                   reflect.TypeFor[adanos.CompareStockItem](),
		"CryptoCompareResponse":              reflect.TypeFor[adanos.CryptoCompareResponse](),
		"CryptoCompareTokenItem":             reflect.TypeFor[adanos.CryptoCompareTokenItem](),
		"CryptoDailyTrendItem":               reflect.TypeFor[adanos.CryptoDailyTrendItem](),
		"CryptoHealthResponse":               reflect.TypeFor[adanos.CryptoHealthResponse](),
		"CryptoMarketSentimentDriver":        reflect.TypeFor[adanos.CryptoMarketSentimentDriver](),
		"CryptoMarketSentimentResponse":      reflect.TypeFor[adanos.CryptoMarketSentimentResponse](),
		"CryptoRawMentionItem":               reflect.TypeFor[adanos.CryptoRawMentionItem](),
		"CryptoRawMentionsResponse":          reflect.TypeFor[adanos.CryptoRawMentionsResponse](),
		"CryptoSchedulerStatus":              reflect.TypeFor[adanos.CryptoSchedulerStatus](),
		"CryptoSearchItem":                   reflect.TypeFor[adanos.CryptoSearchItem](),
		"CryptoSearchResponse":               reflect.TypeFor[adanos.CryptoSearchResponse](),
		"CryptoSearchSummary":                reflect.TypeFor[adanos.CryptoSearchSummary](),
		"CryptoStatsResponse":                reflect.TypeFor[adanos.CryptoStatsResponse](),
		"CryptoSubredditCount":               reflect.TypeFor[adanos.CryptoSubredditCount](),
		"CryptoTokenSentiment":               reflect.TypeFor[adanos.CryptoTokenSentiment](),
		"CryptoTopMention":                   reflect.TypeFor[adanos.CryptoTopMention](),
		"CryptoTrendingToken":                reflect.TypeFor[adanos.CryptoTrendingToken](),
		"DailyTrendItem":                     reflect.TypeFor[adanos.DailyTrendItem](),
		"ErrorResponse":                      reflect.TypeFor[adanos.ErrorResponse](),
		"HTTPValidationError":                reflect.TypeFor[adanos.HTTPValidationError](),
		"HealthResponse":                     reflect.TypeFor[adanos.HealthResponse](),
		"HistoricalLimitError":               reflect.TypeFor[adanos.HistoricalLimitError](),
		"HistoricalLimitErrorDetail":         reflect.TypeFor[adanos.HistoricalLimitErrorDetail](),
		"InvalidPeriodError":                 reflect.TypeFor[adanos.InvalidPeriodError](),
		"InvalidPeriodErrorDetail":           reflect.TypeFor[adanos.InvalidPeriodErrorDetail](),
		"InvalidSearchParametersError":       reflect.TypeFor[adanos.InvalidSearchParametersError](),
		"InvalidSearchParametersErrorDetail": reflect.TypeFor[adanos.InvalidSearchParametersErrorDetail](),
		"NewsCompareResponse":                reflect.TypeFor[adanos.NewsCompareResponse](),
		"NewsCompareStockItem":               reflect.TypeFor[adanos.NewsCompareStockItem](),
		"NewsMarketSentimentDriver":          reflect.TypeFor[adanos.NewsMarketSentimentDriver](),
		"NewsMarketSentimentResponse":        reflect.TypeFor[adanos.NewsMarketSentimentResponse](),
		"NewsRawMentionItem":                 reflect.TypeFor[adanos.NewsRawMentionItem](),
		"NewsRawMentionsResponse":            reflect.TypeFor[adanos.NewsRawMentionsResponse](),
		"NewsSearchResponse":                 reflect.TypeFor[adanos.NewsSearchResponse](),
		"NewsSearchResultItem":               reflect.TypeFor[adanos.NewsSearchResultItem](),
		"NewsSearchSummary":                  reflect.TypeFor[adanos.NewsSearchSummary](),
		"NewsSourceCount":                    reflect.TypeFor[adanos.NewsSourceCount](),
		"NewsStatsResponse":                  reflect.TypeFor[adanos.NewsStatsResponse](),
		"NewsStockSentiment":                 reflect.TypeFor[adanos.NewsStockSentiment](),
		"NewsTopMention":                     reflect.TypeFor[adanos.NewsTopMention](),
		"NewsTrendingCountry":                reflect.TypeFor[adanos.NewsTrendingCountry](),
		"NewsTrendingSector":                 reflect.TypeFor[adanos.NewsTrendingSector](),
		"NewsTrendingStock":                  reflect.TypeFor[adanos.NewsTrendingStock](),
		"PolymarketCompareResponse":          reflect.TypeFor[adanos.PolymarketCompareResponse](),
		"PolymarketCompareStockItem":         reflect.TypeFor[adanos.PolymarketCompareStockItem](),
		"PolymarketDailyTrendItem":           reflect.TypeFor[adanos.PolymarketDailyTrendItem](),
		"PolymarketHealthResponse":           reflect.TypeFor[adanos.PolymarketHealthResponse](),
		"PolymarketMarketSentimentDriver":    reflect.TypeFor[adanos.PolymarketMarketSentimentDriver](),
		"PolymarketMarketSentimentResponse":  reflect.TypeFor[adanos.PolymarketMarketSentimentResponse](),
		"PolymarketPulseEvidence":            reflect.TypeFor[adanos.PolymarketPulseEvidence](),
		"PolymarketRawMentionItem":           reflect.TypeFor[adanos.PolymarketRawMentionItem](),
		"PolymarketRawMentionsResponse":      reflect.TypeFor[adanos.PolymarketRawMentionsResponse](),
		"PolymarketSchedulerStatus":          reflect.TypeFor[adanos.PolymarketSchedulerStatus](),
		"PolymarketSearchResponse":           reflect.TypeFor[adanos.PolymarketSearchResponse](),
		"PolymarketSearchResultItem":         reflect.TypeFor[adanos.PolymarketSearchResultItem](),
		"PolymarketSearchSummary":            reflect.TypeFor[adanos.PolymarketSearchSummary](),
		"PolymarketStatsResponse":            reflect.TypeFor[adanos.PolymarketStatsResponse](),
		"PolymarketStockDetailResponse":      reflect.TypeFor[adanos.PolymarketStockDetailResponse](),
		"PolymarketTickerPulse":              reflect.TypeFor[adanos.PolymarketTickerPulse](),
		"PolymarketTopMention":               reflect.TypeFor[adanos.PolymarketTopMention](),
		"PolymarketTrendingCountry":          reflect.TypeFor[adanos.PolymarketTrendingCountry](),
		"PolymarketTrendingSector":           reflect.TypeFor[adanos.PolymarketTrendingSector](),
		"PolymarketTrendingStock":            reflect.TypeFor[adanos.PolymarketTrendingStock](),
		"RateLimitError":                     reflect.TypeFor[adanos.RateLimitError](),
		"RateLimitErrorDetail":               reflect.TypeFor[adanos.RateLimitErrorDetail](),
		"RedditMarketSentimentDriver":        reflect.TypeFor[adanos.RedditMarketSentimentDriver](),
		"RedditMarketSentimentResponse":      reflect.TypeFor[adanos.RedditMarketSentimentResponse](),
		"RedditRawMentionItem":               reflect.TypeFor[adanos.RedditRawMentionItem](),
		"RedditRawMentionsResponse":          reflect.TypeFor[adanos.RedditRawMentionsResponse](),
		"RedditStatsResponse":                reflect.TypeFor[adanos.RedditStatsResponse](),
		"RootHealthResponse":                 reflect.TypeFor[adanos.RootHealthResponse](),
		"RootHealthSummary":                  reflect.TypeFor[adanos.RootHealthSummary](),
		"SchedulerStatus":                    reflect.TypeFor[adanos.SchedulerStatus](),
		"SearchResponse":                     reflect.TypeFor[adanos.SearchResponse](),
		"SearchResultItem":                   reflect.TypeFor[adanos.SearchResultItem](),
		"SentimentAnalyzeRequest":            reflect.TypeFor[adanos.SentimentAnalyzeRequest](),
		"SentimentAnalyzeResponse":           reflect.TypeFor[adanos.SentimentAnalyzeResponse](),
		"SentimentComponents":                reflect.TypeFor[adanos.SentimentComponents](),
		"SentimentPhraseMatch":               reflect.TypeFor[adanos.SentimentPhraseMatch](),
		"StockExplanationResponse":           reflect.TypeFor[adanos.StockExplanationResponse](),
		"StockSearchSummary":                 reflect.TypeFor[adanos.StockSearchSummary](),
		"StockSentiment":                     reflect.TypeFor[adanos.StockSentiment](),
		"SubredditCount":                     reflect.TypeFor[adanos.SubredditCount](),
		"TopMention":                         reflect.TypeFor[adanos.TopMention](),
		"TrendingCountry":                    reflect.TypeFor[adanos.TrendingCountry](),
		"TrendingSector":                     reflect.TypeFor[adanos.TrendingSector](),
		"TrendingStock":                      reflect.TypeFor[adanos.TrendingStock](),
		"UnsupportedAssetError":              reflect.TypeFor[adanos.UnsupportedAssetError](),
		"UnsupportedAssetErrorDetail":        reflect.TypeFor[adanos.UnsupportedAssetErrorDetail](),
		"ValidationError":                    reflect.TypeFor[adanos.ValidationError](),
		"XCompareResponse":                   reflect.TypeFor[adanos.XCompareResponse](),
		"XCompareStockItem":                  reflect.TypeFor[adanos.XCompareStockItem](),
		"XDailyTrendItem":                    reflect.TypeFor[adanos.XDailyTrendItem](),
		"XHealthResponse":                    reflect.TypeFor[adanos.XHealthResponse](),
		"XMarketSentimentDriver":             reflect.TypeFor[adanos.XMarketSentimentDriver](),
		"XMarketSentimentResponse":           reflect.TypeFor[adanos.XMarketSentimentResponse](),
		"XRawMentionItem":                    reflect.TypeFor[adanos.XRawMentionItem](),
		"XRawMentionsResponse":               reflect.TypeFor[adanos.XRawMentionsResponse](),
		"XSchedulerStatus":                   reflect.TypeFor[adanos.XSchedulerStatus](),
		"XSearchResponse":                    reflect.TypeFor[adanos.XSearchResponse](),
		"XSearchResultItem":                  reflect.TypeFor[adanos.XSearchResultItem](),
		"XSearchSummary":                     reflect.TypeFor[adanos.XSearchSummary](),
		"XStatsResponse":                     reflect.TypeFor[adanos.XStatsResponse](),
		"XStockDetailResponse":               reflect.TypeFor[adanos.XStockDetailResponse](),
		"XTopAuthor":                         reflect.TypeFor[adanos.XTopAuthor](),
		"XTopTweet":                          reflect.TypeFor[adanos.XTopTweet](),
		"XTrendingCountry":                   reflect.TypeFor[adanos.XTrendingCountry](),
		"XTrendingSector":                    reflect.TypeFor[adanos.XTrendingSector](),
		"XTrendingStock":                     reflect.TypeFor[adanos.XTrendingStock]()}
}

func TestAllModelPropertiesMatchContract(t *testing.T) {
	data, err := os.ReadFile("api/openapi.json")
	if err != nil {
		t.Fatal(err)
	}
	var spec struct {
		Components struct {
			Schemas map[string]map[string]any `json:"schemas"`
		} `json:"components"`
	}
	if err := json.Unmarshal(data, &spec); err != nil {
		t.Fatal(err)
	}
	registry := modelRegistry()
	count := 0
	for name, schema := range spec.Components.Schemas {
		typ, ok := registry[name]
		if !ok {
			t.Errorf("no public model %s", name)
			continue
		}
		props := schema["properties"].(map[string]any)
		required := map[string]bool{}
		if r, ok := schema["required"].([]any); ok {
			for _, v := range r {
				required[v.(string)] = true
			}
		}
		fields := map[string]reflect.Type{}
		for i := range typ.NumField() {
			f := typ.Field(i)
			fields[strings.Split(f.Tag.Get("json"), ",")[0]] = f.Type
		}
		if len(fields) != len(props) {
			t.Errorf("%s has %d fields, contract %d", name, len(fields), len(props))
		}
		for key, property := range props {
			count++
			f, ok := fields[key]
			if !ok {
				t.Errorf("%s.%s missing", name, key)
				continue
			}
			checkModelShape(t, name+"."+key, property.(map[string]any), f, !required[key], registry)
		}
	}
	if len(registry) != len(spec.Components.Schemas) {
		t.Fatalf("models %d, schemas %d", len(registry), len(spec.Components.Schemas))
	}
	receipt(t, map[string]any{"contract": "api/openapi.json"}, map[string]any{"schemas": len(registry), "properties": count})
}
func checkModelShape(t *testing.T, path string, s map[string]any, typ reflect.Type, optional bool, registry map[string]reflect.Type) {
	t.Helper()
	if variants, ok := s["anyOf"].([]any); ok {
		var real []map[string]any
		nullable := false
		for _, v := range variants {
			m := v.(map[string]any)
			if m["type"] == "null" {
				nullable = true
			} else {
				real = append(real, m)
			}
		}
		if len(real) > 1 {
			if typ != reflect.TypeFor[json.RawMessage]() {
				t.Errorf("%s union type %s", path, typ)
			}
			return
		}
		checkModelShape(t, path, real[0], typ, optional || nullable, registry)
		return
	}
	kind, _ := s["type"].(string)
	if optional && kind != "array" && kind != "object" {
		if typ.Kind() != reflect.Pointer {
			t.Errorf("%s optional type %s", path, typ)
			return
		}
		typ = typ.Elem()
	}
	if ref, ok := s["$ref"].(string); ok {
		want := registry[strings.TrimPrefix(ref, "#/components/schemas/")]
		if typ != want {
			t.Errorf("%s ref type %s want %s", path, typ, want)
		}
		return
	}
	switch kind {
	case "string":
		if typ.Kind() != reflect.String {
			t.Errorf("%s string type %s", path, typ)
		}
	case "integer":
		if typ != reflect.TypeFor[int64]() {
			t.Errorf("%s signed integer type %s", path, typ)
		}
	case "number":
		if typ != reflect.TypeFor[float64]() {
			t.Errorf("%s number type %s", path, typ)
		}
	case "boolean":
		if typ.Kind() != reflect.Bool {
			t.Errorf("%s bool type %s", path, typ)
		}
	case "array":
		if typ.Kind() != reflect.Slice {
			t.Errorf("%s array type %s", path, typ)
			return
		}
		checkModelShape(t, path+"[]", s["items"].(map[string]any), typ.Elem(), false, registry)
	case "object":
		if typ.Kind() != reflect.Map || typ.Key().Kind() != reflect.String {
			t.Errorf("%s map type %s", path, typ)
			return
		}
		if a, ok := s["additionalProperties"].(map[string]any); ok {
			checkModelShape(t, path+"{}", a, typ.Elem(), false, registry)
		} else if typ.Elem() != reflect.TypeFor[json.RawMessage]() {
			t.Errorf("%s dynamic JSON type %s", path, typ.Elem())
		}
	default:
		t.Errorf("unsupported independent audit shape %s %v", path, s)
	}
}
func TestModelWireSemantics(t *testing.T) {
	var crypto adanos.CryptoRawMentionsResponse
	raw := `{"symbol":"BTC","results":[{"upvotes":-9,"created_utc":"2026-10-13T15:02:00.123456+00:00","is_inherited":false,"sentiment_score":null},{"upvotes":0,"sentiment_score":0}]}`
	if err := json.Unmarshal([]byte(raw), &crypto); err != nil {
		t.Fatal(err)
	}
	if crypto.Results[0].Upvotes != -9 || crypto.Results[0].CreatedUTC != "2026-10-13T15:02:00.123456+00:00" || crypto.Results[0].SentimentScore != nil || crypto.Results[1].SentimentScore == nil || *crypto.Results[1].SentimentScore != 0 {
		t.Fatalf("crypto %#v", crypto.Results)
	}
	var sentiment adanos.SentimentAnalyzeResponse
	if err := json.Unmarshal([]byte(`{"components":{"model_score":null,"vader_compound":0,"roberta_score":null}}`), &sentiment); err != nil {
		t.Fatal(err)
	}
	if sentiment.Components.ModelScore != nil || sentiment.Components.VaderCompound == nil || *sentiment.Components.VaderCompound != 0 || sentiment.Components.RobertaScore != nil {
		t.Fatalf("components %#v", sentiment.Components)
	}
	var poly adanos.PolymarketStockDetailResponse
	if err := json.Unmarshal([]byte(`{"ticker":"AAPL","found":false,"unique_traders":null,"sentiment_score":0,"daily_trend":[],"top_mentions":null,"pulse":{"evidence":{"avg_spread":null,"snapshot_at":"2026-10-13T00:00:00Z"}}}`), &poly); err != nil {
		t.Fatal(err)
	}
	if poly.Found || poly.UniqueTraders != nil || poly.SentimentScore == nil || *poly.SentimentScore != 0 || poly.DailyTrend == nil || len(poly.DailyTrend) != 0 || poly.TopMentions != nil || poly.Pulse.Evidence.AvgSpread != nil || *poly.Pulse.Evidence.SnapshotAt != "2026-10-13T00:00:00Z" {
		t.Fatalf("polymarket %#v", poly)
	}
	var health adanos.RootHealthResponse
	if err := json.Unmarshal([]byte(`{"services":{"news":{"large":9007199254740993,"array":[false,null],"object":{"x":"y"}}}}`), &health); err != nil {
		t.Fatal(err)
	}
	if string(health.Services["news"]["large"]) != "9007199254740993" || string(health.Services["news"]["array"]) != "[false,null]" {
		t.Fatalf("health %#v", health)
	}
	for _, raw := range []string{`{"value":9007199254740993}`, `{"value":"2026-10-13"}`, `{"value":null}`} {
		var v adanos.InvalidPeriodErrorDetail
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		out, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		var got, want map[string]json.RawMessage
		json.Unmarshal(out, &got)
		json.Unmarshal([]byte(raw), &want)
		if string(got["value"]) != string(want["value"]) {
			t.Fatalf("union %s -> %s", raw, out)
		}
	}
	for _, raw := range []string{`{"detail":"invalid search"}`, `{"detail":{"message":"invalid parameter"}}`} {
		var v adanos.InvalidSearchParametersError
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		var in map[string]json.RawMessage
		json.Unmarshal([]byte(raw), &in)
		if string(v.Detail) != string(in["detail"]) {
			t.Fatalf("detail %s", v.Detail)
		}
	}
	var validation adanos.ValidationError
	if err := json.Unmarshal([]byte(`{"loc":["body",9007199254740993]}`), &validation); err != nil {
		t.Fatal(err)
	}
	if string(validation.Loc[1]) != "9007199254740993" {
		t.Fatal("location lost precision")
	}
	receipt(t, raw, map[string]any{"crypto": crypto, "polymarket": poly, "sentiment": sentiment, "health": health, "location": validation})
}
