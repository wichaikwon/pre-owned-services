package models

import "github.com/google/uuid"

type ViewModels struct {
	ModelID   uuid.UUID `json:"modelId"`
	BrandID   uuid.UUID `json:"brandId"`
	BrandName string    `json:"brandName"`
	ModelName string    `json:"modelName"`
	IsDeleted bool      `json:"isDeleted"`
}

func (b *ViewModels) TableName() string {
	return "view_models"
}
