package model

import "time"

type ShortLinkData struct {
	LongLink string    `json:"long_link"`
	ExpireAt time.Time `json:"expire_at"`
}

type Command struct {
	Link string `json:"link"`
}
