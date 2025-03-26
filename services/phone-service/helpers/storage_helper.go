package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Storages struct {
	ID           uuid.UUID `json:"id"`
	StorageCode  string    `json:"storageCode"`
	StorageValue string    `json:"storageValue"`
	IsDeleted    bool      `json:"isDeleted"`
}

func FindStorageByID(id uuid.UUID) (Storages, error) {
	resp, err := http.Get("https://pre-owned-brand-service-production.up.railway.app/storages/storage?id=" + id.String())
	if err != nil {
		return Storages{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Storages{}, fmt.Errorf("failed to fetch storage: HTTP %d %s", resp.StatusCode, resp.Status)
	}
	var storage Storages
	if err := json.NewDecoder(resp.Body).Decode(&storage); err != nil {
		return Storages{}, fmt.Errorf("failed to decode response: %w", err)
	}
	return storage, nil
}
