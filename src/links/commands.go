package links

import "time"

type CreateShortenedUrlCommand struct {
	OriginUrl string
}

type CreateShortenedUrlCommandResponse struct {
	ShortenedUrl string
	ExpiresAt    time.Time
}

func (c CreateShortenedUrlCommand) getLinkKey() string {
	return "origin-link " + c.OriginUrl
}

func (c CreateShortenedUrlCommand) getExpirationKey() string {
	return "origin-link-expires-at " + c.OriginUrl
}
