package service

import (
	"ShortLand/internal/agent/link"
	"fmt"
	"strings"
)

type LinkService interface {
	CreateShortLink(origLink string) (string, error)
	GetOriginalLink(shortLink string) (string, error)
}

type linkService struct {
}

func NewLinkService() LinkService {
	return &linkService{}
}

func (s *linkService) CreateShortLink(origLink string) (string, error) {
	var shortLink string

	if origLink == "" {
		return "", fmt.Errorf("invalid link")
	}

	shortLink = link.CreateShortLink(origLink)

	return shortLink, nil
}

func (s *linkService) GetOriginalLink(shortLink string) (string, error) {
	var originalLink string

	if shortLink == "" || strings.HasPrefix(shortLink, "http://") || strings.HasPrefix(shortLink, "https://") {
		return "", fmt.Errorf("invalid link")
	}

	originalLink = link.GetShortLink(shortLink)
	if originalLink == "" {
		return "", fmt.Errorf("origin link not found")
	}

	return originalLink, nil
}
