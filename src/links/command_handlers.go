package links

import (
	"context"
	"time"
)

type LinkCommandHandler struct {
	linkRepository ILinkRepository
}

func NewLinkCommandHandler(linkRepository ILinkRepository) *LinkCommandHandler {
	return &LinkCommandHandler{
		linkRepository: linkRepository,
	}
}

func (h *LinkCommandHandler) CreateShortenedUrl(ctx context.Context, command CreateShortenedUrlCommand) (*CreateShortenedUrlCommandResponse, error) {
	newLink := &Link{OriginLink: command.OriginUrl, ShortenedLink: nil, ExpiresAt: time.Now()}

	err := h.linkRepository.Save(ctx, newLink)
	if err != nil {
		return nil, err
	}

	// todo: refactor

	fetchedNewLink, err := h.linkRepository.FindByOriginLink(ctx, command.OriginUrl)
	if err != nil {
		return nil, err
	}
	fetchedNewLink.ShortenedLink = md5Encode(fetchedNewLink.ID)

	err = h.linkRepository.Save(ctx, fetchedNewLink)
	if err != nil {
		return nil, err
	}
	err = h.linkRepository.SetUrl(ctx, command.getLinkKey(), *fetchedNewLink.ShortenedLink, time.Minute*2)
	if err != nil {
		return nil, err
	}
	err = h.linkRepository.SetUrlExpiresAt(ctx, command.getExpirationKey(), fetchedNewLink.ExpiresAt, time.Minute*2)
	if err != nil {
		return nil, err
	}
	response := &CreateShortenedUrlCommandResponse{ShortenedUrl: *fetchedNewLink.ShortenedLink, ExpiresAt: fetchedNewLink.ExpiresAt}
	return response, nil
}
