package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	InternalID int64          `json:"internal_id" db:"internal_id" gorm:"column:internal_id;primaryKey;not null;autoIncrement"`
	PublicID   uuid.UUID      `json:"public_id" db:"public_id" gorm:"column:public_id;type:uuid;not null"`
	Name       string         `json:"name" db:"name" gorm:"column:name;type:varchar(255);not null"`
	Email      string         `json:"email" db:"email" gorm:"column:email;type:varchar(255);not null;uniqueIndex"`
	Password   string         `json:"password" db:"password" gorm:"column:password;type:varchar(255);not null"`
	Role       string         `json:"role" db:"role" gorm:"column:role;type:varchar(100);not null"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at" gorm:"column:created_at;not null"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated_at" gorm:"column:updated_at;not null"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"column:deleted_at;index"`
}

type UserResponse struct {
	PublicID  uuid.UUID `json:"public_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
