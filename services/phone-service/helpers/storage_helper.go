package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Storage struct {
	ID uuid.UUID `json:"id"`
}

func FindStorageByID(id uuid.UUID) (Storage, error) {
	resp, err := http.Get("http://localhost:8080/storages/storage?id=" + id.String())
	if err != nil {
		return Storage{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Storage{}, fmt.Errorf("failed to fetch storage: HTTP %d %s", resp.StatusCode, resp.Status)
	}
	var storage Storage
	if err := json.NewDecoder(resp.Body).Decode(&storage); err != nil {
		return Storage{}, fmt.Errorf("failed to decode response: %w", err)
	}
	return storage, nil
}
