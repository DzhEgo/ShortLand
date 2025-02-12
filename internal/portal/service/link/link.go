package link

import (
	"ShortLand/internal/common/model"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"strings"
	"time"
)

const alphabet = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789_"

var Expired = errors.New("short link expired")
var Exist = errors.New("short link exist")
var Invalid = errors.New("invalid link")
var NotFound = errors.New("failed to find link")

type LinkService interface {
	CreateShortLink(origLink string) (*model.LinkOutput, error)
	GetOriginalLink(shortLink string) (*model.LinkOutput, error)
}

type linkService struct {
	stor StorageLink
}

func NewLinkService(stor StorageLink) LinkService {
	return &linkService{
		stor: stor,
	}
}

func (s *linkService) CreateShortLink(origLink string) (*model.LinkOutput, error) {
	var out *model.LinkOutput

	if origLink == "" || !(strings.HasPrefix(origLink, "http://") || strings.HasPrefix(origLink, "https://")) {
		return nil, Invalid
	}

	short, err := s.createShortLink(origLink)
	if short == "" || err != nil {
		return nil, err
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	expireTime := time.Now().Add(time.Minute * 2).Unix()
	data := model.LinkTable{
		ID:         id.String(),
		OriginLink: origLink,
		ShortLink:  short,
		ExpireAt:   expireTime,
	}

	if err := s.stor.SaveLink(data); err != nil {
		return nil, fmt.Errorf("failed to save short link: %w", err)
	}

	out = &model.LinkOutput{
		Link:     short,
		ExpireAt: expireTime,
	}

	return out, nil
}

func (s *linkService) GetOriginalLink(shortLink string) (*model.LinkOutput, error) {
	var out *model.LinkOutput

	if shortLink == "" || strings.HasPrefix(shortLink, "http://") || strings.HasPrefix(shortLink, "https://") {
		return nil, Invalid
	}

	data, err := s.stor.GetLink(shortLink)
	if err != nil {
		return nil, NotFound
	}

	if time.Now().Unix() > data.ExpireAt {
		if err := s.stor.DeleteLink(shortLink); err != nil {
			return nil, fmt.Errorf("failed to delete short link")
		}

		return nil, Expired
	}

	out = &model.LinkOutput{
		Link:     data.OriginLink,
		ExpireAt: data.ExpireAt,
	}

	return out, nil
}

func (s *linkService) createShortLink(link string) (string, error) {
	short := generateShortLink(link)

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

	return "", Exist
}

func generateShortLink(link string) string {
	hash := sha256.New()

	hash.Write([]byte(link))
	sum := hash.Sum(nil)
	short := make([]byte, 10)

	for i := 0; i < 10; i++ {
		short[i] = alphabet[sum[i]%byte(len(alphabet))]
	}

	return string(short)
}
