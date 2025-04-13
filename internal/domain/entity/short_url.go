package entity

import (
	"time"
)

type ShortURL struct {
	ID         string    `bson:"_id,omitempty" json:"id"`
	Original   string    `bson:"original" json:"original"`
	Code       string    `bson:"code" json:"code"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
	Accessed   int64     `bson:"accessed" json:"accessed"`
	LastAccess time.Time `bson:"last_access" json:"last_access"`
	IsActive   bool      `bson:"is_active" json:"is_active"`
}

func NewShortURL(original, code string) *ShortURL {
	return &ShortURL{
		Original:   original,
		Code:       code,
		CreatedAt:  time.Now(),
		Accessed:   0,
		LastAccess: time.Time{},
		IsActive:   true,
	}
}
