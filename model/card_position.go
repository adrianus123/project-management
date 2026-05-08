package model

import (
	"github.com/adrianus123/project-management/model/types"
	"github.com/google/uuid"
)

type CardPosition struct {
	InternalID int64           `json:"internal_id" db:"internal_id" gorm:"column:internal_id;type:bigint;not null;primaryKey;autoIncrement"`
	PublicID   uuid.UUID       `json:"public_id" db:"public_id" gorm:"column:public_id;type:uuid;not null"`
	ListID     int64           `json:"list_internal_id" db:"list_internal_id" gorm:"column:list_internal_id;type:bigint;not null"`
	CardOrder  types.UUIDArray `json:"card_order" gorm:"column:card_order;type:uuid[]"`
}
