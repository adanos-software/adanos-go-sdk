// Code generated from API 1.54.0, SHA-256 dd8a500f064cac27f8ea6f5e76e8cabdf48b5c6347116bbb06c3f4fcd16258de; DO NOT EDIT.

package adanos

import "encoding/json"

// CompareLimitError: Error response for compare item-count violations (HTTP 400).
type CompareLimitError struct {
	// Detail: Structured error details
	Detail CompareLimitErrorDetail `json:"detail"`
}

// CompareLimitErrorDetail: Structured error detail for compare item-count violations.
type CompareLimitErrorDetail struct {
	// Error: Machine-readable error code
	Error string `json:"error"`
	// ItemName: Compared item type, such as ticker or symbol
	ItemName string `json:"item_name"`
	// MaxItems: Maximum accepted compare items
	MaxItems int64 `json:"max_items"`
	// Message: Human-readable error message
	Message string `json:"message"`
}

// CompareResponse: Comparison of multiple stocks.
type CompareResponse struct {
	// PeriodDays: Analysis period in days
	PeriodDays int64 `json:"period_days"`
	// Stocks: Stocks sorted by buzz_score descending
	Stocks []CompareStockItem `json:"stocks"`
}

// CompareStockItem: Individual stock in comparison.
type CompareStockItem struct {
	// BearishPct: Percentage of bearish mentions
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Percentage of bullish mentions
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Buzz Score (0-100). Asymptotic scaling above 50.
	BuzzScore *float64 `json:"buzz_score"`
	// CompanyName: Company name
	CompanyName *string `json:"company_name,omitempty"`
	// Mentions: Total attributed mentions in period, including inherited thread-context mentions
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment (-1 to +1, null if no mentions)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// SubredditCount: Number of subreddits with mentions in period
	SubredditCount int64 `json:"subreddit_count"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// TotalUpvotes: Total upvotes from attributed mentions
	TotalUpvotes int64 `json:"total_upvotes"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, upvotes and subreddit breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Null when no qualifying Reddit data exists.
	Trend *string `json:"trend,omitempty"`
	// TrendHistory: Daily buzz scores (oldest→newest). Length follows the resolved period, with a minimum of 7 values when data exists. Empty when no qualifying data exists.
	TrendHistory []float64 `json:"trend_history,omitempty"`
	// UniquePosts: Number of unique posts in period
	UniquePosts int64 `json:"unique_posts"`
}

// CryptoCompareResponse: Comparison response for multiple crypto symbols.
type CryptoCompareResponse struct {
	// PeriodDays: Analysis period in days
	PeriodDays int64 `json:"period_days"`
	// Tokens: Requested tokens sorted by buzz_score descending
	Tokens []CryptoCompareTokenItem `json:"tokens"`
}

// CryptoCompareTokenItem: Individual crypto token in a compare response.
type CryptoCompareTokenItem struct {
	BearishPct *int64   `json:"bearish_pct,omitempty"`
	BullishPct *int64   `json:"bullish_pct,omitempty"`
	BuzzScore  *float64 `json:"buzz_score"`
	// Mentions: Attributed mentions in the selected period, including inherited thread-context mentions
	Mentions       int64    `json:"mentions"`
	Name           *string  `json:"name,omitempty"`
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// SubredditCount: Unique subreddit count
	SubredditCount int64  `json:"subreddit_count"`
	Symbol         string `json:"symbol"`
	// TotalUpvotes: Net upvotes (Reddit score) across attributed mentions; negative when downvotes prevail
	TotalUpvotes int64 `json:"total_upvotes"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, upvotes and subreddit breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Null when no qualifying data exists.
	Trend *string `json:"trend,omitempty"`
	// TrendHistory: Daily buzz scores (oldest to newest). Length follows the resolved period, with a minimum of 7 values when data exists. Empty when no qualifying data exists.
	TrendHistory []float64 `json:"trend_history,omitempty"`
	// UniquePosts: Unique Reddit post count
	UniquePosts int64 `json:"unique_posts"`
}

// CryptoDailyTrendItem: Daily trend datapoint for a crypto symbol.
type CryptoDailyTrendItem struct {
	// BearishPct: Bearish mention percentage for the day, calculated from negative sentiment-classified mentions
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Bullish mention percentage for the day, calculated from positive sentiment-classified mentions
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Buzz score for this date (0-100); null until the day is aggregated, while the activity fields are already filled
	BuzzScore *float64 `json:"buzz_score,omitempty"`
	// Date: Date in YYYY-MM-DD format Format: date.
	Date string `json:"date"`
	// Mentions: Attributed mention count for the day, including inherited thread-context mentions
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score for the day
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
}

// CryptoHealthResponse: Health response for Reddit Crypto service.
type CryptoHealthResponse struct {
	Error     *string                `json:"error,omitempty"`
	Scheduler *CryptoSchedulerStatus `json:"scheduler,omitempty"`
	Service   *string                `json:"service,omitempty"`
	Status    string                 `json:"status"`
	// TokensTracked: Distinct symbols observed in the most recent scrape cycle; returns 0 when scrape telemetry is unavailable.
	TokensTracked *int64 `json:"tokens_tracked,omitempty"`
	// TotalMentions: Mentions observed in the most recent scrape cycle; returns 0 when scrape telemetry is unavailable.
	TotalMentions *int64  `json:"total_mentions,omitempty"`
	Version       *string `json:"version,omitempty"`
}

// CryptoMarketSentimentDriver: Top asset contributing to the service-level Reddit Crypto market sentiment.
type CryptoMarketSentimentDriver struct {
	// BuzzScore: Asset buzz score in the selected period
	BuzzScore *float64 `json:"buzz_score"`
	// Mentions: Attributed Reddit mentions in the selected period, including inherited thread-context mentions
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score for this asset
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Symbol: Crypto symbol
	Symbol string `json:"symbol"`
}

// CryptoMarketSentimentResponse: Service-level Reddit Crypto market sentiment across all tracked assets.
type CryptoMarketSentimentResponse struct {
	// ActiveTickers: Number of crypto symbols with activity in the selected period
	ActiveTickers int64 `json:"active_tickers"`
	// BearishPct: Bearish mention percentage
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Bullish mention percentage
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Service-wide Reddit Crypto heat score relative to the trailing 90-day baseline for the selected period
	BuzzScore *float64 `json:"buzz_score"`
	// Drivers: Top 5 tokens of the selected period by buzz_score. mentions and sentiment_score cover every day of the period.
	Drivers []CryptoMarketSentimentDriver `json:"drivers,omitempty"`
	// Mentions: Total Reddit mentions across the selected period, including inherited thread-context mentions
	Mentions int64 `json:"mentions"`
	// NegativeCount: Deprecated. Prefer bearish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NegativeCount int64 `json:"negative_count"`
	// NeutralCount: Deprecated. Prefer bullish_pct, bearish_pct and mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NeutralCount int64 `json:"neutral_count"`
	// PositiveCount: Deprecated. Prefer bullish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	PositiveCount int64 `json:"positive_count"`
	// SentimentScore: Service-wide average sentiment score across the selected period
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// SubredditCount: Distinct subreddits contributing in the selected period
	SubredditCount int64 `json:"subreddit_count"`
	// TotalUpvotes: Total upvotes across attributed mentions in the selected period
	TotalUpvotes int64 `json:"total_upvotes"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, upvotes and subreddit breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
	// TrendHistory: Daily service-wide Reddit Crypto heat scores (oldest→newest), each calibrated relative to the trailing 90-day baseline. Length = max(requested_days, 7) for successful requests; windows before platform availability return HTTP 422.
	TrendHistory []float64 `json:"trend_history,omitempty"`
	// UniquePosts: Exact distinct Reddit post count across the selected UTC period
	UniquePosts int64 `json:"unique_posts"`
}

// CryptoRawMentionItem: Raw Reddit crypto mention row for one symbol.
type CryptoRawMentionItem struct {
	// Author: Public Reddit author name, if scraped
	Author *string `json:"author,omitempty"`
	// CommentID: Reddit comment identifier when the mention comes from a comment
	CommentID *string `json:"comment_id,omitempty"`
	// CreatedUTC: UTC timestamp of creation (ISO 8601) Format: date-time.
	CreatedUTC string `json:"created_utc"`
	// IsInherited: Whether symbol context was inherited from the parent thread instead of directly mentioned
	IsInherited bool `json:"is_inherited"`
	// PostID: Reddit post identifier
	PostID *string `json:"post_id,omitempty"`
	// SentimentLabel: Sentiment label (positive/negative/neutral)
	SentimentLabel *string `json:"sentiment_label,omitempty"`
	// SentimentScore: Sentiment score (-1 to +1)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Subreddit: Source subreddit
	Subreddit *string `json:"subreddit,omitempty"`
	// TextSnippet: Stored mention snippet (truncated at ingestion time)
	TextSnippet string `json:"text_snippet"`
	// Upvotes: Reddit score (upvotes) for this mention row
	Upvotes int64 `json:"upvotes"`
}

// CryptoRawMentionsResponse: Paginated raw Reddit crypto mentions for a single symbol.
type CryptoRawMentionsResponse struct {
	// Count: Total number of matching raw rows for the requested include_inherited filter before limit and offset are applied; only the requested page is returned
	Count int64 `json:"count"`
	// PeriodDays: Lookback window used for the query
	PeriodDays int64 `json:"period_days"`
	// Results: Raw mention rows ordered by newest first
	Results []CryptoRawMentionItem `json:"results"`
	// Symbol: Crypto symbol
	Symbol string `json:"symbol"`
}

// CryptoSchedulerStatus: Reddit Crypto data collection status.
type CryptoSchedulerStatus struct {
	// ErrorCount: Total failed scrapes
	ErrorCount *int64 `json:"error_count,omitempty"`
	// LastScrape: Last scrape timestamp
	LastScrape *string `json:"last_scrape,omitempty"`
	// Note: Additional status information
	Note *string `json:"note,omitempty"`
	// Running: Whether data collection is running
	Running bool `json:"running"`
	// ScrapeCount: Total scrapes
	ScrapeCount *int64 `json:"scrape_count,omitempty"`
	// SuccessRate: Success ratio
	SuccessRate *float64 `json:"success_rate,omitempty"`
	// TimedOutScrapeThreads: Timed-out collection attempts since the last successful reset.
	TimedOutScrapeThreads *int64 `json:"timed_out_scrape_threads,omitempty"`
}

// CryptoSearchItem: Single search result for crypto reference data.
type CryptoSearchItem struct {
	Aliases       []string `json:"aliases,omitempty"`
	MarketCapRank *int64   `json:"market_cap_rank,omitempty"`
	MarketCapUsd  *float64 `json:"market_cap_usd,omitempty"`
	Name          string   `json:"name"`
	// Summary: Last-7-days Reddit summary for this asset
	Summary CryptoSearchSummary `json:"summary"`
	Symbol  string              `json:"symbol"`
}

// CryptoSearchResponse: Search response for crypto assets.
type CryptoSearchResponse struct {
	// Count: Total number of matching assets before the limit is applied
	Count int64 `json:"count"`
	// PeriodDays: Fixed lookback window used for each summary block
	PeriodDays int64 `json:"period_days"`
	// Query: Original search query
	Query string `json:"query"`
	// Results: Top matching crypto assets capped by the requested limit
	Results []CryptoSearchItem `json:"results"`
}

// CryptoSearchSummary: Compact last-7-days summary attached to crypto search results.
type CryptoSearchSummary struct {
	// BuzzScore: Buzz score over the last 7 UTC calendar days
	BuzzScore *float64 `json:"buzz_score"`
	// Mentions: Attributed mentions in the last 7 UTC calendar days, including inherited thread-context mentions
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score over the last 7 UTC calendar days
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, upvotes and subreddit breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
}

// CryptoStatsResponse: High-level crypto dataset statistics.
type CryptoStatsResponse struct {
	MentionsToday   *int64 `json:"mentions_today,omitempty"`
	SupportedTokens int64  `json:"supported_tokens"`
	// TotalMentions: All-time number of stored Reddit Crypto mentions
	TotalMentions     int64  `json:"total_mentions"`
	UniqueTokens      int64  `json:"unique_tokens"`
	UniqueTokensToday *int64 `json:"unique_tokens_today,omitempty"`
}

// CryptoSubredditCount: Subreddit contributor metrics for a crypto detail response.
type CryptoSubredditCount struct {
	// BuzzScore: Contributor-level buzz score for this subreddit in the selected period
	BuzzScore *float64 `json:"buzz_score,omitempty"`
	// Count: Deprecated alias for mentions. Use mentions instead.
	// Deprecated: Retained for API compatibility.
	Count int64 `json:"count"`
	// Mentions: Attributed mentions in this subreddit
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score for this subreddit in the selected period
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Subreddit: Subreddit name
	Subreddit string `json:"subreddit"`
}

// CryptoTokenSentiment: Detailed Reddit sentiment for a single crypto symbol.
type CryptoTokenSentiment struct {
	// BearishPct: Bearish mention percentage
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Bullish mention percentage
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Buzz score
	BuzzScore *float64 `json:"buzz_score,omitempty"`
	// DailyTrend: Daily attributed mention breakdown for the selected period, with sentiment and buzz_score
	DailyTrend []CryptoDailyTrendItem `json:"daily_trend,omitempty"`
	// Found: Legacy v1 data-availability flag. `true` means this service has qualifying data for this asset/window. `false` means the asset is supported, but this service has no qualifying data for the requested window. Planned for removal in v2.
	// Deprecated: Retained for API compatibility.
	Found bool `json:"found"`
	// Mentions: Attributed mentions in the selected period, including inherited thread-context mentions
	Mentions *int64 `json:"mentions,omitempty"`
	// Name: Token name
	Name *string `json:"name,omitempty"`
	// NegativeCount: Deprecated. Prefer bearish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NegativeCount *int64 `json:"negative_count,omitempty"`
	// NeutralCount: Deprecated. Prefer bullish_pct, bearish_pct and mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NeutralCount *int64 `json:"neutral_count,omitempty"`
	// PeriodDays: Analysis period in days
	PeriodDays *int64 `json:"period_days,omitempty"`
	// PositiveCount: Deprecated. Prefer bullish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	PositiveCount *int64 `json:"positive_count,omitempty"`
	// SentimentScore: Average sentiment
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// SubredditCount: Distinct subreddit count
	SubredditCount *int64 `json:"subreddit_count,omitempty"`
	// Symbol: Token symbol
	Symbol string `json:"symbol"`
	// TopMentions: Top explicit mentions from the selected period, ranked by engagement
	TopMentions []CryptoTopMention `json:"top_mentions,omitempty"`
	// TopSubreddits: Top subreddits by mentions
	TopSubreddits []CryptoSubredditCount `json:"top_subreddits,omitempty"`
	// TotalUpvotes: Net upvotes (Reddit score) across attributed mentions; negative when downvotes prevail
	TotalUpvotes *int64 `json:"total_upvotes,omitempty"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, upvotes and subreddit breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
	// UniquePosts: Distinct post count
	UniquePosts *int64 `json:"unique_posts,omitempty"`
}

// CryptoTopMention: Top Reddit mention ordered by engagement.
type CryptoTopMention struct {
	// CreatedUTC: UTC timestamp of creation (ISO 8601) Format: date-time.
	CreatedUTC string `json:"created_utc"`
	// SentimentLabel: Sentiment label
	SentimentLabel *string `json:"sentiment_label,omitempty"`
	// SentimentScore: Sentiment score (-1 to +1)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Subreddit: Source subreddit
	Subreddit *string `json:"subreddit,omitempty"`
	// TextSnippet: Mention text snippet
	TextSnippet string `json:"text_snippet"`
	// Upvotes: Upvote count
	Upvotes int64 `json:"upvotes"`
}

// CryptoTrendingToken: Trending crypto token summary.
type CryptoTrendingToken struct {
	// BearishPct: Bearish mention percentage
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Bullish mention percentage
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Buzz score (0-100)
	BuzzScore float64 `json:"buzz_score"`
	// Mentions: Attributed mention count, including inherited thread-context mentions
	Mentions int64 `json:"mentions"`
	// Name: Token name
	Name *string `json:"name,omitempty"`
	// SentimentScore: Average sentiment score
	SentimentScore float64 `json:"sentiment_score"`
	// SubredditCount: Unique subreddit count
	SubredditCount int64 `json:"subreddit_count"`
	// Symbol: Token symbol
	Symbol string `json:"symbol"`
	// TotalUpvotes: Net upvotes (Reddit score) across attributed mentions; negative when downvotes prevail
	TotalUpvotes int64 `json:"total_upvotes"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, upvotes and subreddit breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Values: rising, falling, stable.
	Trend string `json:"trend"`
	// TrendHistory: Daily buzz scores (oldest to newest). Length follows the resolved period, with a minimum of 7 values.
	TrendHistory []float64 `json:"trend_history,omitempty"`
	// UniquePosts: Unique Reddit post count
	UniquePosts int64 `json:"unique_posts"`
}

// DailyTrendItem: Daily trend data point.
type DailyTrendItem struct {
	// BearishPct: Bearish mention percentage for this date, calculated from negative sentiment-classified mentions
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Bullish mention percentage for this date, calculated from positive sentiment-classified mentions
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Buzz score for this date (0-100); null until the day is aggregated, while the activity fields are already filled
	BuzzScore *float64 `json:"buzz_score,omitempty"`
	// Date: Date in YYYY-MM-DD format Format: date.
	Date string `json:"date"`
	// Mentions: Attributed mention count for this date, including inherited thread-context mentions
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score for this date
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
}

// ErrorResponse: Standard error response.
type ErrorResponse struct {
	// Detail: Human-readable error message
	Detail string `json:"detail"`
}

// HTTPValidationError: FastAPI/Pydantic validation error response.
type HTTPValidationError struct {
	// Detail: Validation error details
	Detail []ValidationError `json:"detail,omitempty"`
}

// HealthResponse: API health status.
type HealthResponse struct {
	// Error: Error message if unhealthy
	Error *string `json:"error,omitempty"`
	// Scheduler: Scheduler status information
	Scheduler *SchedulerStatus `json:"scheduler,omitempty"`
	// Service: Service identifier
	Service *string `json:"service,omitempty"`
	// Status: Current health status
	Status string `json:"status"`
	// TickersTracked: Distinct tickers observed in the most recent scrape cycle; returns 0 when scrape telemetry is unavailable.
	TickersTracked *int64 `json:"tickers_tracked,omitempty"`
	// TotalMentions: Mentions observed in the most recent scrape cycle; returns 0 when scrape telemetry is unavailable.
	TotalMentions *int64 `json:"total_mentions,omitempty"`
	// Version: Service version
	Version *string `json:"version,omitempty"`
}

// HistoricalLimitError: Error response for historical data limit exceeded (HTTP 403).
type HistoricalLimitError struct {
	// Detail: Structured error details
	Detail HistoricalLimitErrorDetail `json:"detail"`
}

// HistoricalLimitErrorDetail: Structured error detail for historical data limit exceeded.
type HistoricalLimitErrorDetail struct {
	// AccountType: User's account type (free, hobby, professional, premium)
	AccountType string `json:"account_type"`
	// EnforcedQuantity: Quantity checked against the account tier limit
	EnforcedQuantity *string `json:"enforced_quantity,omitempty"`
	// Error: Error type identifier
	Error string `json:"error"`
	// MaxDays: Maximum days allowed for account tier
	MaxDays int64 `json:"max_days"`
	// Message: Human-readable error message
	Message string `json:"message"`
	// RequestedDays: Number of days requested
	RequestedDays int64 `json:"requested_days"`
}

// InvalidPeriodError: Error response for invalid period requests (HTTP 422).
type InvalidPeriodError struct {
	// Detail: Structured error details
	Detail InvalidPeriodErrorDetail `json:"detail"`
}

// InvalidPeriodErrorDetail: Structured error detail for invalid period windows.
type InvalidPeriodErrorDetail struct {
	// AvailableSince: Earliest public data date for the platform
	AvailableSince *string `json:"available_since,omitempty"`
	// Error: Error type identifier
	Error string `json:"error"`
	// Field: Invalid query parameter name
	Field *string `json:"field,omitempty"`
	// MaxDays: Maximum days accepted by the endpoint
	MaxDays *int64 `json:"max_days,omitempty"`
	// Message: Human-readable error message
	Message string `json:"message"`
	// PeriodFrom: Resolved inclusive UTC start date
	PeriodFrom *string `json:"period_from,omitempty"`
	// PeriodTo: Resolved inclusive UTC end date
	PeriodTo *string `json:"period_to,omitempty"`
	// Platform: Platform identifier
	Platform *string `json:"platform,omitempty"`
	// RequestedDays: Number of days requested
	RequestedDays *int64 `json:"requested_days,omitempty"`
	// RetentionFrom: Earliest date retained for the endpoint
	RetentionFrom *string `json:"retention_from,omitempty"`
	// Today: Current UTC date used for validation
	Today *string `json:"today,omitempty"`
	// Value: Invalid query parameter value
	Value json.RawMessage `json:"value,omitempty"`
}

// InvalidSearchParametersError: Error response for invalid search query parameters.
type InvalidSearchParametersError struct {
	// Detail: Structured search-parameter error or normalized-query validation message
	Detail json.RawMessage `json:"detail"`
}

// InvalidSearchParametersErrorDetail: Structured error detail for unsupported search query parameters.
type InvalidSearchParametersErrorDetail struct {
	// Error: Error type identifier
	Error string `json:"error"`
	// Message: Human-readable error message
	Message string `json:"message"`
	// UnsupportedParameters: Search query parameters that are not supported
	UnsupportedParameters []string `json:"unsupported_parameters"`
}

// NewsCompareResponse: Comparison of multiple stocks in news.
type NewsCompareResponse struct {
	// PeriodDays: Analysis period in days
	PeriodDays int64 `json:"period_days"`
	// Stocks: Stocks sorted by buzz_score descending
	Stocks []NewsCompareStockItem `json:"stocks"`
}

// NewsCompareStockItem: Individual stock in news comparison.
type NewsCompareStockItem struct {
	// BearishPct: Percentage of bearish mentions
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Percentage of bullish mentions
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Buzz Score (0-100). Asymptotic scaling above 50.
	BuzzScore *float64 `json:"buzz_score"`
	// CompanyName: Company or fund name from resolved instrument metadata or the Legacy reference
	CompanyName *string `json:"company_name,omitempty"`
	// Mentions: Total mentions in period
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment (-1 to +1, null if no mentions)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// SourceCount: Number of distinct news sources with mentions in period
	SourceCount int64 `json:"source_count"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions/articles and source breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Null when no qualifying news data exists.
	Trend *string `json:"trend,omitempty"`
	// TrendHistory: Daily buzz scores (oldest→newest). Length follows the resolved period, with a minimum of 7 values when data exists. Empty when no qualifying data exists.
	TrendHistory []float64 `json:"trend_history,omitempty"`
}

// NewsMarketSentimentDriver: Top asset contributing to the service-level News market sentiment.
type NewsMarketSentimentDriver struct {
	// BuzzScore: Asset buzz score in the selected period
	BuzzScore *float64 `json:"buzz_score"`
	// Mentions: News mentions in the selected period
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average news sentiment score for this asset
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
}

// NewsMarketSentimentResponse: Service-level News market sentiment across all tracked stocks.
type NewsMarketSentimentResponse struct {
	// ActiveTickers: Number of tickers with news activity in the selected period
	ActiveTickers int64 `json:"active_tickers"`
	// BearishPct: Bearish news mention percentage
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Bullish news mention percentage
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Service-wide News heat score relative to the service's trailing 90-day baseline. Around 50 = normal coverage heat, higher values = hotter-than-usual coverage.
	BuzzScore *float64 `json:"buzz_score"`
	// Drivers: Top assets by current buzz_score driving the service-level reading
	Drivers []NewsMarketSentimentDriver `json:"drivers,omitempty"`
	// Mentions: News mentions across the selected period
	Mentions int64 `json:"mentions"`
	// NegativeCount: Deprecated. Prefer bearish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NegativeCount int64 `json:"negative_count"`
	// NeutralCount: Deprecated. Prefer bullish_pct, bearish_pct and mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NeutralCount int64 `json:"neutral_count"`
	// PositiveCount: Deprecated. Prefer bullish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	PositiveCount int64 `json:"positive_count"`
	// SentimentScore: Service-wide average news sentiment score
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// SourceCount: Exact distinct news-source breadth across the selected UTC period
	SourceCount int64 `json:"source_count"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions/articles and source breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
	// TrendHistory: Daily service-wide buzz scores (oldest→newest) using the same relative baseline calibration. Length = max(requested_days, 7) for successful requests; windows before platform availability return HTTP 422.
	TrendHistory []float64 `json:"trend_history,omitempty"`
	// UniqueArticles: Exact distinct article count across the selected UTC period
	UniqueArticles int64 `json:"unique_articles"`
}

// NewsRawMentionItem: Raw news mention row for one stock ticker.
type NewsRawMentionItem struct {
	// ArticleID: Article identifier when available
	ArticleID *string `json:"article_id,omitempty"`
	// Author: Public article author/byline when available
	Author *string `json:"author,omitempty"`
	// CreatedUTC: ISO timestamp of creation when available
	CreatedUTC *string `json:"created_utc,omitempty"`
	// SentimentLabel: Sentiment label (positive/negative/neutral)
	SentimentLabel *string `json:"sentiment_label,omitempty"`
	// SentimentScore: Sentiment score (-1 to +1)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Source: Source publication identifier
	Source string `json:"source"`
	// Summary: Article summary
	Summary *string `json:"summary,omitempty"`
	// TextSnippet: Stored article snippet (truncated at ingestion time)
	TextSnippet string `json:"text_snippet"`
	// Title: Article title
	Title *string `json:"title,omitempty"`
	// URL: Article URL
	URL *string `json:"url,omitempty"`
}

// NewsRawMentionsResponse: Paginated raw news mentions for a single stock ticker.
type NewsRawMentionsResponse struct {
	// Count: Total number of matching raw rows before limit and offset are applied; only the requested page is returned
	Count int64 `json:"count"`
	// PeriodDays: Lookback window used for the query
	PeriodDays int64 `json:"period_days"`
	// Results: Raw news mention rows ordered by newest first
	Results []NewsRawMentionItem `json:"results"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
}

// NewsSearchResponse: Search response for news assets.
type NewsSearchResponse struct {
	// Count: Total number of matching results before the limit is applied
	Count int64 `json:"count"`
	// PeriodDays: Fixed lookback window used for each summary block
	PeriodDays int64 `json:"period_days"`
	// Query: Original search query
	Query string `json:"query"`
	// Results: Top matching stocks capped by the requested limit
	Results []NewsSearchResultItem `json:"results"`
}

// NewsSearchResultItem: Individual news search result item.
type NewsSearchResultItem struct {
	// Aliases: Configured News graph matching terms, including terms with no mentions in the summary window; article acceptance also depends on worker safety and context checks
	Aliases []string `json:"aliases,omitempty"`
	// Country: Country of headquarters
	Country *string `json:"country,omitempty"`
	// Exchange: Stock exchange (NYSE, NASDAQ, etc.)
	Exchange *string `json:"exchange,omitempty"`
	// Name: Company name
	Name string `json:"name"`
	// Sector: Industry sector
	Sector *string `json:"sector,omitempty"`
	// Summary: Last-7-days news summary for this asset
	Summary NewsSearchSummary `json:"summary"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// Type: Asset type (Stock, ETF, etc.)
	Type *string `json:"type,omitempty"`
}

// NewsSearchSummary: Compact last-7-days summary attached to news search results.
type NewsSearchSummary struct {
	// BuzzScore: Buzz score over the last 7 UTC calendar days
	BuzzScore *float64 `json:"buzz_score"`
	// Mentions: News mentions in the last 7 UTC calendar days
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score over the last 7 UTC calendar days
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions/articles and source breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
}

// NewsSourceCount: News source contributor metrics for a stock detail response.
type NewsSourceCount struct {
	// BuzzScore: Contributor-level buzz score for this source in the selected period
	BuzzScore *float64 `json:"buzz_score,omitempty"`
	// Count: Deprecated alias for mentions. Use mentions instead.
	// Deprecated: Retained for API compatibility.
	Count int64 `json:"count"`
	// Mentions: Mentions from this source
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score for this source in the selected period
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Source: News source identifier
	Source string `json:"source"`
}

// NewsStatsResponse: News service statistics.
type NewsStatsResponse struct {
	// MentionsToday: News mention rows created since today's UTC midnight
	MentionsToday *int64 `json:"mentions_today,omitempty"`
	// SupportedTickers: Number of supported public tickers in the News reference catalog
	SupportedTickers int64 `json:"supported_tickers"`
	// TotalMentions: All-time number of stored news stock mentions
	TotalMentions int64 `json:"total_mentions"`
	// UniqueTickers: Number of unique tickers with news mentions
	UniqueTickers int64 `json:"unique_tickers"`
	// UniqueTickersToday: Unique tickers with news mention rows since today's UTC midnight
	UniqueTickersToday *int64 `json:"unique_tickers_today,omitempty"`
}

// NewsStockSentiment: Detailed sentiment analysis for a specific stock ticker from news mentions.
type NewsStockSentiment struct {
	// BearishPct: Percentage of bearish mentions
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Percentage of bullish mentions
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Buzz Score (0-100). Asymptotic scaling above 50.
	BuzzScore *float64 `json:"buzz_score,omitempty"`
	// CompanyName: Company or fund name from resolved instrument metadata or the Legacy reference (null if unavailable)
	CompanyName *string `json:"company_name,omitempty"`
	// DailyTrend: Daily mention breakdown for the selected period, with sentiment and buzz_score
	DailyTrend []DailyTrendItem `json:"daily_trend,omitempty"`
	// Found: Legacy v1 data-availability flag. `true` means this service has qualifying data for this asset/window. `false` means the asset is supported, but this service has no qualifying data for the requested window. Planned for removal in v2.
	// Deprecated: Retained for API compatibility.
	Found bool `json:"found"`
	// Mentions: Total mentions in the selected period
	Mentions *int64 `json:"mentions,omitempty"`
	// NegativeCount: Deprecated. Prefer bearish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NegativeCount *int64 `json:"negative_count,omitempty"`
	// NeutralCount: Deprecated. Prefer bullish_pct, bearish_pct and mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NeutralCount *int64 `json:"neutral_count,omitempty"`
	// PeriodDays: Analysis period in days
	PeriodDays *int64 `json:"period_days,omitempty"`
	// PositiveCount: Deprecated. Prefer bullish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	PositiveCount *int64 `json:"positive_count,omitempty"`
	// SentimentScore: Average sentiment score (-1 bearish to +1 bullish)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// SourceCount: Number of sources with mentions
	SourceCount *int64 `json:"source_count,omitempty"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// TopMentions: Top news mentions from the selected period, ranked by recency
	TopMentions []NewsTopMention `json:"top_mentions,omitempty"`
	// TopSources: Top sources by mention count
	TopSources []NewsSourceCount `json:"top_sources,omitempty"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions/articles and source breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
}

// NewsTopMention: Top mention item from news.
type NewsTopMention struct {
	// CreatedUTC: ISO timestamp of creation
	CreatedUTC string `json:"created_utc"`
	// SentimentLabel: Sentiment label (positive/negative/neutral)
	SentimentLabel string `json:"sentiment_label"`
	// SentimentScore: Sentiment score (-1 to +1)
	SentimentScore float64 `json:"sentiment_score"`
	// Source: Source publication identifier
	Source string `json:"source"`
	// TextSnippet: Excerpt from the mention
	TextSnippet string `json:"text_snippet"`
}

// NewsTrendingCountry: Trending country with aggregated buzz metrics from news.
type NewsTrendingCountry struct {
	// BearishPct: Percentage of bearish mentions
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Percentage of bullish mentions
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Aggregated Buzz Score (0-100)
	BuzzScore float64 `json:"buzz_score"`
	// Country: Country name
	Country string `json:"country"`
	// Mentions: Total mentions
	Mentions int64 `json:"mentions"`
	// SentimentScore: Weighted average sentiment (-1 bearish to +1 bullish)
	SentimentScore float64 `json:"sentiment_score"`
	// SourceCount: Number of distinct news sources with mentions
	SourceCount int64 `json:"source_count"`
	// TopTickers: Top 5 tickers by mentions
	TopTickers []string `json:"top_tickers"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions/articles and source breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Values: rising, falling, stable.
	Trend string `json:"trend"`
	// TrendHistory: Daily dimension buzz scores (oldest to newest). Length follows the resolved period, with a minimum of 7 values. Every slot describes only its UTC calendar day. The live final value describes today so far and can differ from period buzz_score. Closed historical windows end at `to`.
	TrendHistory []float64 `json:"trend_history"`
	// UniqueTickers: Number of unique tickers mentioned
	UniqueTickers int64 `json:"unique_tickers"`
}

// NewsTrendingSector: Trending sector with aggregated buzz metrics from news.
type NewsTrendingSector struct {
	// BearishPct: Percentage of bearish mentions
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Percentage of bullish mentions
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Aggregated Buzz Score (0-100)
	BuzzScore float64 `json:"buzz_score"`
	// Mentions: Total mentions
	Mentions int64 `json:"mentions"`
	// Sector: Industry sector name
	Sector string `json:"sector"`
	// SentimentScore: Weighted average sentiment (-1 bearish to +1 bullish)
	SentimentScore float64 `json:"sentiment_score"`
	// SourceCount: Number of distinct news sources with mentions
	SourceCount int64 `json:"source_count"`
	// TopTickers: Top 5 tickers by mentions
	TopTickers []string `json:"top_tickers"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions/articles and source breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Values: rising, falling, stable.
	Trend string `json:"trend"`
	// TrendHistory: Daily dimension buzz scores (oldest to newest). Length follows the resolved period, with a minimum of 7 values. Every slot describes only its UTC calendar day. The live final value describes today so far and can differ from period buzz_score. Closed historical windows end at `to`.
	TrendHistory []float64 `json:"trend_history"`
	// UniqueTickers: Number of unique tickers mentioned
	UniqueTickers int64 `json:"unique_tickers"`
}

// NewsTrendingStock: Trending stock with buzz metrics from news sources.
type NewsTrendingStock struct {
	// BearishPct: Percentage of bearish mentions
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Percentage of bullish mentions
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Buzz Score (0-100). Asymptotic scaling above 50.
	BuzzScore float64 `json:"buzz_score"`
	// CompanyName: Company or fund name from resolved instrument metadata or the Legacy reference (null if unavailable)
	CompanyName *string `json:"company_name,omitempty"`
	// Mentions: Total number of mentions
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score (-1 bearish to +1 bullish)
	SentimentScore float64 `json:"sentiment_score"`
	// SourceCount: Number of distinct news sources with mentions
	SourceCount int64 `json:"source_count"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions/articles and source breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Values: rising, falling, stable.
	Trend string `json:"trend"`
	// TrendHistory: Daily buzz scores (oldest→newest). Length follows the resolved period, with a minimum of 7 values. Values 0-100.
	TrendHistory []float64 `json:"trend_history,omitempty"`
}

// PolymarketCompareResponse: Response for compare endpoint.
type PolymarketCompareResponse struct {
	// PeriodDays: Analysis period in days
	PeriodDays int64 `json:"period_days"`
	// Stocks: Stocks sorted by buzz_score descending
	Stocks []PolymarketCompareStockItem `json:"stocks"`
}

// PolymarketCompareStockItem: Single stock item for Polymarket compare endpoint.
type PolymarketCompareStockItem struct {
	// BearishPct: Outcome-aware bearish market percentage
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Outcome-aware bullish market percentage
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Buzz score (0-100)
	BuzzScore *float64 `json:"buzz_score"`
	// CompanyName: Company name
	CompanyName *string `json:"company_name,omitempty"`
	// CurrentMarketCount: Number of currently open markets in the latest UTC-day snapshot for ticker; use this for live-only market breadth
	CurrentMarketCount int64 `json:"current_market_count"`
	// MarketCount: Distinct markets that existed for ticker within the selected UTC window, counted by market-catalog lifetime overlap
	MarketCount int64 `json:"market_count"`
	// SentimentScore: Orderbook-aware implied sentiment; null when no directional sentiment evidence exists in the selected period
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// TotalLiquidity: Windowed aggregated liquidity signal in USD over the selected period; not a current snapshot
	TotalLiquidity float64 `json:"total_liquidity"`
	// TradeCount: Trade count in selected period
	TradeCount int64 `json:"trade_count"`
	// Trend: Flow momentum over the current 3 UTC days vs previous 3 UTC days using trades, volume, market breadth and liquidity; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Null when no qualifying activity exists.
	Trend *string `json:"trend,omitempty"`
	// TrendHistory: Recent buzz history (oldest→newest). Length follows the resolved period, with a minimum of 7 values when data exists. Closed historical windows end at `to`. Empty when no qualifying data exists.
	TrendHistory []float64 `json:"trend_history,omitempty"`
	// UniqueTraders: Exact distinct union of observed proxy and explicit maker wallet hashes for the delivered scope and UTC window; null when retained wallet-level trades do not fully cover the window; not a complete market-wide participant census
	UniqueTraders *int64 `json:"unique_traders,omitempty"`
}

// PolymarketDailyTrendItem: Daily trend item for stock detail endpoint.
type PolymarketDailyTrendItem struct {
	// BearishPct: Outcome-aware bearish market-direction percentage for this date
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Outcome-aware bullish market-direction percentage for this date
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Buzz score on this date; null when neither a daily aggregate nor a live fallback could be computed for this date
	BuzzScore *float64 `json:"buzz_score,omitempty"`
	// Date: Date in YYYY-MM-DD format
	Date string `json:"date"`
	// SentimentScore: Outcome-aware implied sentiment on this date; can come from orderbook evidence or legacy traded price evidence, null when no sentiment evidence exists
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// TradeCount: Trade count on this date
	TradeCount int64 `json:"trade_count"`
}

// PolymarketHealthResponse: Polymarket health response.
type PolymarketHealthResponse struct {
	// Error: Error message when unhealthy
	Error *string `json:"error,omitempty"`
	// Scheduler: Polymarket worker status
	Scheduler *PolymarketSchedulerStatus `json:"scheduler,omitempty"`
	// Service: Service identifier
	Service string `json:"service"`
	// Status: Current health status
	Status string `json:"status"`
	// TickersTracked: Distinct tickers observed in the most recent scrape cycle; returns 0 when scrape telemetry is unavailable.
	TickersTracked int64 `json:"tickers_tracked"`
	// TotalMentions: Mentions observed in the most recent scrape cycle; returns 0 when scrape telemetry is unavailable.
	TotalMentions int64 `json:"total_mentions"`
	// Version: API version
	Version string `json:"version"`
}

// PolymarketMarketSentimentDriver: Top asset contributing to the service-level Polymarket market sentiment.
type PolymarketMarketSentimentDriver struct {
	// BuzzScore: Asset buzz score in the selected period
	BuzzScore *float64 `json:"buzz_score"`
	// SentimentScore: Implied sentiment score for this asset in the selected period
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// TradeCount: Trade count for this asset in the selected period
	TradeCount int64 `json:"trade_count"`
}

// PolymarketMarketSentimentResponse: Service-level Polymarket market sentiment across all tracked stocks.
type PolymarketMarketSentimentResponse struct {
	// ActiveTickers: Number of tickers with market activity in the selected period
	ActiveTickers int64 `json:"active_tickers"`
	// BearishPct: Outcome-aware bearish market percentage
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Outcome-aware bullish market percentage
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Service-wide Polymarket heat score relative to the service's trailing 90-day baseline. Around 50 = normal activity, higher values = hotter-than-usual Polymarket activity.
	BuzzScore *float64 `json:"buzz_score"`
	// CurrentMarketCount: Sum of ticker-level currently open markets in the latest UTC-day snapshot; use this for live-only market breadth
	CurrentMarketCount int64 `json:"current_market_count"`
	// Drivers: Top 5 stocks of the selected period by buzz_score. trade_count and sentiment_score cover every day of the period.
	Drivers []PolymarketMarketSentimentDriver `json:"drivers,omitempty"`
	// MarketCount: Sum of ticker-level catalog lifetime-overlap market counts in the selected UTC window; not a cross-ticker global condition-id union
	MarketCount int64 `json:"market_count"`
	// NegativeCount: Deprecated. Outcome-aware bearish market count retained for v1 compatibility; prefer bearish_pct with trade_count and market_count for public analysis.
	// Deprecated: Retained for API compatibility.
	NegativeCount int64 `json:"negative_count"`
	// NeutralCount: Deprecated. Outcome-aware neutral/unclassified market count retained for v1 compatibility; prefer bullish_pct, bearish_pct, trade_count and market_count for public analysis.
	// Deprecated: Retained for API compatibility.
	NeutralCount int64 `json:"neutral_count"`
	// PositiveCount: Deprecated. Outcome-aware bullish market count retained for v1 compatibility; prefer bullish_pct with trade_count and market_count for public analysis.
	// Deprecated: Retained for API compatibility.
	PositiveCount int64 `json:"positive_count"`
	// SentimentScore: Service-wide weighted orderbook-aware implied sentiment score
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// TotalLiquidity: Windowed aggregated liquidity signal in USD over the selected period
	TotalLiquidity float64 `json:"total_liquidity"`
	// TradeCount: Service-wide trade count in the selected period
	TradeCount int64 `json:"trade_count"`
	// Trend: Flow momentum over the current 3 UTC days vs previous 3 UTC days using trades, volume, market breadth and liquidity; not price movement. Null when the selected window has no measurable Polymarket heat. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
	// TrendHistory: Daily service-wide buzz scores (oldest→newest) using the same relative baseline calibration. Length = max(requested_days, 7) for successful requests; windows before platform availability return HTTP 422.
	TrendHistory []float64 `json:"trend_history,omitempty"`
	// UniqueTraders: Exact service-wide observed participant/proxy-wallet union in the selected period; null when retained wallet-level trades do not fully cover the requested window
	UniqueTraders *int64 `json:"unique_traders,omitempty"`
}

// PolymarketPulseEvidence: Compact data-quality evidence for the ticker-level Polymarket pulse.
type PolymarketPulseEvidence struct {
	// AvgSpread: Average YES-token spread across current open snapshot markets with bid/ask spread evidence
	AvgSpread *float64 `json:"avg_spread,omitempty"`
	// DirectionalCoverage: Share of current open snapshot markets with outcome-aware directional evidence
	DirectionalCoverage *float64 `json:"directional_coverage,omitempty"`
	// SnapshotAt: Latest stored snapshot timestamp used by the pulse read model
	SnapshotAt *string `json:"snapshot_at,omitempty"`
	// TradedMarketPct: Share of current open snapshot markets with trades in the latest UTC day
	TradedMarketPct *float64 `json:"traded_market_pct,omitempty"`
	// ZeroTradeMarketPct: Share of current open snapshot markets without trades in the latest UTC day
	ZeroTradeMarketPct *float64 `json:"zero_trade_market_pct,omitempty"`
}

// PolymarketRawMentionItem: Raw Polymarket market snapshot row for one stock ticker.
type PolymarketRawMentionItem struct {
	// Active: Whether the market is currently open; false for closed or ended markets
	Active bool `json:"active"`
	// BuyTrades: UTC-day Polymarket BUY trades for this snapshot; activity counter, not bullishness
	BuyTrades int64 `json:"buy_trades"`
	// ConditionID: Polymarket condition ID
	ConditionID string `json:"condition_id"`
	// EndDate: Market end timestamp when available. Unlike summary `top_mentions`, raw rows preserve the full timestamp from the stored snapshot.
	EndDate *string `json:"end_date,omitempty"`
	// EventID: Polymarket event ID
	EventID string `json:"event_id"`
	// FetchedAt: Snapshot fetch timestamp Format: date-time.
	FetchedAt string `json:"fetched_at"`
	// Liquidity: Market liquidity in USD at snapshot time
	Liquidity float64 `json:"liquidity"`
	// MarketSlug: Polymarket market slug
	MarketSlug *string `json:"market_slug,omitempty"`
	// MarketStatus: Compact current market status derived from active, closed, accepting_orders and end_date Values: tradable, open, resolved, inactive, expired.
	MarketStatus string `json:"market_status"`
	// MarketType: Derived market type (`up_down`, `close_above`, `hit_target`, `earnings` or `other`)
	MarketType string `json:"market_type"`
	// NoPrice: NO price from this snapshot
	NoPrice *float64 `json:"no_price,omitempty"`
	// Question: Market question/title
	Question string `json:"question"`
	// ReferencePrice: Stock reference price attached during scrape when needed for range-ladder classification
	ReferencePrice *float64 `json:"reference_price,omitempty"`
	// ReferencePriceTimestamp: Timestamp for `reference_price` when available
	ReferencePriceTimestamp *string `json:"reference_price_timestamp,omitempty"`
	// SellTrades: UTC-day Polymarket SELL trades for this snapshot; activity counter, not bearishness
	SellTrades int64 `json:"sell_trades"`
	// SentimentLabel: Classification derived from `sentiment_score`: positive=bullish, negative=bearish, neutral=flat; null when `sentiment_score` is null
	SentimentLabel *string `json:"sentiment_label,omitempty"`
	// SentimentScore: Outcome-aware market price sentiment for this snapshot; prefers orderbook midpoint when available, falls back to YES/NO price and is null when direction or evidence is unavailable
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// StrikePrice: Extracted strike price when available
	StrikePrice *float64 `json:"strike_price,omitempty"`
	// TradeCount: UTC-day trade count observed for this market snapshot
	TradeCount int64 `json:"trade_count"`
	// UniqueTraders: UTC-day observed participant/proxy-wallet counter for this market snapshot
	UniqueTraders int64 `json:"unique_traders"`
	// Volume24h: Rolling 24h market volume in USD at snapshot time
	Volume24h float64 `json:"volume_24h"`
	// YesPrice: YES price from this snapshot
	YesPrice *float64 `json:"yes_price,omitempty"`
}

// PolymarketRawMentionsResponse: Paginated raw Polymarket snapshot rows for a single stock ticker.
type PolymarketRawMentionsResponse struct {
	// Count: Total number of matching raw rows before limit and offset are applied; only the requested page is returned
	Count int64 `json:"count"`
	// PeriodDays: Lookback window used for the query
	PeriodDays int64 `json:"period_days"`
	// Results: Raw market snapshot rows ordered by newest first
	Results []PolymarketRawMentionItem `json:"results"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
}

// PolymarketSchedulerStatus: Polymarket data collection status.
type PolymarketSchedulerStatus struct {
	// ErrorCount: Total failed cycles
	ErrorCount *int64 `json:"error_count,omitempty"`
	// LastScrape: ISO timestamp of last scrape
	LastScrape *string `json:"last_scrape,omitempty"`
	// Note: Additional worker status note
	Note *string `json:"note,omitempty"`
	// Running: Whether data collection is running
	Running bool `json:"running"`
	// ScrapeCount: Total scrape cycles
	ScrapeCount *int64 `json:"scrape_count,omitempty"`
	// SuccessRate: Success ratio
	SuccessRate *float64 `json:"success_rate,omitempty"`
}

// PolymarketSearchResponse: Search results for Polymarket stocks.
type PolymarketSearchResponse struct {
	// Count: Total number of matching results before the limit is applied
	Count int64 `json:"count"`
	// PeriodDays: Fixed lookback window used for each summary block
	PeriodDays int64 `json:"period_days"`
	// Query: Original search query
	Query string `json:"query"`
	// Results: Top matching stocks capped by the requested limit
	Results []PolymarketSearchResultItem `json:"results"`
}

// PolymarketSearchResultItem: Individual Polymarket search result item.
type PolymarketSearchResultItem struct {
	// Aliases: Alternative names/aliases
	Aliases []string `json:"aliases,omitempty"`
	// Country: Country of headquarters
	Country *string `json:"country,omitempty"`
	// Exchange: Stock exchange
	Exchange *string `json:"exchange,omitempty"`
	// Name: Company name
	Name string `json:"name"`
	// Sector: Industry sector
	Sector *string `json:"sector,omitempty"`
	// Summary: Last-7-days Polymarket summary for this asset
	Summary PolymarketSearchSummary `json:"summary"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// Type: Asset type (Stock, ETF, etc.)
	Type *string `json:"type,omitempty"`
}

// PolymarketSearchSummary: Compact last-7-days summary attached to Polymarket search results.
type PolymarketSearchSummary struct {
	// BuzzScore: Buzz score over the last 7 UTC calendar days
	BuzzScore *float64 `json:"buzz_score"`
	// SentimentScore: Orderbook-aware implied sentiment over the last 7 UTC calendar days
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// TradeCount: Trade count in the last 7 UTC calendar days
	TradeCount int64 `json:"trade_count"`
	// Trend: Flow momentum over the current 3 UTC days vs previous 3 UTC days using trades, volume, market breadth and liquidity; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
}

// PolymarketStatsResponse: Polymarket service statistics.
type PolymarketStatsResponse struct {
	// OpenMarketsCurrent: Currently open Polymarket markets where active=true, closed=false and end_date has not passed
	OpenMarketsCurrent *int64 `json:"open_markets_current,omitempty"`
	// OpenTickersCurrent: Distinct tickers with at least one currently open Polymarket market
	OpenTickersCurrent *int64 `json:"open_tickers_current,omitempty"`
	// SupportedTickers: Number of supported stock tickers
	SupportedTickers int64 `json:"supported_tickers"`
	// TotalMarkets: Distinct Polymarket condition_id count
	TotalMarkets int64 `json:"total_markets"`
	// TotalTrades: Total Polymarket trade activity from daily aggregates
	TotalTrades int64 `json:"total_trades"`
	// TradedMarketsToday: Distinct Polymarket condition_id values with UTC-day trade activity
	TradedMarketsToday *int64 `json:"traded_markets_today,omitempty"`
	// TradedTickersToday: Distinct tickers with UTC-day Polymarket trade activity
	TradedTickersToday *int64 `json:"traded_tickers_today,omitempty"`
	// TradesToday: Polymarket trade activity during today's UTC date
	TradesToday *int64 `json:"trades_today,omitempty"`
	// UniqueTickers: Distinct tickers with Polymarket market rows in the database
	UniqueTickers int64 `json:"unique_tickers"`
}

// PolymarketStockDetailResponse: Detailed Polymarket data for a single stock.
type PolymarketStockDetailResponse struct {
	// BearishPct: Outcome-aware bearish market percentage
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Outcome-aware bullish market percentage
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Buzz score (0-100)
	BuzzScore *float64 `json:"buzz_score,omitempty"`
	// CompanyName: Company name
	CompanyName *string `json:"company_name,omitempty"`
	// CurrentMarketCount: Number of currently open markets in the latest UTC-day snapshot for ticker; use this for live-only market breadth
	CurrentMarketCount *int64 `json:"current_market_count,omitempty"`
	// DailyTrend: Daily activity breakdown for the selected period
	DailyTrend []PolymarketDailyTrendItem `json:"daily_trend,omitempty"`
	// Found: Legacy v1 data-availability flag. `true` means this service has qualifying data for this asset/window. `false` means the asset is supported, but this service has no qualifying data for the requested window. Planned for removal in v2.
	// Deprecated: Retained for API compatibility.
	Found bool `json:"found"`
	// MarketCount: Distinct markets that existed for ticker within the selected UTC window, counted by market-catalog lifetime overlap
	MarketCount *int64 `json:"market_count,omitempty"`
	// NegativeCount: Deprecated. Outcome-aware bearish market count retained for v1 compatibility; prefer bearish_pct with trade_count and market_count for public analysis.
	// Deprecated: Retained for API compatibility.
	NegativeCount *int64 `json:"negative_count,omitempty"`
	// NeutralCount: Deprecated. Outcome-aware neutral/unclassified market count retained for v1 compatibility; prefer bullish_pct, bearish_pct, trade_count and market_count for public analysis.
	// Deprecated: Retained for API compatibility.
	NeutralCount *int64 `json:"neutral_count,omitempty"`
	// PeriodDays: Analysis period in days
	PeriodDays *int64 `json:"period_days,omitempty"`
	// PositiveCount: Deprecated. Outcome-aware bullish market count retained for v1 compatibility; prefer bullish_pct with trade_count and market_count for public analysis.
	// Deprecated: Retained for API compatibility.
	PositiveCount *int64 `json:"positive_count,omitempty"`
	// Pulse: Compact current Polymarket interpretation built from locally stored latest snapshot evidence. Does not duplicate top-level aggregate fields; concrete market evidence remains in `top_mentions`.
	Pulse *PolymarketTickerPulse `json:"pulse,omitempty"`
	// SentimentScore: Orderbook-aware implied sentiment; null when no directional sentiment evidence exists in the selected period
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// TopMentions: Top markets with an active/open snapshot during the retained snapshot portion of the selected UTC-day period for this ticker (including the current UTC day so far), capped at 10 and sorted as representative sentiment evidence: outcome-aware directional markets first, opposing signals kept visible when present, then trade count, with non-open markets de-prioritized and 24h volume and liquidity as tie-breakers. For windows longer than raw snapshot retention, ticker analytics still use daily aggregates, while this list is limited to retained market snapshots. Item prices and liquidity come from the selected in-period snapshot, not necessarily the current market state. `active` indicates whether the market is currently open, so historical in-period evidence can now have a non-open current state. `market_status` gives the compact current state.
	TopMentions []PolymarketTopMention `json:"top_mentions,omitempty"`
	// TotalLiquidity: Windowed aggregated liquidity signal in USD over the selected period; not a current snapshot
	TotalLiquidity *float64 `json:"total_liquidity,omitempty"`
	// TradeCount: Trade count in period
	TradeCount *int64 `json:"trade_count,omitempty"`
	// Trend: Flow momentum over the current 3 UTC days vs previous 3 UTC days using trades, volume, market breadth and liquidity; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
	// UniqueTraders: Exact distinct union of observed proxy and explicit maker wallet hashes for the delivered scope and UTC window; null when retained wallet-level trades do not fully cover the window; not a complete market-wide participant census
	UniqueTraders *int64 `json:"unique_traders,omitempty"`
}

// PolymarketTickerPulse: Compact interpretation layer for current ticker-level Polymarket data.
type PolymarketTickerPulse struct {
	// Confidence: Interpretation confidence capped by coverage, trade breadth, spread, liquidity and recency
	Confidence float64 `json:"confidence"`
	// Evidence: Compact non-duplicative quality summary for the pulse inputs
	Evidence PolymarketPulseEvidence `json:"evidence"`
	// Mood: Current compact Polymarket interpretation for this ticker Values: bullish, bearish, mixed, event_driven, illiquid, unclear.
	Mood string `json:"mood"`
	// ThinData: True when current evidence is insufficient for a strong directional interpretation
	ThinData bool `json:"thin_data"`
	// Warnings: Stable machine-readable data quality, freshness and tradability warnings
	Warnings []string `json:"warnings,omitempty"`
	// Why: Stable machine-readable reason codes explaining the pulse interpretation
	Why []string `json:"why"`
}

// PolymarketTopMention: Top Polymarket market for a ticker with current open/closed state.
type PolymarketTopMention struct {
	// Active: Whether the representative market state is open; snapshot-backed rows use current state, historical daily examples use stored representative state
	Active bool `json:"active"`
	// BuzzScore: Market-level activity buzz score for this selected in-period snapshot
	BuzzScore *float64 `json:"buzz_score,omitempty"`
	// ConditionID: Polymarket condition ID
	ConditionID string `json:"condition_id"`
	// EndDate: Market end date (ISO)
	EndDate *string `json:"end_date,omitempty"`
	// Liquidity: Selected in-period snapshot liquidity in USD; not necessarily live
	Liquidity float64 `json:"liquidity"`
	// MarketStatus: Compact market status for this representative item; snapshot-backed rows use current state, historical daily examples use stored representative state Values: tradable, open, resolved, inactive, expired.
	MarketStatus string `json:"market_status"`
	// MarketType: Derived market type
	MarketType string `json:"market_type"`
	// NoPrice: NO price from the selected in-period snapshot; not necessarily live
	NoPrice *float64 `json:"no_price,omitempty"`
	// Question: Market question/title
	Question string `json:"question"`
	// SentimentScore: Outcome-aware market price sentiment for the selected in-period snapshot; prefers orderbook midpoint when available, falls back to YES/NO price and returns neutral 0.0 when direction is unclassified
	SentimentScore float64 `json:"sentiment_score"`
	// TradeCount: UTC-day trade count from the selected in-period snapshot
	TradeCount *int64 `json:"trade_count,omitempty"`
	// Volume24h: Selected in-period snapshot rolling 24h volume in USD; not necessarily live
	Volume24h float64 `json:"volume_24h"`
	// YesPrice: YES price from the selected in-period snapshot; not necessarily live
	YesPrice *float64 `json:"yes_price,omitempty"`
}

// PolymarketTrendingCountry: Trending country on Polymarket.
type PolymarketTrendingCountry struct {
	// BearishPct: Outcome-aware bearish market percentage
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Outcome-aware bullish market percentage
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Aggregated activity/flow buzz score using trades, volume, breadth, liquidity and participation
	BuzzScore *float64 `json:"buzz_score"`
	// Country: Country name
	Country string `json:"country"`
	// CurrentMarketCount: Sum of ticker-level currently open markets in the latest UTC-day snapshot; use this for live-only market breadth
	CurrentMarketCount int64 `json:"current_market_count"`
	// MarketCount: Sum of ticker-level catalog lifetime-overlap market counts in the selected UTC window for this dimension; not a cross-ticker condition-id union
	MarketCount int64 `json:"market_count"`
	// SentimentScore: Weighted implied sentiment; null when no underlying trades exist in the selected period
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// TopTickers: Top 5 tickers by trade_count
	TopTickers []string `json:"top_tickers"`
	// TotalLiquidity: Windowed aggregated liquidity signal in USD over the selected period; not a current snapshot
	TotalLiquidity float64 `json:"total_liquidity"`
	// TradeCount: Total trade count in period
	TradeCount int64 `json:"trade_count"`
	// Trend: Flow momentum over the current 3 UTC days vs previous 3 UTC days using trades, volume, market breadth and liquidity; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Values: rising, falling, stable.
	Trend string `json:"trend"`
	// TrendHistory: Daily dimension buzz scores (oldest to newest). Length follows the resolved period, with a minimum of 7 values. Every slot describes only its UTC calendar day. The live final value describes today so far and can differ from period buzz_score. Closed historical windows end at `to`.
	TrendHistory []float64 `json:"trend_history"`
	// UniqueTickers: Unique tickers in dimension
	UniqueTickers int64 `json:"unique_tickers"`
	// UniqueTraders: Exact distinct union of observed proxy and explicit maker wallet hashes for the dimension and UTC window; null when retained wallet-level trades do not fully cover the window
	UniqueTraders *int64 `json:"unique_traders,omitempty"`
}

// PolymarketTrendingSector: Trending sector on Polymarket.
type PolymarketTrendingSector struct {
	// BearishPct: Outcome-aware bearish market percentage
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Outcome-aware bullish market percentage
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Aggregated activity/flow buzz score using trades, volume, breadth, liquidity and participation
	BuzzScore *float64 `json:"buzz_score"`
	// CurrentMarketCount: Sum of ticker-level currently open markets in the latest UTC-day snapshot; use this for live-only market breadth
	CurrentMarketCount int64 `json:"current_market_count"`
	// MarketCount: Sum of ticker-level catalog lifetime-overlap market counts in the selected UTC window for this dimension; not a cross-ticker condition-id union
	MarketCount int64 `json:"market_count"`
	// Sector: Sector name
	Sector string `json:"sector"`
	// SentimentScore: Weighted implied sentiment; null when no underlying trades exist in the selected period
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// TopTickers: Top 5 tickers by trade_count
	TopTickers []string `json:"top_tickers"`
	// TotalLiquidity: Windowed aggregated liquidity signal in USD over the selected period; not a current snapshot
	TotalLiquidity float64 `json:"total_liquidity"`
	// TradeCount: Total trade count in period
	TradeCount int64 `json:"trade_count"`
	// Trend: Flow momentum over the current 3 UTC days vs previous 3 UTC days using trades, volume, market breadth and liquidity; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Values: rising, falling, stable.
	Trend string `json:"trend"`
	// TrendHistory: Daily dimension buzz scores (oldest to newest). Length follows the resolved period, with a minimum of 7 values. Every slot describes only its UTC calendar day. The live final value describes today so far and can differ from period buzz_score. Closed historical windows end at `to`.
	TrendHistory []float64 `json:"trend_history"`
	// UniqueTickers: Unique tickers in dimension
	UniqueTickers int64 `json:"unique_tickers"`
	// UniqueTraders: Exact distinct union of observed proxy and explicit maker wallet hashes for the dimension and UTC window; null when retained wallet-level trades do not fully cover the window
	UniqueTraders *int64 `json:"unique_traders,omitempty"`
}

// PolymarketTrendingStock: Trending stock from Polymarket in Reddit/X-compatible shape.
type PolymarketTrendingStock struct {
	// BearishPct: Outcome-aware bearish market percentage
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Outcome-aware bullish market percentage
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Polymarket activity/flow discovery score (0-100); not directional sentiment
	BuzzScore *float64 `json:"buzz_score"`
	// CompanyName: Company name
	CompanyName *string `json:"company_name,omitempty"`
	// CurrentMarketCount: Number of currently open markets in the latest UTC-day snapshot for ticker; use this for live-only market breadth
	CurrentMarketCount int64 `json:"current_market_count"`
	// MarketCount: Distinct markets that existed for ticker within the selected UTC window, counted by market-catalog lifetime overlap
	MarketCount int64 `json:"market_count"`
	// SentimentScore: Orderbook-aware implied sentiment (-1 to +1); null when no directional sentiment evidence exists in the selected period
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// TotalLiquidity: Windowed aggregated liquidity signal in USD over the selected period; not a current snapshot
	TotalLiquidity float64 `json:"total_liquidity"`
	// TradeCount: Trade count in selected period
	TradeCount int64 `json:"trade_count"`
	// Trend: Flow momentum over the current 3 UTC days vs previous 3 UTC days using trades, volume, market breadth and liquidity; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Values: rising, falling, stable.
	Trend string `json:"trend"`
	// TrendHistory: Recent buzz history (oldest→newest). Length follows the resolved period, with a minimum of 7 values. Closed historical windows end at `to`.
	TrendHistory []float64 `json:"trend_history,omitempty"`
	// UniqueTraders: Exact distinct union of observed proxy and explicit maker wallet hashes for the delivered scope and UTC window; null when retained wallet-level trades do not fully cover the window; not a complete market-wide participant census
	UniqueTraders *int64 `json:"unique_traders,omitempty"`
}

// RateLimitError: Error response for rate-limit exceeded responses.
type RateLimitError struct {
	// Detail: Structured rate-limit details
	Detail RateLimitErrorDetail `json:"detail"`
}

// RateLimitErrorDetail: Structured rate-limit error detail.
type RateLimitErrorDetail struct {
	// AccountType: Resolved account type or fallback tier
	AccountType string `json:"account_type"`
	// Contact: Plan or support URL for resolving the limit
	Contact *string `json:"contact,omitempty"`
	// Error: Error type identifier
	Error string `json:"error"`
	// Limit: Applied request limit for the window
	Limit *int64 `json:"limit,omitempty"`
	// Message: Human-readable error message
	Message string `json:"message"`
	// Remaining: Requests remaining in the current window
	Remaining *int64 `json:"remaining,omitempty"`
	// ResetAt: ISO 8601 timestamp when the current window resets
	ResetAt *string `json:"reset_at,omitempty"`
	// Used: Requests used in the current window
	Used *int64 `json:"used,omitempty"`
}

// RedditMarketSentimentDriver: Top asset contributing to the service-level Reddit market sentiment.
type RedditMarketSentimentDriver struct {
	// BuzzScore: Asset buzz score in the selected period
	BuzzScore *float64 `json:"buzz_score"`
	// Mentions: Attributed mentions in the selected period, including inherited thread-context mentions
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score for this asset in the selected period
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
}

// RedditMarketSentimentResponse: Service-level Reddit market sentiment across all tracked stocks.
type RedditMarketSentimentResponse struct {
	// ActiveTickers: Number of tickers with qualifying attributed Reddit activity in the selected period
	ActiveTickers int64 `json:"active_tickers"`
	// BearishPct: Bearish mention percentage
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Bullish mention percentage
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Service-wide Reddit heat score relative to the service's trailing 90-day baseline. Around 50 = normal activity, higher values = hotter-than-usual service activity.
	BuzzScore *float64 `json:"buzz_score"`
	// Drivers: Top 5 stocks of the selected period by buzz_score. mentions and sentiment_score cover every day of the period.
	Drivers []RedditMarketSentimentDriver `json:"drivers,omitempty"`
	// Mentions: Total Reddit mentions across the selected period, including inherited thread-context mentions
	Mentions int64 `json:"mentions"`
	// NegativeCount: Deprecated. Prefer bearish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NegativeCount int64 `json:"negative_count"`
	// NeutralCount: Deprecated. Prefer bullish_pct, bearish_pct and mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NeutralCount int64 `json:"neutral_count"`
	// PositiveCount: Deprecated. Prefer bullish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	PositiveCount int64 `json:"positive_count"`
	// SentimentScore: Service-wide average sentiment score across all qualifying mentions in the selected period
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// SubredditCount: Distinct subreddits contributing in the selected period
	SubredditCount int64 `json:"subreddit_count"`
	// TotalUpvotes: Total upvotes across attributed mentions in the selected period
	TotalUpvotes int64 `json:"total_upvotes"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, upvotes and subreddit breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
	// TrendHistory: Daily service-wide buzz scores (oldest→newest) using the same relative baseline calibration. Length = max(requested_days, 7) for successful requests; windows before platform availability return HTTP 422.
	TrendHistory []float64 `json:"trend_history,omitempty"`
	// UniquePosts: Exact distinct Reddit post count across the selected UTC period
	UniquePosts int64 `json:"unique_posts"`
}

// RedditRawMentionItem: Raw Reddit mention row for one ticker.
type RedditRawMentionItem struct {
	// Author: Public Reddit author name, if scraped
	Author *string `json:"author,omitempty"`
	// CommentID: Reddit comment identifier when the mention comes from a comment
	CommentID *string `json:"comment_id,omitempty"`
	// CreatedUTC: UTC timestamp of creation (ISO 8601) Format: date-time.
	CreatedUTC string `json:"created_utc"`
	// IsInherited: Whether ticker context was inherited from the parent thread instead of directly mentioned
	IsInherited bool `json:"is_inherited"`
	// PostID: Reddit post identifier
	PostID *string `json:"post_id,omitempty"`
	// SentimentLabel: Sentiment label (positive/negative/neutral)
	SentimentLabel *string `json:"sentiment_label,omitempty"`
	// SentimentScore: Sentiment score (-1 to +1)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Subreddit: Source subreddit
	Subreddit *string `json:"subreddit,omitempty"`
	// TextSnippet: Stored mention snippet (truncated at ingestion time)
	TextSnippet string `json:"text_snippet"`
	// Upvotes: Reddit score (upvotes) for this mention row
	Upvotes int64 `json:"upvotes"`
}

// RedditRawMentionsResponse: Paginated raw Reddit mentions for a single ticker.
type RedditRawMentionsResponse struct {
	// Count: Total number of matching raw rows for the requested include_inherited filter before limit and offset are applied; only the requested page is returned
	Count int64 `json:"count"`
	// PeriodDays: Lookback window used for the query
	PeriodDays int64 `json:"period_days"`
	// Results: Raw mention rows ordered by newest first
	Results []RedditRawMentionItem `json:"results"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
}

// RedditStatsResponse: Service statistics.
type RedditStatsResponse struct {
	// MentionsToday: Stock mention rows created since today's UTC midnight, including posts, comments and inherited thread-context mentions
	MentionsToday *int64 `json:"mentions_today,omitempty"`
	// SupportedTickers: Number of supported stock tickers
	SupportedTickers int64 `json:"supported_tickers"`
	// TotalMentions: All-time number of stored Reddit stock mentions
	TotalMentions int64 `json:"total_mentions"`
	// UniqueTickers: Number of unique tickers with mentions
	UniqueTickers int64 `json:"unique_tickers"`
	// UniqueTickersToday: Unique tickers (detected using case-sensitive, uppercase-only patterns) including posts, comments and inherited thread-context mentions
	UniqueTickersToday *int64 `json:"unique_tickers_today,omitempty"`
}

// RootHealthResponse: Root API health status aggregated across platforms.
type RootHealthResponse struct {
	// Service: Root service identifier
	Service *string `json:"service,omitempty"`
	// Services: Per-platform health payloads keyed by service identifier
	Services map[string]map[string]json.RawMessage `json:"services"`
	// Status: Aggregated root health status Values: healthy, unhealthy.
	Status string `json:"status"`
	// Summary: Aggregated health counts
	Summary RootHealthSummary `json:"summary"`
	// Version: Service version
	Version string `json:"version"`
}

// RootHealthSummary: Aggregated root health summary.
type RootHealthSummary struct {
	// Healthy: Number of checks that reported healthy
	Healthy int64 `json:"healthy"`
	// Total: Number of platform health checks
	Total int64 `json:"total"`
	// Unhealthy: Number of checks that reported unhealthy or timed out
	Unhealthy int64 `json:"unhealthy"`
	// UnhealthyServices: Service identifiers that reported unhealthy or timed out
	UnhealthyServices []string `json:"unhealthy_services"`
}

// SchedulerStatus: Scheduler status information.
type SchedulerStatus struct {
	// ErrorCount: Total number of scrape errors/failures
	ErrorCount *int64 `json:"error_count,omitempty"`
	// LastScrape: ISO timestamp of last scrape (null if never run)
	LastScrape *string `json:"last_scrape,omitempty"`
	// Note: Additional status information, for example why collection is not running
	Note *string `json:"note,omitempty"`
	// Running: Whether data collection is running
	Running bool `json:"running"`
	// ScrapeCount: Total number of scrapes performed
	ScrapeCount *int64 `json:"scrape_count,omitempty"`
	// SuccessRate: Scrape success rate (0.0-1.0)
	SuccessRate *float64 `json:"success_rate,omitempty"`
}

// SearchResponse: Search results for stocks.
type SearchResponse struct {
	// Count: Total number of matching results before the limit is applied
	Count int64 `json:"count"`
	// PeriodDays: Fixed lookback window used for each summary block
	PeriodDays int64 `json:"period_days"`
	// Query: Original search query
	Query string `json:"query"`
	// Results: Top matching stocks capped by the requested limit
	Results []SearchResultItem `json:"results"`
}

// SearchResultItem: Individual search result item.
type SearchResultItem struct {
	// Aliases: Alternative names/aliases
	Aliases []string `json:"aliases,omitempty"`
	// Country: Country of headquarters
	Country *string `json:"country,omitempty"`
	// Exchange: Stock exchange (NYSE, NASDAQ, etc.)
	Exchange *string `json:"exchange,omitempty"`
	// Name: Company name
	Name string `json:"name"`
	// Sector: Industry sector
	Sector *string `json:"sector,omitempty"`
	// Summary: Last-7-days Reddit summary for this asset
	Summary StockSearchSummary `json:"summary"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// Type: Asset type (Stock, ETF, etc.)
	Type *string `json:"type,omitempty"`
}

// SentimentAnalyzeRequest: Direct text sentiment analysis request.
type SentimentAnalyzeRequest struct {
	// Target: Optional: the company to judge the text for, as a name or ticker, for example `Apple` or `AAPL`. It should appear in the text; a company name works best. "Goldman Sachs cuts Apple price target" is negative for `Apple` and neutral for `Goldman Sachs`.
	Target *string `json:"target,omitempty"`
	// Text: Financial text to analyze, for example a news headline or a sentence about a company.
	Text string `json:"text"`
}

// SentimentAnalyzeResponse: Direct text sentiment analysis response.
type SentimentAnalyzeResponse struct {
	// Components: Whitelisted finance sentiment engine components
	Components SentimentComponents `json:"components"`
	// SentimentLabel: Adanos three-class sentiment label, matching top_mentions payloads Values: positive, neutral, negative.
	SentimentLabel string `json:"sentiment_label"`
	// SentimentScore: -1.0 bearish to +1.0 bullish
	SentimentScore float64 `json:"sentiment_score"`
	// Text: Normalized input text that was analyzed
	Text string `json:"text"`
}

// SentimentComponents: Whitelisted sentiment engine components exposed publicly.
type SentimentComponents struct {
	// ContextualFinanceMatches: Engine 5.4.1 only; since engine 5.5 it is empty. Contextual finance phrases that informed scoring.
	// Deprecated: Retained for API compatibility.
	ContextualFinanceMatches []string `json:"contextual_finance_matches,omitempty"`
	// EmojiScore: Engine 5.4.1 only; since engine 5.5 it is 0. Emoji-derived sentiment score.
	// Deprecated: Retained for API compatibility.
	EmojiScore float64 `json:"emoji_score"`
	// EngineVersion: Sentiment engine version
	EngineVersion string `json:"engine_version"`
	// ModelScore: Raw model score from -1 to +1: the probability of positive minus the probability of negative. `sentiment_score` maps it onto the shared scale, where ±0.14 separates positive and negative from neutral.
	ModelScore *float64 `json:"model_score,omitempty"`
	// PhraseAdjustment: Engine 5.4.1 only; since engine 5.5 it is 0. Directional finance phrase adjustment.
	// Deprecated: Retained for API compatibility.
	PhraseAdjustment float64 `json:"phrase_adjustment"`
	// PhraseMatches: Engine 5.4.1 only; since engine 5.5 it is empty. Directional finance phrases matched by the engine.
	// Deprecated: Retained for API compatibility.
	PhraseMatches []SentimentPhraseMatch `json:"phrase_matches,omitempty"`
	// RobertaScore: Engine 5.4.1 only; since engine 5.5 it is null; see `model_score`. RoBERTa ensemble score.
	// Deprecated: Retained for API compatibility.
	RobertaScore *float64 `json:"roberta_score,omitempty"`
	// VaderCompound: Engine 5.4.1 only; since engine 5.5 it is null. VADER compound score after finance lexicon tuning.
	// Deprecated: Retained for API compatibility.
	VaderCompound *float64 `json:"vader_compound,omitempty"`
}

// SentimentPhraseMatch: Whitelisted directional finance phrase metadata.
type SentimentPhraseMatch struct {
	// Phrase: Matched finance phrase
	Phrase string `json:"phrase"`
	// Score: Internal phrase sentiment score
	Score float64 `json:"score"`
}

// StockExplanationResponse: AI-generated trend explanation for a stock.
type StockExplanationResponse struct {
	// Cached: Whether the explanation was served from cache
	Cached bool `json:"cached"`
	// CompanyName: Company name
	CompanyName *string `json:"company_name,omitempty"`
	// Explanation: AI-generated explanation of why this stock is trending
	Explanation string `json:"explanation"`
	// GeneratedAt: Timezone-aware UTC ISO timestamp when explanation was generated
	GeneratedAt string `json:"generated_at"`
	// Model: LLM model used for generation
	Model *string `json:"model,omitempty"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
}

// StockSearchSummary: Compact last-7-days summary attached to stock search results.
type StockSearchSummary struct {
	// BuzzScore: Buzz score over the last 7 UTC calendar days
	BuzzScore *float64 `json:"buzz_score"`
	// Mentions: Attributed mentions in the last 7 UTC calendar days, including inherited thread-context mentions
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score over the last 7 UTC calendar days
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, upvotes and subreddit breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
}

// StockSentiment: Detailed sentiment analysis for a specific stock ticker.
type StockSentiment struct {
	// BearishPct: Percentage of bearish mentions
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Percentage of bullish mentions
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Buzz Score (0-100). Asymptotic scaling above 50.
	BuzzScore *float64 `json:"buzz_score,omitempty"`
	// CompanyName: Company name (null if unknown)
	CompanyName *string `json:"company_name,omitempty"`
	// DailyTrend: Daily attributed mention breakdown for the selected period, with sentiment and buzz_score
	DailyTrend []DailyTrendItem `json:"daily_trend,omitempty"`
	// Found: Legacy v1 data-availability flag. `true` means this service has qualifying data for this asset/window. `false` means the asset is supported, but this service has no qualifying data for the requested window. Planned for removal in v2.
	// Deprecated: Retained for API compatibility.
	Found bool `json:"found"`
	// Mentions: Total attributed mentions in the selected period, including inherited thread-context mentions
	Mentions *int64 `json:"mentions,omitempty"`
	// NegativeCount: Deprecated. Prefer bearish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NegativeCount *int64 `json:"negative_count,omitempty"`
	// NeutralCount: Deprecated. Prefer bullish_pct, bearish_pct and mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NeutralCount *int64 `json:"neutral_count,omitempty"`
	// PeriodDays: Analysis period in days
	PeriodDays *int64 `json:"period_days,omitempty"`
	// PositiveCount: Deprecated. Prefer bullish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	PositiveCount *int64 `json:"positive_count,omitempty"`
	// SentimentScore: Average sentiment score (-1 bearish to +1 bullish)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// SubredditCount: Number of subreddits with mentions
	SubredditCount *int64 `json:"subreddit_count,omitempty"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// TopMentions: Top explicit mentions from the selected period, ranked by upvotes
	TopMentions []TopMention `json:"top_mentions,omitempty"`
	// TopSubreddits: Top subreddits by attributed mention count
	TopSubreddits []SubredditCount `json:"top_subreddits,omitempty"`
	// TotalUpvotes: Total upvotes across attributed mentions
	TotalUpvotes *int64 `json:"total_upvotes,omitempty"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, upvotes and subreddit breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
	// UniquePosts: Number of unique posts
	UniquePosts *int64 `json:"unique_posts,omitempty"`
}

// SubredditCount: Subreddit contributor metrics for a stock detail response.
type SubredditCount struct {
	// BuzzScore: Contributor-level buzz score for this subreddit in the selected period
	BuzzScore *float64 `json:"buzz_score,omitempty"`
	// Count: Deprecated alias for mentions. Use mentions instead.
	// Deprecated: Retained for API compatibility.
	Count int64 `json:"count"`
	// Mentions: Attributed mentions in this subreddit
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score for this subreddit in the selected period
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Subreddit: Subreddit name
	Subreddit string `json:"subreddit"`
}

// TopMention: Top mention by upvotes.
type TopMention struct {
	// CreatedUTC: UTC timestamp of creation (ISO 8601) Format: date-time.
	CreatedUTC string `json:"created_utc"`
	// SentimentLabel: Sentiment label (positive/negative/neutral)
	SentimentLabel string `json:"sentiment_label"`
	// SentimentScore: Sentiment score (-1 to +1)
	SentimentScore float64 `json:"sentiment_score"`
	// Subreddit: Source subreddit
	Subreddit string `json:"subreddit"`
	// TextSnippet: Clean Reddit paragraph around the detected ticker mention (up to 500 chars)
	TextSnippet string `json:"text_snippet"`
	// Upvotes: Number of upvotes
	Upvotes int64 `json:"upvotes"`
}

// TrendingCountry: Trending country with aggregated buzz metrics from Reddit.
type TrendingCountry struct {
	// BearishPct: Percentage of sentiment-labeled mentions that are bearish
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Percentage of sentiment-labeled mentions that are bullish
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Aggregated Buzz Score (0-100)
	BuzzScore float64 `json:"buzz_score"`
	// Country: Country name
	Country string `json:"country"`
	// Mentions: Total attributed mentions, including inherited thread-context mentions
	Mentions int64 `json:"mentions"`
	// SentimentScore: Weighted average sentiment (-1 bearish to +1 bullish)
	SentimentScore float64 `json:"sentiment_score"`
	// SubredditCount: Number of subreddits with mentions
	SubredditCount int64 `json:"subreddit_count"`
	// TopTickers: Top 5 tickers by mentions
	TopTickers []string `json:"top_tickers"`
	// TotalUpvotes: Total upvotes across attributed mentions
	TotalUpvotes int64 `json:"total_upvotes"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, upvotes and subreddit breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Values: rising, falling, stable.
	Trend string `json:"trend"`
	// TrendHistory: Daily dimension buzz scores (oldest to newest). Length follows the resolved period, with a minimum of 7 values. Every slot describes only its UTC calendar day. The live final value describes today so far and can differ from period buzz_score. Closed historical windows end at `to`.
	TrendHistory []float64 `json:"trend_history"`
	// UniqueTickers: Number of unique tickers mentioned
	UniqueTickers int64 `json:"unique_tickers"`
}

// TrendingSector: Trending sector with aggregated buzz metrics from Reddit.
type TrendingSector struct {
	// BearishPct: Percentage of sentiment-labeled mentions that are bearish
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Percentage of sentiment-labeled mentions that are bullish
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Aggregated Buzz Score (0-100)
	BuzzScore float64 `json:"buzz_score"`
	// Mentions: Total attributed mentions, including inherited thread-context mentions
	Mentions int64 `json:"mentions"`
	// Sector: Industry sector name
	Sector string `json:"sector"`
	// SentimentScore: Weighted average sentiment (-1 bearish to +1 bullish)
	SentimentScore float64 `json:"sentiment_score"`
	// SubredditCount: Number of subreddits with mentions
	SubredditCount int64 `json:"subreddit_count"`
	// TopTickers: Top 5 tickers by mentions
	TopTickers []string `json:"top_tickers"`
	// TotalUpvotes: Total upvotes across attributed mentions
	TotalUpvotes int64 `json:"total_upvotes"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, upvotes and subreddit breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Values: rising, falling, stable.
	Trend string `json:"trend"`
	// TrendHistory: Daily dimension buzz scores (oldest to newest). Length follows the resolved period, with a minimum of 7 values. Every slot describes only its UTC calendar day. The live final value describes today so far and can differ from period buzz_score. Closed historical windows end at `to`.
	TrendHistory []float64 `json:"trend_history"`
	// UniqueTickers: Number of unique tickers mentioned
	UniqueTickers int64 `json:"unique_tickers"`
}

// TrendingStock: Trending stock with buzz metrics from Reddit.
type TrendingStock struct {
	// BearishPct: Percentage of bearish mentions
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Percentage of bullish mentions
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Buzz Score (0-100). Asymptotic scaling above 50.
	BuzzScore float64 `json:"buzz_score"`
	// CompanyName: Company name (null if unknown)
	CompanyName *string `json:"company_name,omitempty"`
	// Mentions: Total attributed mentions, including inherited thread-context mentions
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score (-1 bearish to +1 bullish)
	SentimentScore float64 `json:"sentiment_score"`
	// SubredditCount: Number of subreddits with mentions
	SubredditCount int64 `json:"subreddit_count"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// TotalUpvotes: Total upvotes across attributed mentions
	TotalUpvotes int64 `json:"total_upvotes"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, upvotes and subreddit breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Values: rising, falling, stable.
	Trend string `json:"trend"`
	// TrendHistory: Daily buzz scores (oldest→newest). Length follows the resolved period, with a minimum of 7 values. Values 0-100.
	TrendHistory []float64 `json:"trend_history,omitempty"`
	// UniquePosts: Number of unique posts
	UniquePosts int64 `json:"unique_posts"`
}

// UnsupportedAssetError: Error response for unsupported v1 stock/crypto asset requests.
type UnsupportedAssetError struct {
	// Detail: Structured unsupported asset details
	Detail UnsupportedAssetErrorDetail `json:"detail"`
}

// UnsupportedAssetErrorDetail: Structured detail for unsupported v1 stock/crypto asset requests.
type UnsupportedAssetErrorDetail struct {
	// ErrorCode: Machine-readable unsupported asset error code
	ErrorCode string `json:"error_code"`
	// Message: Human-readable unsupported asset message
	Message string `json:"message"`
}

// ValidationError: FastAPI/Pydantic validation error item.
type ValidationError struct {
	// Loc: Location of the invalid input
	Loc []json.RawMessage `json:"loc"`
	// Msg: Validation failure message
	Msg string `json:"msg"`
	// Type: Validation error type
	Type string `json:"type"`
}

// XCompareResponse: Comparison response for multiple X/Twitter stocks.
type XCompareResponse struct {
	// PeriodDays: Analysis period in days
	PeriodDays int64 `json:"period_days"`
	// Stocks: Stocks sorted by buzz_score descending
	Stocks []XCompareStockItem `json:"stocks"`
}

// XCompareStockItem: Individual stock in X/Twitter comparison.
type XCompareStockItem struct {
	// BearishPct: Percentage of bearish tweet mentions
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Percentage of bullish tweet mentions
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Aggregated Buzz Score (0-100) from X mentions
	BuzzScore *float64 `json:"buzz_score"`
	// CompanyName: Company name
	CompanyName *string `json:"company_name,omitempty"`
	// Mentions: Total tweet mentions
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score from tweet analysis (-1 to +1)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// TotalUpvotes: Total likes across all tweet mentions
	TotalUpvotes int64 `json:"total_upvotes"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days, not price movement. Null when no qualifying X data exists.
	Trend *string `json:"trend,omitempty"`
	// TrendHistory: Daily buzz history (oldest to newest). Length follows the resolved period, with a minimum of 7 values when data exists. Closed historical windows end at `to`. Empty when no qualifying data exists.
	TrendHistory []float64 `json:"trend_history,omitempty"`
	// UniqueTweets: Number of unique tweets mentioning this ticker (distinct tweet_id)
	UniqueTweets int64 `json:"unique_tweets"`
}

// XDailyTrendItem: Daily trend data point for X/Twitter.
type XDailyTrendItem struct {
	// BearishPct: Bearish tweet percentage for this date, calculated from negative sentiment-classified tweets
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Bullish tweet percentage for this date, calculated from positive sentiment-classified tweets
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Buzz score for this date (0-100); null until the day is aggregated, while the activity fields are already filled
	BuzzScore *float64 `json:"buzz_score,omitempty"`
	// Date: Date in YYYY-MM-DD format
	Date string `json:"date"`
	// Mentions: Number of tweet mentions on this date
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score for the day
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
}

// XHealthResponse: X/Twitter service health status - aligned with Reddit health response.
type XHealthResponse struct {
	// Error: Error message if unhealthy
	Error *string `json:"error,omitempty"`
	// Scheduler: Data collection status
	Scheduler *XSchedulerStatus `json:"scheduler,omitempty"`
	// Service: Service identifier
	Service string `json:"service"`
	// Status: Current health status (healthy/unhealthy)
	Status string `json:"status"`
	// TickersTracked: Distinct tickers observed in the most recent scrape cycle; returns 0 when scrape telemetry is unavailable.
	TickersTracked int64 `json:"tickers_tracked"`
	// TotalMentions: Mentions observed in the most recent scrape cycle; returns 0 when scrape telemetry is unavailable.
	TotalMentions int64 `json:"total_mentions"`
	// Version: Current API version
	Version string `json:"version"`
}

// XMarketSentimentDriver: Top asset contributing to the service-level X market sentiment.
type XMarketSentimentDriver struct {
	// BuzzScore: Asset buzz score in the selected period
	BuzzScore *float64 `json:"buzz_score"`
	// Mentions: Tweet mentions in the selected period
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average tweet sentiment score for this asset
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
}

// XMarketSentimentResponse: Service-level X/Twitter market sentiment across all tracked stocks.
type XMarketSentimentResponse struct {
	// ActiveTickers: Number of tickers with X activity in the selected period
	ActiveTickers int64 `json:"active_tickers"`
	// BearishPct: Bearish tweet percentage
	BearishPct int64 `json:"bearish_pct"`
	// BullishPct: Bullish tweet percentage
	BullishPct int64 `json:"bullish_pct"`
	// BuzzScore: Service-wide X/Twitter heat score relative to the service's trailing 90-day baseline. Around 50 = normal activity, higher values = hotter-than-usual X activity.
	BuzzScore *float64 `json:"buzz_score"`
	// Drivers: Top 5 stocks of the selected period by buzz_score. mentions and sentiment_score cover every day of the period.
	Drivers []XMarketSentimentDriver `json:"drivers,omitempty"`
	// Mentions: Tweet mentions across the selected period
	Mentions int64 `json:"mentions"`
	// NegativeCount: Deprecated. Prefer bearish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NegativeCount int64 `json:"negative_count"`
	// NeutralCount: Deprecated. Prefer bullish_pct, bearish_pct and mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NeutralCount int64 `json:"neutral_count"`
	// PositiveCount: Deprecated. Prefer bullish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	PositiveCount int64 `json:"positive_count"`
	// SentimentScore: Service-wide average tweet sentiment score
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// TotalUpvotes: Total likes across the selected period
	TotalUpvotes int64 `json:"total_upvotes"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, likes and author breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
	// TrendHistory: Daily service-wide buzz scores (oldest→newest) using the same relative baseline calibration. Length = max(requested_days, 7) for successful requests; windows before platform availability return HTTP 422.
	TrendHistory []float64 `json:"trend_history,omitempty"`
	// UniqueAuthors: Exact distinct author breadth across the selected UTC period
	UniqueAuthors int64 `json:"unique_authors"`
	// UniqueTweets: Exact distinct tweet count across the selected UTC period
	UniqueTweets int64 `json:"unique_tweets"`
}

// XRawMentionItem: Raw tweet mention row for one stock ticker.
type XRawMentionItem struct {
	// Author: Public X/Twitter username, if scraped
	Author *string `json:"author,omitempty"`
	// CreatedUTC: Tweet creation time in UTC (ISO 8601) when available
	CreatedUTC *string `json:"created_utc,omitempty"`
	// IsReply: Whether this row comes from a reply tweet
	IsReply bool `json:"is_reply"`
	// Likes: Number of likes
	Likes int64 `json:"likes"`
	// ParentTweetID: Parent tweet identifier for replies
	ParentTweetID *string `json:"parent_tweet_id,omitempty"`
	// Retweets: Number of retweets
	Retweets int64 `json:"retweets"`
	// SentimentLabel: Sentiment label (positive/negative/neutral)
	SentimentLabel *string `json:"sentiment_label,omitempty"`
	// SentimentScore: Sentiment score (-1 to +1)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// TextSnippet: Tweet text snippet
	TextSnippet string `json:"text_snippet"`
	// TweetID: Tweet identifier
	TweetID string `json:"tweet_id"`
	// Views: Number of views
	Views *int64 `json:"views,omitempty"`
}

// XRawMentionsResponse: Paginated raw tweet mentions for a single stock ticker.
type XRawMentionsResponse struct {
	// Count: Total number of matching raw rows before limit and offset are applied; only the requested page is returned
	Count int64 `json:"count"`
	// PeriodDays: Lookback window used for the query
	PeriodDays int64 `json:"period_days"`
	// Results: Raw tweet rows ordered by newest first
	Results []XRawMentionItem `json:"results"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
}

// XSchedulerStatus: X/Twitter data collection status.
type XSchedulerStatus struct {
	// ErrorCount: Total failed scrape cycles
	ErrorCount *int64 `json:"error_count,omitempty"`
	// LastScrape: ISO timestamp of last completed scrape cycle
	LastScrape *string `json:"last_scrape,omitempty"`
	// Note: Additional status info (e.g., worker crash warning)
	Note *string `json:"note,omitempty"`
	// Running: Whether data collection is running (reported within the last 90 minutes)
	Running bool `json:"running"`
	// ScrapeCount: Total completed scrape cycles since worker start
	ScrapeCount *int64 `json:"scrape_count,omitempty"`
	// SuccessRate: Success rate (success_count / scrape_count)
	SuccessRate *float64 `json:"success_rate,omitempty"`
}

// XSearchResponse: Search response for X/Twitter stocks.
type XSearchResponse struct {
	// Count: Total number of matching results before the limit is applied
	Count int64 `json:"count"`
	// PeriodDays: Fixed lookback window used for each summary block
	PeriodDays int64 `json:"period_days"`
	// Query: Original search query
	Query string `json:"query"`
	// Results: Top matching stocks capped by the requested limit
	Results []XSearchResultItem `json:"results"`
}

// XSearchResultItem: Individual X/Twitter search result item.
type XSearchResultItem struct {
	// Aliases: Alternative names/aliases
	Aliases []string `json:"aliases,omitempty"`
	// Country: Country of headquarters
	Country *string `json:"country,omitempty"`
	// Exchange: Stock exchange
	Exchange *string `json:"exchange,omitempty"`
	// Name: Company name
	Name string `json:"name"`
	// Sector: Industry sector
	Sector *string `json:"sector,omitempty"`
	// Summary: Last-7-days X summary for this asset
	Summary XSearchSummary `json:"summary"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// Type: Asset type (Stock, ETF, etc.)
	Type *string `json:"type,omitempty"`
}

// XSearchSummary: Compact last-7-days summary attached to X search results.
type XSearchSummary struct {
	// BuzzScore: Buzz score over the last 7 UTC calendar days
	BuzzScore *float64 `json:"buzz_score"`
	// Mentions: Tweet mentions in the last 7 UTC calendar days
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score over the last 7 UTC calendar days
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, likes and author breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
}

// XStatsResponse: X/Twitter service statistics.
type XStatsResponse struct {
	// MentionsToday: X mention rows created since today's UTC midnight
	MentionsToday *int64 `json:"mentions_today,omitempty"`
	// SupportedTickers: Number of supported stock tickers
	SupportedTickers *int64 `json:"supported_tickers,omitempty"`
	// TotalMentions: All-time number of stored X mentions
	TotalMentions *int64 `json:"total_mentions,omitempty"`
	// UniqueTickers: Number of unique tickers with X mentions in the database
	UniqueTickers *int64 `json:"unique_tickers,omitempty"`
	// UniqueTickersToday: Unique tickers with X mentions since today's UTC midnight
	UniqueTickersToday *int64 `json:"unique_tickers_today,omitempty"`
}

// XStockDetailResponse: Detailed X/Twitter data for a single stock, in the same format as the Reddit API.
type XStockDetailResponse struct {
	// BearishPct: Percentage of bearish tweet mentions
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Percentage of bullish tweet mentions
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Buzz score (0-100) from tweet activity, including author diversity (null if not found)
	BuzzScore *float64 `json:"buzz_score,omitempty"`
	// CompanyName: Company name
	CompanyName *string `json:"company_name,omitempty"`
	// DailyTrend: Daily mention breakdown for the selected period
	DailyTrend []XDailyTrendItem `json:"daily_trend,omitempty"`
	// Found: Legacy v1 data-availability flag. `true` means this service has qualifying data for this asset/window. `false` means the asset is supported, but this service has no qualifying data for the requested window. Planned for removal in v2.
	// Deprecated: Retained for API compatibility.
	Found *bool `json:"found,omitempty"`
	// Mentions: Total tweet mentions within the selected period
	Mentions *int64 `json:"mentions,omitempty"`
	// NegativeCount: Deprecated. Prefer bearish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NegativeCount *int64 `json:"negative_count,omitempty"`
	// NeutralCount: Deprecated. Prefer bullish_pct, bearish_pct and mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	NeutralCount *int64 `json:"neutral_count,omitempty"`
	// PeriodDays: Analysis period in days
	PeriodDays *int64 `json:"period_days,omitempty"`
	// PositiveCount: Deprecated. Prefer bullish_pct with mentions for public analysis; retained for v1 compatibility.
	// Deprecated: Retained for API compatibility.
	PositiveCount *int64 `json:"positive_count,omitempty"`
	// SentimentScore: Average sentiment score from tweet analysis (-1 to +1)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// TopAuthors: Top X authors in the selected period with mentions, sentiment_score, buzz_score and deprecated count alias
	TopAuthors []XTopAuthor `json:"top_authors,omitempty"`
	// TopTweets: Top tweets from the selected period, ranked by engagement (likes + retweets)
	TopTweets []XTopTweet `json:"top_tweets,omitempty"`
	// TotalUpvotes: Total likes across all tweet mentions
	TotalUpvotes *int64 `json:"total_upvotes,omitempty"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, likes and author breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today).
	Trend *string `json:"trend,omitempty"`
	// UniqueTweets: Number of unique tweets (distinct tweet_id)
	UniqueTweets *int64 `json:"unique_tweets,omitempty"`
}

// XTopAuthor: X author contributor metrics for a stock detail response.
type XTopAuthor struct {
	// Author: Public X/Twitter username
	Author string `json:"author"`
	// BuzzScore: Contributor-level buzz score for this author in the selected period
	BuzzScore *float64 `json:"buzz_score,omitempty"`
	// Count: Deprecated alias for mentions. Use mentions instead.
	// Deprecated: Retained for API compatibility.
	Count int64 `json:"count"`
	// Mentions: Tweet mentions by this author in the selected period
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score for this author's mentions in the selected period
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
}

// XTopTweet: Top tweet mention for a stock - highest engagement tweets.
type XTopTweet struct {
	// Author: Author username
	Author *string `json:"author,omitempty"`
	// CreatedAt: Tweet creation time (ISO 8601)
	CreatedAt *string `json:"created_at,omitempty"`
	// Likes: Number of likes
	Likes int64 `json:"likes"`
	// Retweets: Number of retweets
	Retweets int64 `json:"retweets"`
	// SentimentLabel: Sentiment label (positive/negative/neutral)
	SentimentLabel *string `json:"sentiment_label,omitempty"`
	// SentimentScore: Sentiment score (-1 to +1)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// TextSnippet: Clean tweet paragraph around the detected ticker mention (truncated to 280 chars for this view)
	TextSnippet string `json:"text_snippet"`
	// Views: Number of views
	Views *int64 `json:"views,omitempty"`
}

// XTrendingCountry: Trending country with aggregated buzz metrics from X/Twitter.
type XTrendingCountry struct {
	// BearishPct: Percentage of bearish mentions
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Percentage of bullish mentions
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Aggregated Buzz Score (0-100) from X mentions
	BuzzScore *float64 `json:"buzz_score"`
	// Country: Country name
	Country string `json:"country"`
	// Mentions: Total tweet mentions
	Mentions int64 `json:"mentions"`
	// SentimentScore: Weighted average sentiment (-1 bearish to +1 bullish)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// TopTickers: Top 5 tickers by mention count in this dimension
	TopTickers []string `json:"top_tickers"`
	// TotalUpvotes: Total likes across all mentions
	TotalUpvotes *int64 `json:"total_upvotes,omitempty"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, likes and author breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Values: rising, falling, stable.
	Trend string `json:"trend"`
	// TrendHistory: Daily dimension buzz scores (oldest to newest). Length follows the resolved period, with a minimum of 7 values. Every slot describes only its UTC calendar day. The live final value describes today so far and can differ from period buzz_score. Closed historical windows end at `to`. Historical author breadth is rebuilt from daily author_distribution where available and may conservatively approximate exact cross-ticker author de-duplication.
	TrendHistory []float64 `json:"trend_history"`
	// UniqueAuthors: Exact distinct author breadth across the selected UTC period
	UniqueAuthors int64 `json:"unique_authors"`
	// UniqueTickers: Number of unique tickers mentioned
	UniqueTickers int64 `json:"unique_tickers"`
}

// XTrendingSector: Trending sector with aggregated buzz metrics from X/Twitter.
type XTrendingSector struct {
	// BearishPct: Percentage of bearish mentions
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Percentage of bullish mentions
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Aggregated Buzz Score (0-100) from X mentions
	BuzzScore *float64 `json:"buzz_score"`
	// Mentions: Total tweet mentions
	Mentions int64 `json:"mentions"`
	// Sector: Industry sector name
	Sector string `json:"sector"`
	// SentimentScore: Weighted average sentiment (-1 bearish to +1 bullish)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// TopTickers: Top 5 tickers by mention count in this dimension
	TopTickers []string `json:"top_tickers"`
	// TotalUpvotes: Total likes across all mentions
	TotalUpvotes *int64 `json:"total_upvotes,omitempty"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days using mentions, likes and author breadth; not price movement. For `from`/`to`, anchors at `to` (or now when `to` is today). Values: rising, falling, stable.
	Trend string `json:"trend"`
	// TrendHistory: Daily dimension buzz scores (oldest to newest). Length follows the resolved period, with a minimum of 7 values. Every slot describes only its UTC calendar day. The live final value describes today so far and can differ from period buzz_score. Closed historical windows end at `to`. Historical author breadth is rebuilt from daily author_distribution where available and may conservatively approximate exact cross-ticker author de-duplication.
	TrendHistory []float64 `json:"trend_history"`
	// UniqueAuthors: Exact distinct author breadth across the selected UTC period
	UniqueAuthors int64 `json:"unique_authors"`
	// UniqueTickers: Number of unique tickers mentioned
	UniqueTickers int64 `json:"unique_tickers"`
}

// XTrendingStock: Trending stock on X/Twitter, in the same format as the Reddit API. Metrics come from tweets and replies that mention the ticker. Author diversity rewards discussion spread across many accounts.
type XTrendingStock struct {
	// BearishPct: Percentage of bearish tweet mentions
	BearishPct *int64 `json:"bearish_pct,omitempty"`
	// BullishPct: Percentage of bullish tweet mentions
	BullishPct *int64 `json:"bullish_pct,omitempty"`
	// BuzzScore: Buzz score from tweet activity. Components: mentions (20), sentiment (20), quality (10), author diversity (14) and trend (-10 to +20). Asymptotic scaling caps at 100.
	BuzzScore *float64 `json:"buzz_score"`
	// CompanyName: Company name
	CompanyName *string `json:"company_name,omitempty"`
	// Mentions: Number of tweets mentioning the ticker
	Mentions int64 `json:"mentions"`
	// SentimentScore: Average sentiment score from tweet analysis (-1 to +1)
	SentimentScore *float64 `json:"sentiment_score,omitempty"`
	// Ticker: Stock ticker symbol
	Ticker string `json:"ticker"`
	// TotalUpvotes: Total likes across all tweet mentions
	TotalUpvotes *int64 `json:"total_upvotes,omitempty"`
	// Trend: Activity momentum over the current 3 UTC days vs previous 3 UTC days, not price movement. Values: rising, falling, stable.
	Trend string `json:"trend"`
	// TrendHistory: Daily buzz scores (oldest→newest). Length follows the resolved period, with a minimum of 7 values. For live windows, the final value is the current live buzz_score; closed historical windows end at `to`.
	TrendHistory []float64 `json:"trend_history,omitempty"`
	// UniqueTweets: Number of unique tweets mentioning this ticker (distinct tweet_id)
	UniqueTweets *int64 `json:"unique_tweets,omitempty"`
}
