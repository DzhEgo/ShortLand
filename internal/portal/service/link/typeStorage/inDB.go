package typeStorage

import (
	"ShortLand/internal/common/model"
	"fmt"
	"gorm.io/gorm"
)

type InDB struct {
	db *gorm.DB
}

func NewInDB(db *gorm.DB) *InDB {
	return &InDB{
		db: db,
	}
}

func (d *InDB) SaveLink(data model.LinkTable) error {
	return d.db.Model(&model.LinkTable{}).Create(&data).Error
}

func (d *InDB) GetLink(shortLink string) (*model.LinkTable, error) {
	var link model.LinkTable

	if err := d.db.Model(&model.LinkTable{}).Where("short_link = ?", shortLink).Take(&link).Error; err != nil {
		return nil, fmt.Errorf("failed to get link")
	}

	return &link, nil
}

func (d *InDB) DeleteLink(shortLink string) error {
	return d.db.Where("short_link = ?", shortLink).Delete(&model.LinkTable{}).Error
}
