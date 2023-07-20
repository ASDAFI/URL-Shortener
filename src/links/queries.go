package links

import "time"

type GetShortenedLinkQuery struct {
	OriginLink string
}

type GetOriginLinkQuery struct {
	ShortenedLink string
}

type GetShortenedLinkQueryResponse struct {
	ShortenedLink string
	ExpiresAt     time.Time
}

type GetOriginLinkQueryResponse struct {
	OriginLink string
	ExpiresAt  time.Time
}

func (q GetShortenedLinkQuery) getLinkKey() string {
	return "origin-link " + q.OriginLink
}

func (q GetOriginLinkQuery) getLinkKey() string {
	return "shortened-link " + q.ShortenedLink
}

func (q GetShortenedLinkQuery) getExpirationKey() string {
	return "origin-link-expires-at " + q.OriginLink
}

func (q GetOriginLinkQuery) getExpirationKey() string {
	return "shortened-link-expires-at " + q.ShortenedLink
}
