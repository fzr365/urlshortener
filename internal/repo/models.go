// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repo

import (
	"time"
)

type Url struct {
	ID          int64     `json:"id"`
	OriginalUrl string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`
	IsCustom    bool      `json:"is_custom"`
	ExpiredAt   time.Time `json:"expired_at"`
	CreatedAt   time.Time `json:"created_at"`
}
