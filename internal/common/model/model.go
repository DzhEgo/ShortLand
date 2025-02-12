package model

type Command struct {
	Link string `json:"link"`
}

type LinkTable struct {
	ID         string `json:"id" gorm:"primary_key"`
	OriginLink string `json:"origin_link" gorm:"column:origin_link"`
	ShortLink  string `json:"short_link" gorm:"column:short_link;unique"`
	ExpireAt   int64  `json:"expire_at" gorm:"column:expire_at"`
}

type LinkOutput struct {
	Link     string `json:"link"`
	ExpireAt int64  `json:"expire_at"`
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
