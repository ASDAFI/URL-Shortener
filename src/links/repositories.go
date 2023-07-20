package links

import (
	"context"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"time"
	"url-shortener/src/infrastructure/cache"
	"url-shortener/src/infrastructure/db"
)

type LinkRepository struct {
	dBInfrastructure    db.Provider
	cacheInfrastructure cache.Provider
}

func NewLinkRepository(dBInfrastructure db.Provider, cacheInfrastructure cache.Provider) *LinkRepository {
	return &LinkRepository{dBInfrastructure: dBInfrastructure, cacheInfrastructure: cacheInfrastructure}
}

func (r *LinkRepository) Save(ctx context.Context, link *Link) error {

	return r.dBInfrastructure.DB.WithContext(ctx).Save(link).Error
}

func (r *LinkRepository) FindByShortenedLink(ctx context.Context, shortenedLink *string) (*Link, error) {
	link := &Link{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("shortened_link = ?", shortenedLink).Find(link).Error
	return link, err
}

func (r *LinkRepository) FindByOriginLink(ctx context.Context, originLink string) (*Link, error) {
	link := &Link{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("origin_link = ?", originLink).Find(link).Error
	return link, err
}

func (r *LinkRepository) IsExistKey(ctx context.Context, key string) (bool, error) {

	exists, err := r.cacheInfrastructure.Client.Exists(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	if exists == 0 {
		return false, nil
	}
	return true, nil

}

func (r *LinkRepository) GetUrl(ctx context.Context, key string) (*string, error) {

	result, err := r.cacheInfrastructure.Client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	return &result, nil

}

func (r *LinkRepository) GetUrlExpiresAt(ctx context.Context, key string) (*time.Time, error) {
	result, err := r.cacheInfrastructure.Client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	parsedTime, err := time.Parse(time.RFC3339, result) // Convert string to time.Time
	return &parsedTime, err
}

func (r *LinkRepository) SetUrl(ctx context.Context, key string, value string, expiration time.Duration) error {
	log.Info("hiii")
	err := r.cacheInfrastructure.Client.Set(ctx, key, value, expiration).Err()
	if err != nil && err != redis.Nil {
		return err
	}
	return nil

}

func (r *LinkRepository) SetUrlExpiresAt(ctx context.Context, key string, value time.Time, expiration time.Duration) error {
	strValue := value.Format(time.RFC3339) // Convert time.Time to string
	err := r.cacheInfrastructure.Client.Set(ctx, key, strValue, expiration).Err()
	return err
}
