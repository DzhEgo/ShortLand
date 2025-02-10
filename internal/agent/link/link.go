package link

import (
	"ShortLand/internal/model"
	"crypto/sha256"
	"strconv"
	"time"
)

const alphabet = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789_"

var store = make(map[string]model.LinkTable)

func CreateShortLink(link string) model.LinkTable {
	i := 0

	for {
		salt := strconv.Itoa(i)
		short := generateShortLink(link, salt)

		data, exist := store[short]
		if !exist {
			store[short] = model.LinkTable{
				OriginLink: link,
				ShortLink:  short,
				ExpireAt:   time.Now().Add(5 * time.Minute).Unix(),
			}

			return store[short]
		}

		if data.OriginLink == link {
			return store[short]
		}

		i++
	}
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

func GetShortLink(link string) string {
	data, exist := store[link]
	if !exist {
		return ""
	}

	if time.Now().Unix() > data.ExpireAt {
		clean(link)
		return ""
	}

	return data.OriginLink
}

func clean(link string) {
	delete(store, link)
}
