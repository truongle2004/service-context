package core

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SQLModel struct {
	ID        string    `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime;column:updated_at"`
}

func (sqlModel *SQLModel) BeforeCreate(tx *gorm.DB) (err error) {
	if sqlModel.ID == "" {
		sqlModel.ID = uuid.New().String()
	}
	return
}
