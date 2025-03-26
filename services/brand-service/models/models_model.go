package models

import (
	"time"

	"github.com/google/uuid"
)

type Models struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BrandID   uuid.UUID `json:"brandId" gorm:"type:uuid;not null"`
	ModelCode string    `json:"modelCode" gorm:"unique;index"`
	ModelName string    `json:"modelName" gorm:"index"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	IsDeleted bool      `json:"isDeleted" gorm:"default:false"`
}
