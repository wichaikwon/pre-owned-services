package models

import (
	"time"

	"github.com/google/uuid"
)

type Defects struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	DefectCode string    `json:"defectCode" gorm:"unique;idex"`
	DefectName string    `json:"defectName" gorm:"index"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	IsDeleted  bool      `json:"isDeleted" gorm:"default:false"`
}
