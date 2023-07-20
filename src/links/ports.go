package links

import (
	"context"
	"time"
)

type ILinkRepository interface {
	Save(ctx context.Context, link *Link) error
	FindByShortenedLink(ctx context.Context, shortenedLink *string) (*Link, error)
	FindByOriginLink(ctx context.Context, originLink string) (*Link, error)
	IsExistKey(ctx context.Context, key string) (bool, error)
	GetUrl(ctx context.Context, key string) (*string, error)
	SetUrl(ctx context.Context, key string, value string, expiration time.Duration) error
	GetUrlExpiresAt(ctx context.Context, key string) (*time.Time, error)
	SetUrlExpiresAt(ctx context.Context, key string, value time.Time, expiration time.Duration) error
}
