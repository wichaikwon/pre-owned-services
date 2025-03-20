package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Phone struct {
	ID uuid.UUID `json:"id"`
}

func FetchPhones(ID uuid.UUID) ([]Phone, error) {
	url := fmt.Sprintf("http://localhost:8080/phones/phone?id=%s", ID.String())
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)

	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch phones : HTTP %d %s", resp.StatusCode, resp.Status)
	}
	var phones []Phone
	if err := json.NewDecoder(resp.Body).Decode(&phones); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return phones, nil
}
