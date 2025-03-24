package models

import (
	"time"

	"github.com/google/uuid"
)

type DefectChoices struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	DefectID   uuid.UUID `json:"defectId" gorm:"type:uuid;not null"`
	Index      int       `json:"index" gorm:"index"`
	ChoiceCode string    `json:"choiceCode" gorm:"not null;index"`
	ChoiceName string    `json:"choiceName" gorm:"index"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	IsDeleted  bool      `json:"isDeleted" gorm:"default:false"`
}
