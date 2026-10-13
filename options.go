package adanos

import (
	"net/url"
	"strconv"
)

// PeriodOptions selects the query parameters sent to the API. Nil fields are omitted.
type PeriodOptions struct {
	From *string
	To   *string
	// Days is the legacy date-window parameter retained for SDK compatibility.
	Days *int
}

func (o *PeriodOptions) query() url.Values {
	q := url.Values{}
	if o == nil {
		return q
	}
	if o.From != nil {
		q.Set("from", *o.From)
	}
	if o.To != nil {
		q.Set("to", *o.To)
	}
	if o.Days != nil {
		q.Set("days", strconv.Itoa(*o.Days))
	}
	return q
}

// TrendingOptions selects the query parameters sent to the API. Nil fields are omitted.
type TrendingOptions struct {
	From *string
	To   *string
	// Days is the legacy date-window parameter retained for SDK compatibility.
	Days   *int
	Limit  *int
	Offset *int
	Type   *string
}

func (o *TrendingOptions) query() url.Values {
	q := url.Values{}
	if o == nil {
		return q
	}
	if o.From != nil {
		q.Set("from", *o.From)
	}
	if o.To != nil {
		q.Set("to", *o.To)
	}
	if o.Days != nil {
		q.Set("days", strconv.Itoa(*o.Days))
	}
	if o.Limit != nil {
		q.Set("limit", strconv.Itoa(*o.Limit))
	}
	if o.Offset != nil {
		q.Set("offset", strconv.Itoa(*o.Offset))
	}
	if o.Type != nil {
		q.Set("type", *o.Type)
	}
	return q
}

// TrendingGroupOptions selects the query parameters sent to the API. Nil fields are omitted.
type TrendingGroupOptions struct {
	From *string
	To   *string
	// Days is the legacy date-window parameter retained for SDK compatibility.
	Days   *int
	Limit  *int
	Offset *int
}

func (o *TrendingGroupOptions) query() url.Values {
	q := url.Values{}
	if o == nil {
		return q
	}
	if o.From != nil {
		q.Set("from", *o.From)
	}
	if o.To != nil {
		q.Set("to", *o.To)
	}
	if o.Days != nil {
		q.Set("days", strconv.Itoa(*o.Days))
	}
	if o.Limit != nil {
		q.Set("limit", strconv.Itoa(*o.Limit))
	}
	if o.Offset != nil {
		q.Set("offset", strconv.Itoa(*o.Offset))
	}
	return q
}

// NewsTrendingOptions selects the query parameters sent to the API. Nil fields are omitted.
type NewsTrendingOptions struct {
	From *string
	To   *string
	// Days is the legacy date-window parameter retained for SDK compatibility.
	Days   *int
	Limit  *int
	Offset *int
	Type   *string
	Source *string
}

func (o *NewsTrendingOptions) query() url.Values {
	q := url.Values{}
	if o == nil {
		return q
	}
	if o.From != nil {
		q.Set("from", *o.From)
	}
	if o.To != nil {
		q.Set("to", *o.To)
	}
	if o.Days != nil {
		q.Set("days", strconv.Itoa(*o.Days))
	}
	if o.Limit != nil {
		q.Set("limit", strconv.Itoa(*o.Limit))
	}
	if o.Offset != nil {
		q.Set("offset", strconv.Itoa(*o.Offset))
	}
	if o.Type != nil {
		q.Set("type", *o.Type)
	}
	if o.Source != nil {
		q.Set("source", *o.Source)
	}
	return q
}

// NewsTrendingGroupOptions selects the query parameters sent to the API. Nil fields are omitted.
type NewsTrendingGroupOptions struct {
	From *string
	To   *string
	// Days is the legacy date-window parameter retained for SDK compatibility.
	Days   *int
	Limit  *int
	Offset *int
	Source *string
}

func (o *NewsTrendingGroupOptions) query() url.Values {
	q := url.Values{}
	if o == nil {
		return q
	}
	if o.From != nil {
		q.Set("from", *o.From)
	}
	if o.To != nil {
		q.Set("to", *o.To)
	}
	if o.Days != nil {
		q.Set("days", strconv.Itoa(*o.Days))
	}
	if o.Limit != nil {
		q.Set("limit", strconv.Itoa(*o.Limit))
	}
	if o.Offset != nil {
		q.Set("offset", strconv.Itoa(*o.Offset))
	}
	if o.Source != nil {
		q.Set("source", *o.Source)
	}
	return q
}

// MentionsOptions selects the query parameters sent to the API. Nil fields are omitted.
type MentionsOptions struct {
	From *string
	To   *string
	// Days is the legacy date-window parameter retained for SDK compatibility.
	Days   *int
	Limit  *int
	Offset *int
}

func (o *MentionsOptions) query() url.Values {
	q := url.Values{}
	if o == nil {
		return q
	}
	if o.From != nil {
		q.Set("from", *o.From)
	}
	if o.To != nil {
		q.Set("to", *o.To)
	}
	if o.Days != nil {
		q.Set("days", strconv.Itoa(*o.Days))
	}
	if o.Limit != nil {
		q.Set("limit", strconv.Itoa(*o.Limit))
	}
	if o.Offset != nil {
		q.Set("offset", strconv.Itoa(*o.Offset))
	}
	return q
}

// RedditMentionsOptions selects the query parameters sent to the API. Nil fields are omitted.
type RedditMentionsOptions struct {
	From *string
	To   *string
	// Days is the legacy date-window parameter retained for SDK compatibility.
	Days             *int
	Limit            *int
	Offset           *int
	IncludeInherited *bool
}

func (o *RedditMentionsOptions) query() url.Values {
	q := url.Values{}
	if o == nil {
		return q
	}
	if o.From != nil {
		q.Set("from", *o.From)
	}
	if o.To != nil {
		q.Set("to", *o.To)
	}
	if o.Days != nil {
		q.Set("days", strconv.Itoa(*o.Days))
	}
	if o.Limit != nil {
		q.Set("limit", strconv.Itoa(*o.Limit))
	}
	if o.Offset != nil {
		q.Set("offset", strconv.Itoa(*o.Offset))
	}
	if o.IncludeInherited != nil {
		q.Set("include_inherited", strconv.FormatBool(*o.IncludeInherited))
	}
	return q
}

// SearchOptions selects the query parameters sent to the API. Nil fields are omitted.
type SearchOptions struct {
	Limit *int
}

func (o *SearchOptions) query() url.Values {
	q := url.Values{}
	if o == nil {
		return q
	}
	if o.Limit != nil {
		q.Set("limit", strconv.Itoa(*o.Limit))
	}
	return q
}

// SentimentAnalyzeOptions selects an optional analysis target.
type SentimentAnalyzeOptions struct{ Target *string }
