package db

import (
	"errors"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

const shortCodeLength = 6
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func CreateShortURL(db *gorm.DB, originalURL string) (*ShortURL, error) {
	for {
		code := randomString(shortCodeLength)
		url := ShortURL{
			OriginalURL: originalURL,
			ShortCode:   code,
		}

		if err := db.Create(&url).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				continue
			}
			return nil, err
		}
		return &url, nil
	}
}

func GetShortURL(db *gorm.DB, code string) (*ShortURL, error) {
	var url ShortURL
	if err := db.Where("short_code = ?", code).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}
