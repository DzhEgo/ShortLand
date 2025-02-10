package model

import "time"

type ShortLinkData struct {
	LongLink string    `json:"long_link"`
	ExpireAt time.Time `json:"expire_at"`
}

type Command struct {
	Link string `json:"link"`
}

type LinkTable struct {
	ID         int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	OriginLink string `json:"origin_link" gorm:"column:origin_link"`
	ShortLink  string `json:"short_link" gorm:"column:short_link"`
	ExpireAt   int64  `json:"expire_at" gorm:"column:expire_at"`
}
