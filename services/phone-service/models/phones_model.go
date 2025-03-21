package models

import (
	"time"

	"github.com/google/uuid"
)

type Phones struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BrandID   uuid.UUID `json:"brandId" gorm:"type:uuid;not null"`
	ModelID   uuid.UUID `json:"modelId" gorm:"type:uuid;not null"`
	StorageID uuid.UUID `json:"storageId" gorm:"type:uuid;not null"`
	PhoneCode string    `json:"phoneCode" gorm:"unique;index"`
	PhoneName string    `json:"phoneName" gorm:"index"`
	Price     float64   `json:"price" gorm:"default:0;index"`
	MinPrice  float64   `json:"minPrice" gorm:"default:0;index"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	IsDeleted bool      `json:"isDeleted" gorm:"default:false"`
}
