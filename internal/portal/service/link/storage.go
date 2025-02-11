package link

import "ShortLand/internal/common/model"

type StorageLink interface {
	SaveLink(data model.LinkTable) error
	GetLink(shortLink string) (*model.LinkTable, error)
	DeleteLink(shortLink string) error
}
