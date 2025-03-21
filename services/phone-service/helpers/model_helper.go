package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Model struct {
	ID      uuid.UUID `json:"id"`
	BrandID uuid.UUID `json:"brand_id"`
}

func FindModelByID(id uuid.UUID) (Model, error) {
	resp, err := http.Get("http://localhost:8080/models/model?id=" + id.String())
	if err != nil {
		return Model{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Model{}, fmt.Errorf("failed to fetch model: HTTP %d %s", resp.StatusCode, resp.Status)
	}
	var model Model
	if err := json.NewDecoder(resp.Body).Decode(&model); err != nil {
		return Model{}, fmt.Errorf("failed to decode response: %w", err)
	}
	return model, nil
}
