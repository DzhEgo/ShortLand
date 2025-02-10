package link

import (
	"ShortLand/internal/model"
	"crypto/sha256"
	"strconv"
	"time"
)

const alphabet = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789_"

var store = make(map[string]model.ShortLinkData)

func CreateShortLink(link string) string {
	i := 0

	for {
		salt := strconv.Itoa(i)
		short := generateShortLink(link, salt)

		data, exist := store[short]
		if !exist {
			store[short] = model.ShortLinkData{
				LongLink: link,
				ExpireAt: time.Now().Add(5 * time.Minute),
			}
			return short
		}

		if data.LongLink == link {
			return short
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

	if time.Now().After(data.ExpireAt) {
		clean(link)
		return ""
	}

	return data.LongLink
}

func clean(link string) {
	delete(store, link)
}
