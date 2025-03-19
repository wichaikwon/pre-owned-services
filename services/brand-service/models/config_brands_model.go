package models

import (
	"time"

	"github.com/google/uuid"
)

type ConfigBrands struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default;primaryKey;default:uuid_generate_v4()"`
	BrandID    uuid.UUID `json:"brandId" gorm:"type:uuid;not null"`
	DefectID   uuid.UUID `json:"defectId" gorm:"type:uuid;not null"`
	BrandCode  string    `json:"brandCode" gorm:"not null"`
	DefectCode string    `json:"defectCode" gorm:"not null"`
	BrandName  string    `json:"brandName" gorm:"not null"`
	DefectName string    `json:"defectName" gorm:"not null"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	IsDeleted  bool      `json:"isDeleted" gorm:"default:true"`
}
