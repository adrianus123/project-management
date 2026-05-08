package model

import (
	"time"

	"github.com/google/uuid"
)

type Board struct {
	InternalID    int64      `json:"internal_id" db:"internal_id" gorm:"column:internal_id;primaryKey;not null;autoIncrement"`
	PublicID      uuid.UUID  `json:"public_id" db:"public_id" gorm:"column:public_id;type:uuid;not null"`
	Title         string     `json:"title" db:"title" gorm:"column:title;type:varchar(255);not null"`
	Description   string     `json:"description" db:"description" gorm:"column:description;type:text;not null"`
	OwnerID       int64      `json:"owner_internal_id" db:"owner_internal_id" gorm:"column:owner_internal_id;type:bigint;not null"`
	OwnerPublicID uuid.UUID  `json:"owner_public_id" db:"owner_public_id" gorm:"column:owner_public_id;type:uuid;not null"`
	DueDate       *time.Time `json:"due_date,omitempty" db:"due_date" gorm:"column:due_date"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at" gorm:"column:created_at;not null"`
}
