package model

import (
	"time"

	"github.com/google/uuid"
)

type Card struct {
	InternalID  int64      `json:"internal_id" db:"internal_id" gorm:"column:internal_id;type:bigint;not null;primaryKey;autoIncrement"`
	PublicID    uuid.UUID  `json:"public_id" db:"public_id" gorm:"column:public_id;type:uuid;not null"`
	ListID      int64      `json:"list_internal_id" db:"list_internal_id" gorm:"column:list_internal_id;type:bigint;not null"`
	Title       string     `json:"title" db:"title" gorm:"column:title;type:varchar(255);not null"`
	Description string     `json:"description" db:"description" gorm:"column:description;type:text"`
	DueDate     *time.Time `json:"due_date,omitempty" db:"due_date" gorm:"column:due_date;type:timestamp"`
	Position    int        `json:"position" db:"position" gorm:"column:position;type:integer;not null"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at" gorm:"column:created_at;not null"`
}
