package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type DefectChoice struct {
	ID         uuid.UUID `json:"id"`
	DefectID   uuid.UUID `json:"defectId"`
	ChoiceCode string    `json:"choiceCode"`
	ChoiceName string    `json:"choiceName"`
	CreatedAt  string    `json:"createdAt"` // เพิ่มฟิลด์นี้ถ้าจำเป็น
	UpdatedAt  string    `json:"updatedAt"` // เพิ่มฟิลด์นี้ถ้าจำเป็น
	IsDeleted  bool      `json:"isDeleted"` // เพิ่มฟิลด์นี้ถ้าจำเป็น
}

func FetchDefectChoices(defectID uuid.UUID) ([]DefectChoice, error) {
	url := fmt.Sprintf("http://localhost:8080/defect-choices/defect-choice/defects?id=%s", defectID.String())

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch defect choices: HTTP %d %s", resp.StatusCode, resp.Status)
	}

	// แปลง JSON เป็น array ของ DefectChoice
	var defectChoices []DefectChoice
	if err := json.NewDecoder(resp.Body).Decode(&defectChoices); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return defectChoices, nil
}
