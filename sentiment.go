package adanos

import (
	"context"
	"net/http"
)

// SentimentService provides Sentiment API operations.
type SentimentService struct{ client *Client }

// Analyze calls the /sentiment/v1 Analyze endpoint.
func (s *SentimentService) Analyze(ctx context.Context, text string, o *SentimentAnalyzeOptions) (*SentimentAnalyzeResponse, error) {
	request := SentimentAnalyzeRequest{Text: text}
	if o != nil {
		request.Target = o.Target
	}
	return execute[*SentimentAnalyzeResponse](s.client, ctx, http.MethodPost, "/sentiment/v1/analyze", nil, request)
}
