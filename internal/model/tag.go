package model

import "github.com/google/uuid"

type Tag struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name string    `gorm:"type:varchar(255);not null" json:"name"`

	// many post - many tag
	Posts []Post `gorm:"many2many:post_tag"`
}
