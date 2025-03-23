package models

import "github.com/google/uuid"

type ViewBrands struct {
	BrandID   uuid.UUID `json:"brandId"`
	BrandName string    `json:"brandName"`
	IsDeleted bool      `json:"isDeleted"`
}

func (b *ViewBrands) TableName() string {
	return "view_brands"
}
