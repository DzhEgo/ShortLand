package link

import (
	"ShortLand/internal/common/model"
	"crypto/sha256"
	"fmt"
	"github.com/gofrs/uuid"
	"strings"
	"time"
)

const alphabet = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789_"

type LinkService interface {
	CreateShortLink(origLink string) (string, error)
	GetOriginalLink(shortLink string) (string, error)
}

type linkService struct {
	stor LinkStorage
	salt string
}

func NewLinkService(stor LinkStorage) LinkService {
	return &linkService{
		stor: stor,
	}
}

func (s *linkService) CreateShortLink(origLink string) (string, error) {
	if origLink == "" {
		return "", fmt.Errorf("invalid link")
	}

	short, err := s.createShortLink(origLink)
	if short == "" || err != nil {
		return "", fmt.Errorf("failed to create short link")
	}

	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	data := model.LinkTable{
		ID:         id.String(),
		OriginLink: origLink,
		ShortLink:  short,
		ExpireAt:   time.Now().Add(time.Minute * 2).Unix(),
	}

	if err := s.stor.SaveLink(data); err != nil {
		return "", fmt.Errorf("failed to save short link")
	}

	return data.ShortLink, nil
}

func (s *linkService) GetOriginalLink(shortLink string) (string, error) {
	var data *model.LinkTable

	if shortLink == "" || strings.HasPrefix(shortLink, "http://") || strings.HasPrefix(shortLink, "https://") {
		return "", fmt.Errorf("invalid link")
	}

	data, err := s.stor.GetLink(shortLink)
	if err != nil {
		return "", fmt.Errorf("failed to get short link")
	}

	if time.Now().Unix() > data.ExpireAt {
		if err := s.stor.DeleteLink(shortLink); err != nil {
			return "", fmt.Errorf("failed to delete short link")
		}

		return "", fmt.Errorf("short link expired")
	}

	return data.OriginLink, nil
}

func (s *linkService) createShortLink(link string) (string, error) {
	short := generateShortLink(link, s.salt)

	data, err := s.stor.GetLink(short)
	if err != nil {
		return short, nil
	} else {
		if time.Now().Unix() > data.ExpireAt {
			if err := s.stor.DeleteLink(short); err != nil {
				return "", fmt.Errorf("failed to delete short link")
			}
			return short, nil
		}
	}

	return "", fmt.Errorf("short link exist")
}

func generateShortLink(link, salt string) string {
	hash := sha256.New()

	hash.Write([]byte(link + salt))
	sum := hash.Sum(nil)
	short := make([]byte, 10)

	for i := 0; i < 10; i++ {
		short[i] = alphabet[sum[i]%byte(len(alphabet))]
	}

	return string(short)
}
