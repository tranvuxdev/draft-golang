package model

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	PostID   uuid.UUID `gorm:"type:uuid" json:"post_id"`
	UserID   uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Body     string    `gorm:"type:text" json:"body"`
	CreateAt time.Time `gorm:"autoCreateTime"`

	// 1 post - many comment
	Post Post `gorm:"foreignKey:PostID"`
	// 1 user - many comment
	User User `gorm:"foreignKey: UserID"`
}
