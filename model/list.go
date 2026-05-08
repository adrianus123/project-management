package model

import (
	"time"

	"github.com/google/uuid"
)

type List struct {
	InternalID      int64     `json:"internal_id" db:"internal_id" gorm:"column:internal_id;primaryKey;not null;autoIncrement"`
	PublicID        uuid.UUID `json:"public_id" db:"public_id" gorm:"column:public_id;type:uuid;not null"`
	BoardPublicID   uuid.UUID `json:"board_public_id" db:"board_public_id" gorm:"column:board_public_id;type:uuid;not null"`
	Title           string    `json:"title" db:"title" gorm:"column:title;type:varchar(255);not null"`
	CreatedAt       time.Time `json:"created_at" db:"created_at" gorm:"column:created_at;not null"`
	BoardInternalID int64     `json:"-" db:"board_internal_id" gorm:"column:board_internal_id;type:bigint;not null"`
}
