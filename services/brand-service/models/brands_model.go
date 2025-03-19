package models

import (
	"time"

	"github.com/google/uuid"
)

type Brands struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	BrandCode string    `json:"brandCode" gorm:"unique;index"`
	BrandName string    `json:"brandName" gorm:"index"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	IsDeleted bool      `json:"isDeleted" gorm:"default:false"`
}
