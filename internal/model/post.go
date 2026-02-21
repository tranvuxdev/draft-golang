package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID   uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Title    string    `gorm:"type:varchar(255);not null" json:"title"`
	Content  string    `gorm:"type:text" json:"content"`
	CreateAt time.Time `gorm:"autoCreateTime" json:"create_at"`

	// 1 user - many post
	User User `gorm:"foreignKey:UserID"`

	// many post - many tag
	Tags []Tag `gorm:"many2many:post_tag" json:"-"`
}
