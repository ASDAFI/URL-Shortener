package links

import (
	"crypto/md5"
	"fmt"
	"gorm.io/gorm"
	"regexp"
	"strconv"
	"time"
)

type Link struct {
	gorm.Model
	OriginLink    string    `gorm:"column:origin_link;unique_index;not null;"`
	ShortenedLink *string   `gorm:"column:shortened_link;"`
	ExpiresAt     time.Time `gorm:"column:expires_at;"`
}

func (Link) TableName() string {
	return "links"
}

func md5Encode(input uint) *string {

	hashInput := strconv.Itoa(int(input))
	// Create a new hash & write input string
	hash := md5.New()
	_, _ = hash.Write([]byte(hashInput))

	// Get the resulting encoded byte slice
	md5 := hash.Sum(nil)

	// Convert the encoded byte slice to a string
	result := fmt.Sprintf("%x", md5)
	return &result
}

func isValidURL(url string) bool {
	// Regular expression pattern for validating URL format
	pattern := `^(http(s)?:\/\/)?([^\s]+\.){1,2}[^\s]+(\/[^\s]*)?$`

	match, _ := regexp.MatchString(pattern, url)
	return match
}
