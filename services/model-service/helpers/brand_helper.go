package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Brand struct {
	ID        string `json:"id"`
	BrandCode string `json:"brandCode"`
	BrandName string `json:"brandName"`
	IsDeleted bool   `json:"isDeleted"`
}

func CheckBrandExists(brandId uuid.UUID) (bool, error) {
	url := fmt.Sprintf("http://localhost:8080/brands/brand?id=%s", brandId.String())
	resp, err := http.Get(url)
	if err != nil {
		return false, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch brands: HTTP %d %s", resp.StatusCode, resp.Status)
	}
	var brand Brand
	if err := json.NewDecoder(resp.Body).Decode(&brand); err != nil {
		return false, fmt.Errorf("failed to decode response: %w", err)
	}
	if brand.ID == brandId.String() && !brand.IsDeleted {
		return true, nil
	}

	return false, nil
}
