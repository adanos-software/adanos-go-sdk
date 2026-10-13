# Changelog

## 1.0.0

Initial Go SDK release for Adanos API 1.54.0.

- Covers all 53 operations across Reddit Stocks, Reddit Crypto, X, News, Polymarket, direct sentiment analysis, and root health.
- Includes all 113 typed schemas, generated from a frozen OpenAPI contract.
- Preserves nullable scores, signed Crypto upvotes, News ticker suffixes, explicit query values, and legacy date windows.
- Supports context cancellation and inspectable HTTP errors. Blocks credential forwarding across origins.
- Uses the Go standard library. Requires Go 1.26 or newer.
