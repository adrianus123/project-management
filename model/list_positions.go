package model

import (
	"github.com/adrianus123/project-management/model/types"
	"github.com/google/uuid"
)

type ListPosition struct {
	InternalID int64           `json:"internal_id" db:"internal_id" gorm:"column:internal_id;type:bigint;not null;primaryKey;autoIncrement"`
	PublicID   uuid.UUID       `json:"public_id" db:"public_id" gorm:"column:public_id;type:uuid;not null"`
	BoardID    int64           `json:"board_internal_id" db:"board_internal_id" gorm:"column:board_internal_id;type:bigint;not null"`
	ListOrder  types.UUIDArray `json:"list_order"`
}
