package model

import (
	"github.com/google/uuid"
)

type Label struct {
	InternalID int64     `json:"internal_id" db:"internal_id" gorm:"column:internal_id;type:bigint;not null;primaryKey;autoIncrement"`
	PublicID   uuid.UUID `json:"public_id" db:"public_id" gorm:"column:public_id;type:uuid;not null"`
	Name       string    `json:"name" db:"name" gorm:"column:name;type:varchar(100);not null"`
	Color      string    `json:"color" db:"color" gorm:"column:color;type:varchar(7);not null"`
}
