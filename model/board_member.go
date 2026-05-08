package model

import (
	"time"
)

type BoardMember struct {
	BoardID  int64     `json:"board_internal_id" db:"board_internal_id" gorm:"column:board_internal_id;type:bigint;not null;primaryKey"`
	UserID   int64     `json:"user_internal_id" db:"user_internal_id" gorm:"column:user_internal_id;type:bigint;not null;primaryKey"`
	JoinedAt time.Time `json:"joined_at" db:"joined_at" gorm:"column:joined_at;type:timestamp;not null"`
}
