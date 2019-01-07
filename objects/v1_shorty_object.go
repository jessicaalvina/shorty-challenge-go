package objects

import "time"

type V1ShortyObjectResponse struct {
	StartDate	time.Time `json:"start_date"`
	LastSeenDate	time.Time `json:"last_seen_date"`
	Url	string `json:"url"`
	RedirectCount	uint `json:"redirect_count"`
}

type V1ShortyObjectResponseNoLastSeen struct {
	StartDate	time.Time `json:"start_date"`
	Url	string `json:"url"`
	RedirectCount	uint `json:"redirect_count"`
}

type V1ShortyObjectRequest struct {
	Shortcode string `json:"shortcode"`
}
