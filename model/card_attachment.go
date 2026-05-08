package model

import (
	"time"

	"github.com/google/uuid"
)

type CardAttachment struct {
	InternalID int64     `json:"internal_id" db:"internal_id" gorm:"column:internal_id;type:bigint;not null;primaryKey;autoIncrement"`
	PublicID   uuid.UUID `json:"public_id" db:"public_id" gorm:"column:public_id;type:uuid;not null"`
	CardID     int64     `json:"card_internal_id" db:"card_internal_id" gorm:"column:card_internal_id;type:bigint;not null"`
	UserID     int64     `json:"user_internal_id" db:"user_internal_id" gorm:"column:user_internal_id;type:bigint;not null"`
	File       string    `json:"file" db:"file" gorm:"column:file;type:text;not null"`
	CreatedAt  time.Time `json:"created_at" db:"created_at" gorm:"column:created_at;type:timestamp;not null"`
}
