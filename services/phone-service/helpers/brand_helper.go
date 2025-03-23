package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Brand struct {
	ID        uuid.UUID `json:"id"`
	BrandCode string    `json:"brandCode"`
	BrandName string    `json:"brandName"`
	IsDeleted bool      `json:"isDeleted"`
}

func FindBrandByID(id uuid.UUID) (Brand, error) {

	resp, err := http.Get("http://localhost:8080/brands/brand?id=" + id.String())
	if err != nil {
		return Brand{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Brand{}, fmt.Errorf("failed to fetch brand: HTTP %d %s", resp.StatusCode, resp.Status)
	}
	var brand Brand
	if err := json.NewDecoder(resp.Body).Decode(&brand); err != nil {
		return Brand{}, fmt.Errorf("failed to decode response: %w", err)
	}
	return brand, nil
}
func FindBrands() ([]Brand, error) {

	resp, err := http.Get("http://localhost:8080/brands/brands")
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch brands: HTTP %d %s", resp.StatusCode, resp.Status)
	}
	var brands []Brand
	if err := json.NewDecoder(resp.Body).Decode(&brands); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return brands, nil
}
