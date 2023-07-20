package links

import "time"

type GetShortenedUrlRequest struct {
	OriginLink string `json:"origin_link"`
}

type GetShortenedUrlResponse struct {
	ShortenedLink *string    `json:"shortened_link,omitempty"`
	ExpiresAt     *time.Time `json:"expires_at,omitempty"`
	Message       string     `json:"message,omitempty"`
}

type GetOriginUrlRequest struct {
	ShortenedLink string `json:"shortened_link"`
}

type GetOriginUrlResponse struct {
	OriginLink *string    `json:"origin_link,omitempty"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty"`
	Message    string     `json:"message,omitempty"`
}
