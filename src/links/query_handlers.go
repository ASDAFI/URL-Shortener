package links

import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"
)

type LinkQueryHandler struct {
	linkRepository ILinkRepository
}

func NewLinkQueryHandler(linkRepository ILinkRepository) *LinkQueryHandler {
	return &LinkQueryHandler{
		linkRepository: linkRepository,
	}
}

func (h *LinkQueryHandler) GetShortenedLink(ctx context.Context, query GetShortenedLinkQuery) (*GetShortenedLinkQueryResponse, error) {
	isCached, err := h.linkRepository.IsExistKey(ctx, query.getLinkKey())
	if err != nil {
		return nil, err
	}

	if isCached {
		status := true
		log.Info("is cached")
		link, err := h.linkRepository.GetUrl(ctx, query.getLinkKey())
		if err != nil {
			status = false
			log.Info("Some cache error!", err)
		}
		if link == nil && status {
			log.Info("Some cache error!", err)
			status = false
		}
		if status {
			linkExp, err := h.linkRepository.GetUrlExpiresAt(ctx, query.getExpirationKey())
			if err != nil || linkExp == nil {
				log.Info("Some cache error!", err)

				status = false
			}

			if status {
				response := &GetShortenedLinkQueryResponse{ShortenedLink: *link, ExpiresAt: *linkExp}
				return response, nil
			}
		}
	}

	link, err := h.linkRepository.FindByOriginLink(ctx, query.OriginLink)
	if err != nil {
		return nil, err
	}
	if link == nil {
		return nil, nil
	}
	if link.ShortenedLink == nil {
		return nil, nil //errors.New("Shortened link does not exist!")
	}
	err = h.linkRepository.SetUrl(ctx, query.getLinkKey(), *link.ShortenedLink, time.Minute*2)
	if err != nil {
		return nil, err
	}
	err = h.linkRepository.SetUrlExpiresAt(ctx, query.getExpirationKey(), link.ExpiresAt, time.Minute*2)
	if err != nil {
		return nil, err
	}
	response := &GetShortenedLinkQueryResponse{ShortenedLink: *link.ShortenedLink, ExpiresAt: link.ExpiresAt}
	return response, nil
}

func (h *LinkQueryHandler) GetOriginLink(ctx context.Context, query GetOriginLinkQuery) (*GetOriginLinkQueryResponse, error) {
	isCached, err := h.linkRepository.IsExistKey(ctx, query.getLinkKey())
	if err != nil {
		return nil, err
	}

	if isCached {
		link, err := h.linkRepository.GetUrl(ctx, query.getLinkKey())
		if err != nil {
			return nil, err
		}

		linkExp, err := h.linkRepository.GetUrlExpiresAt(ctx, query.getExpirationKey())
		if err != nil {
			return nil, err
		}
		response := &GetOriginLinkQueryResponse{OriginLink: *link, ExpiresAt: *linkExp}
		return response, nil
	}
	link, err := h.linkRepository.FindByShortenedLink(ctx, &query.ShortenedLink)
	if err != nil {
		return nil, err
	}
	if link == nil {
		return nil, nil
	}

	err = h.linkRepository.SetUrl(ctx, query.getLinkKey(), link.OriginLink, time.Minute*2)
	if err != nil {
		return nil, err
	}
	err = h.linkRepository.SetUrlExpiresAt(ctx, query.getExpirationKey(), link.ExpiresAt, time.Minute*2)
	if err != nil {
		return nil, err
	}

	response := &GetOriginLinkQueryResponse{OriginLink: link.OriginLink, ExpiresAt: link.ExpiresAt}
	return response, nil
}
