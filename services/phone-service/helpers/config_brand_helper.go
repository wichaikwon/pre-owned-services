package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type ConfigBrand struct {
	ID         uuid.UUID `json:"id"`
	BrandID    uuid.UUID `json:"brandId"`
	DefectID   uuid.UUID `json:"defectId"`
	BrandCode  string    `json:"brandCode"`
	DefectCode string    `json:"defectCode"`
	BrandName  string    `json:"brandName"`
	DefectName string    `json:"defectName"`
	IsDeleted  bool      `json:"isDeleted"`
}

func FetchConfigBrands() ([]ConfigBrand, error) {
	resp, err := http.Get("https://pre-owned-brand-service-production.up.railway.app/brands/config-brands")
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch config brands: HTTP %d %s", resp.StatusCode, resp.Status)
	}
	var allConfigBrands []ConfigBrand
	if err := json.NewDecoder(resp.Body).Decode(&allConfigBrands); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	var configBrands []ConfigBrand
	for _, brand := range allConfigBrands {
		if !brand.IsDeleted {
			configBrands = append(configBrands, brand)
		}
	}
	return configBrands, nil
}
