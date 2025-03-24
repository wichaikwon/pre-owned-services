package models

import "github.com/google/uuid"

type ViewPhones struct {
	PhoneID      uuid.UUID `json:"phoneId"`
	BrandID      uuid.UUID `json:"brandId"`
	ModelID      uuid.UUID `json:"modelId"`
	StorageID    uuid.UUID `json:"storageId"`
	BrandCode    string    `json:"brandCode"`
	BrandName    string    `json:"brandName"`
	ModelCode    string    `json:"modelCode"`
	ModelName    string    `json:"modelName"`
	PhoneCode    string    `json:"phoneCode"`
	PhoneName    string    `json:"phoneName"`
	StorageCode  string    `json:"storageCode"`
	StorageValue string    `json:"storageValue"`
	Price        float64   `json:"price"`
	MinPrice     float64   `json:"minPrice"`
	IsDeleted    bool      `json:"isDeleted"`
}

func (b *ViewPhones) TableName() string {
	return "view_phones"
}
