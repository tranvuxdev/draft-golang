package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name     string    `gorm:"type:varchar(255);not null" json:"name"`
	Email    string    `gorm:"type:varchar(255);not null" json:"email"`
	CreateAt time.Time `gorm:"autoCreateTime" json:"create_at"`
}
