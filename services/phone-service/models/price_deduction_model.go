package models

import (
	"time"

	"github.com/google/uuid"
)

type PriceDeductions struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PhoneID        uuid.UUID `json:"phoneId" gorm:"type:uuid;not null"`
	ConfigBrandID  uuid.UUID `json:"configBrandId" gorm:"type:uuid;not null"`
	DefectChoiceID uuid.UUID `json:"defectChoiceId" gorm:"type:uuid;not null"`
	Deduction      float64   `json:"deduction" gorm:"default:0"`
	CreatedAt      time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	IsDeleted      bool      `json:"isDeleted" gorm:"default:false"`
}
