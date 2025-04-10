package core

import (
	"time"

	"github.com/google/uuid"
)

type SQLModel struct {
	ID        string    `json:"-" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime;column:updated_at"`
}

// BeforeCreate - Automatically assigns a UUID to FakeId before saving
func (sqlModel *SQLModel) BeforeCreate(objectId int) {
	// Generate uuid and store it as FakeId
	sqlModel.ID = uuid.New().String()
}
