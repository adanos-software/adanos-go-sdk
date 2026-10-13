package adanos_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	adanos "github.com/adanos-software/adanos-go-sdk"
)

type operationCase struct {
	name, method, path, template, fixture, field string
	want                                         any
	query                                        url.Values
	body                                         string
	call                                         func(context.Context, *adanos.Client) (any, error)
}

func receipt(t *testing.T, input, result any) {
	t.Helper()
	dir := os.Getenv("ADANOS_TEST_ARTIFACT_DIR")
	if dir == "" {
		dir = filepath.Join(os.TempDir(), "adanos-go-sdk-tests")
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	data, err := json.MarshalIndent(map[string]any{"test": t.Name(), "command": "go test -count=1 ./...", "input": input, "result": result}, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, strings.ReplaceAll(t.Name(), "/", "_")+".json"), data, 0644); err != nil {
		t.Fatal(err)
	}
}

func newTestClient(t *testing.T, config adanos.Config) *adanos.Client {
	t.Helper()
	c, err := adanos.NewClient(config)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestAllOperations(t *testing.T) {
	cases := operationCases()
	seen := map[string]bool{}
	for _, tc := range cases {
		key := strings.ToLower(tc.method) + " " + tc.template
		if seen[key] {
			t.Fatalf("duplicate operation %s", key)
		}
		seen[key] = true
		t.Run(tc.name, func(t *testing.T) {
			var observed map[string]any
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body, err := io.ReadAll(r.Body)
				if err != nil {
					t.Error(err)
				}
				if r.Method != tc.method || r.URL.EscapedPath() != "/mount"+tc.path {
					t.Errorf("request %s %s, want %s /mount%s", r.Method, r.URL.EscapedPath(), tc.method, tc.path)
				}
				if !reflect.DeepEqual(r.URL.Query(), tc.query) {
					t.Errorf("query %v, want %v", r.URL.Query(), tc.query)
				}
				if r.Header.Get("X-API-Key") != "synthetic-test-key" || r.Header.Get("Accept") != "application/json" {
					t.Errorf("missing auth/accept headers")
				}
				if tc.method == "POST" && r.Header.Get("Content-Type") != "application/json" {
					t.Error("missing content type")
				}
				assertJSONEqual(t, string(body), tc.body)
				observed = map[string]any{"method": r.Method, "uri": r.RequestURI, "body": string(body)}
				w.Header().Set("Content-Type", "application/json")
				io.WriteString(w, tc.fixture)
			}))
			defer srv.Close()
			c := newTestClient(t, adanos.Config{APIKey: "synthetic-test-key", BaseURL: srv.URL + "/mount///"})
			out, err := tc.call(context.Background(), c)
			if err != nil {
				t.Fatal(err)
			}
			data, err := json.Marshal(out)
			if err != nil {
				t.Fatal(err)
			}
			var result any
			if err := json.Unmarshal(data, &result); err != nil {
				t.Fatal(err)
			}
			if a, ok := result.([]any); ok {
				if len(a) != 1 {
					t.Fatalf("array %s", data)
				}
				result = a[0]
			}
			if got := result.(map[string]any)[tc.field]; !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("decoded %s=%v, want %v", tc.field, got, tc.want)
			}
			sum := sha256.Sum256([]byte(tc.fixture))
			receipt(t, map[string]any{"operation": key, "fixture": tc.fixture, "fixture_sha256": hex.EncodeToString(sum[:])}, map[string]any{"request": observed, "decoded": result})
		})
	}
	data, err := os.ReadFile("api/openapi.json")
	if err != nil {
		t.Fatal(err)
	}
	var spec struct {
		Paths map[string]map[string]json.RawMessage `json:"paths"`
	}
	if err := json.Unmarshal(data, &spec); err != nil {
		t.Fatal(err)
	}
	count := 0
	for path, methods := range spec.Paths {
		for method := range methods {
			count++
			if !seen[method+" "+path] {
				t.Errorf("uncovered %s %s", method, path)
			}
		}
	}
	if count != len(cases) {
		t.Fatalf("tested %d operations, contract has %d", len(cases), count)
	}
	if c := newTestClient(t, adanos.Config{}); c.Crypto() != c.RedditCrypto() {
		t.Fatal("crypto alias must share the service")
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestConfigurationAndEscapedIdentifiers(t *testing.T) {
	for _, base := range []string{"ftp://example.com", "/relative", "http://", "https://u:p@example.com", "https://example.com?x=1", "https://example.com#fragment"} {
		if _, err := adanos.NewClient(adanos.Config{BaseURL: base}); err == nil {
			t.Errorf("accepted %q", base)
		}
	}
	for _, config := range []adanos.Config{{APIKey: "bad\rkey"}, {APIKey: "bad\nkey"}, {Timeout: -time.Second}} {
		if _, err := adanos.NewClient(config); err == nil {
			t.Error("accepted invalid config")
		}
	}
	for _, base := range []string{"", "https://api.adanos.org", "https://example.com/a%2Fb/"} {
		want := base + "/health"
		if base == "" {
			want = "https://api.adanos.com/health"
		}
		if strings.HasSuffix(base, "/") {
			want = base + "health"
		}
		client := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			if r.URL.String() != want {
				t.Errorf("URL %q, want %q", r.URL, want)
			}
			return &http.Response{StatusCode: 200, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(`{"status":"healthy"}`))}, nil
		})}
		out, err := newTestClient(t, adanos.Config{BaseURL: base, HTTPClient: client}).Health(context.Background())
		if err != nil || out.Status != "healthy" {
			t.Fatalf("health %v %v", out, err)
		}
	}
	for _, tc := range []struct{ in, path string }{{"NU.V", "NU.V"}, {"BRK.A", "BRK.A"}, {"a/b%? +界", "a%2Fb%25%3F%20+%E7%95%8C"}, {".", "%2E"}, {"..", "%2E%2E"}} {
		t.Run(tc.in, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.EscapedPath() != "/news/stocks/v1/stock/"+tc.path {
					t.Errorf("escaped %s", r.URL.EscapedPath())
				}
				fmt.Fprintf(w, `{"ticker":%q}`, tc.in)
			}))
			defer srv.Close()
			out, err := newTestClient(t, adanos.Config{BaseURL: srv.URL}).News().Stock(context.Background(), tc.in, nil)
			if err != nil || out.Ticker != tc.in {
				t.Fatalf("result %v %v", out, err)
			}
			receipt(t, tc.in, out)
		})
	}
}

func TestOptionalQueryAndSentimentTargets(t *testing.T) {
	var requests []map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		requests = append(requests, map[string]any{"query": r.URL.Query(), "body": jsonValue(t, string(data))})
		if r.Method == "POST" {
			io.WriteString(w, `{"text":"sample","sentiment_score":0.25}`)
		} else {
			io.WriteString(w, `{"ticker":"NU.V"}`)
		}
	}))
	defer srv.Close()
	c := newTestClient(t, adanos.Config{BaseURL: srv.URL})
	if _, err := c.News().Stock(context.Background(), "NU.V", nil); err != nil {
		t.Fatal(err)
	}
	for _, o := range []*adanos.SentimentAnalyzeOptions{nil, {Target: new("")}, {Target: new("Apple")}} {
		out, err := c.Sentiment().Analyze(context.Background(), "sample", o)
		if err != nil || out.SentimentScore != 0.25 {
			t.Fatalf("analysis %v %v", out, err)
		}
	}
	want := []map[string]any{{"query": url.Values{}, "body": nil}, {"query": url.Values{}, "body": map[string]any{"text": "sample"}}, {"query": url.Values{}, "body": map[string]any{"text": "sample", "target": ""}}, {"query": url.Values{}, "body": map[string]any{"text": "sample", "target": "Apple"}}}
	if !reflect.DeepEqual(requests, want) {
		t.Fatalf("requests %#v", requests)
	}
	receipt(t, "nil/empty/named options", requests)
}

func TestErrorPayloadsAndInvalidSuccessJSON(t *testing.T) {
	for _, tc := range []struct {
		status       int
		body, detail string
	}{{400, `{"detail":{"error":"bad_period","message":"invalid dates"}}`, "invalid dates"}, {401, `{"detail":"key rejected"}`, "key rejected"}, {403, `{"detail":{"error":"forbidden"}}`, "forbidden"}, {404, `{"detail":"missing"}`, "missing"}, {422, `{"detail":[{"loc":["query","limit"],"msg":"too small","type":"greater_than"}]}`, "query.limit: too small"}, {429, `{"detail":{"message":"quota exhausted"}}`, "quota exhausted"}, {500, "<html>error</html>", "Internal Server Error"}, {503, `{"status":"unhealthy"}`, `{"status":"unhealthy"}`}} {
		t.Run(fmt.Sprint(tc.status), func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Retry-After", "12")
				w.Header().Set("X-Request-ID", "synthetic-id")
				w.Header().Set("X-RateLimit-Remaining", "0")
				w.WriteHeader(tc.status)
				io.WriteString(w, tc.body)
			}))
			defer srv.Close()
			out, err := newTestClient(t, adanos.Config{BaseURL: srv.URL}).Health(context.Background())
			var apiErr *adanos.APIError
			if out != nil || !errors.As(err, &apiErr) {
				t.Fatalf("result %v error %v", out, err)
			}
			if apiErr.StatusCode != tc.status || apiErr.Detail != tc.detail || string(apiErr.Payload) != tc.body || apiErr.Headers.Get("Retry-After") != "12" || apiErr.Headers.Get("X-Request-ID") != "synthetic-id" || apiErr.Headers.Get("X-RateLimit-Remaining") != "0" {
				t.Fatalf("APIError %#v", apiErr)
			}
			receipt(t, tc.body, map[string]any{"status": apiErr.StatusCode, "detail": apiErr.Detail})
		})
	}
	for _, body := range []string{"", "{", `{"status":"healthy"} {"status":"other"}`, `{"status":"healthy"}junk`, `{"status":7}`} {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, body) }))
		out, err := newTestClient(t, adanos.Config{BaseURL: srv.URL}).Health(context.Background())
		srv.Close()
		if err == nil || out != nil {
			t.Fatalf("accepted %q, result %v", body, out)
		}
	}
}

type failingBody struct {
	closed atomic.Bool
	fail   error
}

func (b *failingBody) Read(p []byte) (int, error) {
	copy(p, `{"detail":"partial"}`)
	return len(`{"detail":"partial"}`), b.fail
}
func (b *failingBody) Close() error { b.closed.Store(true); return nil }

func TestResponseBodyClosureAndReadErrors(t *testing.T) {
	sentinel := errors.New("synthetic read failure")
	for _, status := range []int{200, 429} {
		b := &failingBody{fail: sentinel}
		headers := http.Header{"X-Request-Id": []string{"original"}}
		hc := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: status, Header: headers, Body: b}, nil
		})}
		out, err := newTestClient(t, adanos.Config{HTTPClient: hc}).Health(context.Background())
		if out != nil || !errors.Is(err, sentinel) || !b.closed.Load() {
			t.Fatalf("status=%d out=%v err=%v closed=%v", status, out, err, b.closed.Load())
		}
		if status == 429 {
			var apiErr *adanos.APIError
			if !errors.As(err, &apiErr) || apiErr.StatusCode != 429 || string(apiErr.Payload) != `{"detail":"partial"}` {
				t.Fatalf("error %v", err)
			}
			headers.Set("X-Request-Id", "mutated")
			if apiErr.Headers.Get("X-Request-ID") != "original" {
				t.Fatal("headers not copied")
			}
		}
		receipt(t, status, err.Error())
	}
}

func TestCancellationAndTimeouts(t *testing.T) {
	for _, mode := range []string{"before", "during", "caller-deadline", "SDK-timeout", "client-timeout"} {
		t.Run(mode, func(t *testing.T) {
			arrived := make(chan struct{})
			canceled := make(chan struct{})
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { close(arrived); <-r.Context().Done(); close(canceled) }))
			defer srv.Close()
			config := adanos.Config{BaseURL: srv.URL}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			want := context.Canceled
			if mode == "before" {
				cancel()
			}
			if mode == "caller-deadline" {
				ctx, cancel = context.WithTimeout(context.Background(), 30*time.Millisecond)
				defer cancel()
				want = context.DeadlineExceeded
			}
			if mode == "SDK-timeout" {
				config.Timeout = 30 * time.Millisecond
				want = context.DeadlineExceeded
			}
			if mode == "client-timeout" {
				config.HTTPClient = &http.Client{Timeout: 30 * time.Millisecond}
				want = context.DeadlineExceeded
			}
			if mode == "during" {
				go func() { <-arrived; cancel() }()
			}
			out, err := newTestClient(t, config).Health(ctx)
			if out != nil || !errors.Is(err, want) {
				t.Fatalf("out=%v err=%v want=%v", out, err, want)
			}
			if mode != "before" {
				select {
				case <-canceled:
				case <-time.After(time.Second):
					t.Fatal("server did not observe cancellation")
				}
			}
			receipt(t, mode, err.Error())
		})
	}
}

func TestRedirectsAndInjectedClient(t *testing.T) {
	var foreignRequests atomic.Int64
	foreign := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		foreignRequests.Add(1)
		t.Errorf("foreign server received request with X-API-Key %q", r.Header.Get("X-API-Key"))
		io.WriteString(w, `{"status":"leaked"}`)
	}))
	defer foreign.Close()
	for _, mode := range []string{"foreign", "same-origin", "callback-rewrite", "callback-origin-rewrite", "loop"} {
		t.Run(mode, func(t *testing.T) {
			var callbacks atomic.Int64
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/done" {
					io.WriteString(w, `{"status":"healthy"}`)
					return
				}
				target := "/done"
				if mode == "foreign" {
					target = foreign.URL + "/health"
				}
				if mode == "loop" {
					target = "/health"
				}
				http.Redirect(w, r, target, http.StatusFound)
			}))
			defer srv.Close()
			hc := &http.Client{Timeout: time.Second}
			var callback func(*http.Request, []*http.Request) error
			if mode != "loop" {
				callback = func(r *http.Request, via []*http.Request) error {
					callbacks.Add(1)
					if mode == "callback-rewrite" || mode == "callback-origin-rewrite" {
						r.URL, _ = url.Parse(foreign.URL + "/health")
						if mode == "callback-origin-rewrite" {
							*via[0].URL = *r.URL
						}
					}
					return nil
				}
				hc.CheckRedirect = callback
			}
			out, err := newTestClient(t, adanos.Config{BaseURL: srv.URL, HTTPClient: hc, APIKey: "synthetic-test-key"}).Health(context.Background())
			if mode == "same-origin" {
				if err != nil || out.Status != "healthy" || callbacks.Load() != 1 {
					t.Fatalf("same redirect %v %v callbacks=%d", out, err, callbacks.Load())
				}
			} else {
				if err == nil || out != nil {
					t.Fatalf("allowed redirect %v %v", out, err)
				}
			}
			if hc.Timeout != time.Second || reflect.ValueOf(hc.CheckRedirect).Pointer() != reflect.ValueOf(callback).Pointer() {
				t.Fatal("injected client mutated")
			}
			if mode == "foreign" && callbacks.Load() != 0 {
				t.Fatal("foreign callback invoked")
			}
			if foreignRequests.Load() != 0 {
				t.Fatal("key crossed origin")
			}
			receipt(t, mode, map[string]any{"foreign_requests": foreignRequests.Load(), "callbacks": callbacks.Load()})
		})
	}
}

func TestRedirectHostThroughProxy(t *testing.T) {
	for _, mode := range []string{"foreign-host", "foreign-opaque", "foreign-opaque-authority", "foreign-path", "foreign-rawpath", "foreign-connect-authority", "same-origin-host", "empty-host", "same-origin-opaque-authority", "same-origin-opaque-absolute", "same-origin-opaque-path", "same-origin-escaped-path", "same-origin-query-fragment"} {
		t.Run(mode, func(t *testing.T) {
			var foreignRequests, originRequests, proxyRequests, callbacks atomic.Int64
			var foreignKey atomic.Value
			foreign := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				foreignRequests.Add(1)
				foreignKey.Store(r.Header.Get("X-API-Key"))
				io.WriteString(w, `{"status":"leaked"}`)
			}))
			defer foreign.Close()
			origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				originRequests.Add(1)
				if r.Header.Get("X-API-Key") != "synthetic-proxy-test-key" {
					t.Error("origin did not receive the synthetic key")
				}
				if r.URL.Path == "/done" || r.URL.EscapedPath() == "/done/%2F/%25/%3F/NU.V" || strings.HasPrefix(r.URL.Path, "/done/@") {
					io.WriteString(w, `{"status":"healthy"}`)
					return
				}
				http.Redirect(w, r, "/done", http.StatusFound)
			}))
			defer origin.Close()
			forward := &http.Transport{}
			defer forward.CloseIdleConnections()
			proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				proxyRequests.Add(1)
				request := r.Clone(r.Context())
				request.RequestURI = ""
				if request.URL.Scheme == "" && request.URL.Host == "" {
					request.URL.Scheme = "http"
					request.URL.Host = request.Host
				}
				response, err := forward.RoundTrip(request)
				if err != nil {
					t.Error(err)
					http.Error(w, "proxy forwarding failed", http.StatusBadGateway)
					return
				}
				defer response.Body.Close()
				for name, values := range response.Header {
					w.Header()[name] = values
				}
				w.WriteHeader(response.StatusCode)
				io.Copy(w, response.Body)
			}))
			defer proxy.Close()
			proxyURL, _ := url.Parse(proxy.URL)
			originURL, _ := url.Parse(origin.URL)
			foreignURL, _ := url.Parse(foreign.URL)
			transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
			defer transport.CloseIdleConnections()
			callback := func(r *http.Request, via []*http.Request) error {
				callbacks.Add(1)
				switch mode {
				case "foreign-host":
					r.Host = foreignURL.Host
				case "foreign-opaque":
					r.URL.Opaque = foreign.URL + "/done"
				case "foreign-opaque-authority":
					r.URL.Opaque = "//" + foreignURL.Host + "/done"
				case "foreign-path":
					r.URL.Path = "@" + foreignURL.Host + "/done"
					r.URL.RawPath = ""
				case "foreign-rawpath":
					r.URL.Path = "@" + foreignURL.Host + "/done/x"
					r.URL.RawPath = "@" + foreignURL.Host + "/done/%78"
				case "foreign-connect-authority":
					r.Method = http.MethodConnect
					r.URL.Path = ""
					r.URL.Opaque = foreignURL.Host
				case "same-origin-host":
					r.Host = originURL.Host
				case "empty-host":
					r.Host = ""
				case "same-origin-opaque-authority":
					r.URL.Opaque = "//" + originURL.Host + "/done/%2F/%25/%3F/NU.V"
				case "same-origin-opaque-absolute":
					r.URL.Opaque = origin.URL + "/done/%2F/%25/%3F/NU.V"
				case "same-origin-opaque-path":
					r.URL.Opaque = "/done/%2F/%25/%3F/NU.V"
				case "same-origin-escaped-path":
					r.URL.Path = "/done/@" + foreignURL.Host + "/x"
					r.URL.RawPath = "/done/@" + foreignURL.Host + "/%78"
				case "same-origin-query-fragment":
					r.URL.RawQuery = "@" + foreignURL.Host + "/done"
					r.URL.Fragment = foreign.URL + "/done"
				}
				return nil
			}
			hc := &http.Client{Transport: transport, CheckRedirect: callback, Timeout: time.Second}
			out, err := newTestClient(t, adanos.Config{BaseURL: origin.URL, APIKey: "synthetic-proxy-test-key", HTTPClient: hc}).Health(context.Background())
			receipt(t, map[string]any{"callback": mode, "base_url": origin.URL, "proxy": proxy.URL, "foreign_url": foreign.URL}, map[string]any{"foreign_requests": foreignRequests.Load(), "foreign_key": foreignKey.Load(), "origin_requests": originRequests.Load(), "proxy_requests": proxyRequests.Load(), "callbacks": callbacks.Load(), "response": out, "error": fmt.Sprint(err)})
			if foreignRequests.Load() != 0 {
				t.Errorf("foreign authority received %d request(s) with synthetic key %q", foreignRequests.Load(), foreignKey.Load())
			}
			if strings.HasPrefix(mode, "foreign-") {
				var apiErr *adanos.APIError
				if out != nil || !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusFound || originRequests.Load() != 1 || proxyRequests.Load() != 1 {
					t.Errorf("redirect was not blocked: out=%v err=%v origin=%d proxy=%d", out, err, originRequests.Load(), proxyRequests.Load())
				}
			} else if err != nil || out == nil || out.Status != "healthy" || originRequests.Load() != 2 || proxyRequests.Load() != 2 {
				t.Errorf("same-origin redirect failed: out=%v err=%v origin=%d proxy=%d", out, err, originRequests.Load(), proxyRequests.Load())
			}
			if callbacks.Load() != 1 || hc.Transport != transport || hc.Timeout != time.Second || reflect.ValueOf(hc.CheckRedirect).Pointer() != reflect.ValueOf(callback).Pointer() {
				t.Error("callback semantics or injected client changed")
			}
		})
	}
}

func TestConcurrentClientReuse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ticker := strings.TrimPrefix(r.URL.Path, "/news/stocks/v1/stock/")
		fmt.Fprintf(w, `{"ticker":%q,"sentiment_score":0}`, ticker)
	}))
	defer srv.Close()
	c := newTestClient(t, adanos.Config{BaseURL: srv.URL})
	o := &adanos.PeriodOptions{Days: new(7)}
	var wg sync.WaitGroup
	for i := range 48 {
		wg.Go(func() {
			ticker := fmt.Sprintf("S%d.V", i)
			out, err := c.News().Stock(context.Background(), ticker, o)
			if err != nil || out.Ticker != ticker || out.SentimentScore == nil || *out.SentimentScore != 0 {
				t.Errorf("response %v %v", out, err)
			}
		})
	}
	wg.Wait()
	if *o.Days != 7 {
		t.Fatal("options mutated")
	}
	receipt(t, map[string]any{"requests": 48, "days": 7}, "all distinct tickers retained")
}

func jsonValue(t *testing.T, raw string) any {
	t.Helper()
	if raw == "" {
		return nil
	}
	var value any
	if err := json.Unmarshal([]byte(raw), &value); err != nil {
		t.Fatal(err)
	}
	return value
}
func assertJSONEqual(t *testing.T, got, want string) {
	t.Helper()
	if !reflect.DeepEqual(jsonValue(t, got), jsonValue(t, want)) {
		t.Errorf("JSON body %q, want %q", got, want)
	}
}
