package service

import (
	"ShortLand/internal/agent/link"
	"ShortLand/internal/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

type LinkService interface {
	CreateShortLink(origLink string) (string, error)
	GetOriginalLink(shortLink string) (string, error)
}

type linkService struct {
	db *gorm.DB
}

func NewLinkService() LinkService {
	dsn := fmt.Sprintf("host=localhost user=postgres dbname=postgres sslmode=disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	return &linkService{
		db: db,
	}
}

func (s *linkService) CreateShortLink(origLink string) (string, error) {
	var data model.LinkTable

	if origLink == "" {
		return "", fmt.Errorf("invalid link")
	}

	data = link.CreateShortLink(origLink)
	if data.OriginLink == "" {
		return "", fmt.Errorf("invalid link")
	}

	if err := s.db.Save(&data).Error; err != nil {
		return "", err
	}

	return data.ShortLink, nil
}

func (s *linkService) GetOriginalLink(shortLink string) (string, error) {
	var originalLink string

	if shortLink == "" || strings.HasPrefix(shortLink, "http://") || strings.HasPrefix(shortLink, "https://") {
		return "", fmt.Errorf("invalid link")
	}

	if err := s.db.Model(&model.LinkTable{}).Select("origin_link").Where("short_link = ?", shortLink).Take(&originalLink).Error; err != nil {
		return "", err
	}

	originalLink = link.GetShortLink(shortLink)
	if originalLink == "" {
		return "", fmt.Errorf("origin link not found")
	}

	return originalLink, nil
}
