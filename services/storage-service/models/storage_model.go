package models

import "time"

type Storages struct {
	ID           string    `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	StorageCode  string    `json:"storageCode" gorm:"unique"`
	StorageValue string    `json:"storageValue"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	IsDeleted    bool      `json:"isDeleted" gorm:"default:false"`
}
