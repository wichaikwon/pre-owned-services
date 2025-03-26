package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Defect struct {
	ID         string `json:"id"`
	DefectCode string `json:"defectCode"`
	DefectName string `json:"defectName"`
	IsDeleted  bool   `json:"isDeleted"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

func FetchDefects() ([]Defect, error) {
	resp, err := http.Get("http://localhost:8084/defects/defects")
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch defects: HTTP %d %s", resp.StatusCode, resp.Status)
	}

	var defects []Defect
	if err := json.NewDecoder(resp.Body).Decode(&defects); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return defects, nil
}
