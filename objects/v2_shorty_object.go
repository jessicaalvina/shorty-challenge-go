package objects

type V2ShortyObjectResponse struct {
	Shortcode	string `json:"shortcode"`
}

type V2ShortyObjectRequest struct {
	Url  string `json:"url"`
	Shortcode string `json:"shortcode"`
}
