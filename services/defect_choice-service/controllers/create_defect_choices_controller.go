package controllers

import (
	"defect_choice-service/config"
	"defect_choice-service/models"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckDefectExists(defectId string) (bool, error) {
	req, err := http.NewRequest("GET", "http://localhost:8084/defects/defect?id="+defectId, nil)
	if err != nil {
		return false, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return false, err
	}

	if len(result) > 0 {
		return true, nil
	}

	return false, nil
}

func ValidateDefectChoice(c *gin.Context, defectChoice models.DefectChoices) bool {
	var existingDefect models.DefectChoices
	if err := config.DB.Where("choice_code = ?", defectChoice.ChoiceCode).First(&existingDefect).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "choiceCode already exists"})
		return false
	} else if err.Error() != "record not found" {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "database error: " + err.Error()})
		return false
	}
	return true
}
func ValidateDefect(c *gin.Context, defectId string) bool {
	exists, err := CheckDefectExists(defectId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "failed to check defect: " + err.Error()})
		return false
	}
	if !exists {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": "defect does not exist"})
		return false
	}
	return true
}
func CreateDefectChoice(c *gin.Context) {
	var defectChoices []models.DefectChoices
	if err := c.ShouldBindJSON(&defectChoices); err != nil {
		var singleDefectChoice models.DefectChoices
		if err := c.ShouldBindJSON(&singleDefectChoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request data"})
			return
		}
		defectChoices = append(defectChoices, singleDefectChoice)
	}
	var newDefectChoices []models.DefectChoices
	for _, defectChoice := range defectChoices {
		if !ValidateDefectChoice(c, defectChoice) {
			return
		}
		if !ValidateDefect(c, defectChoice.DefectID.String()) {
			return
		}
		newDefectChoices = append(newDefectChoices, defectChoice)
	}
	if len(newDefectChoices) == 0 {
		c.JSON(http.StatusConflict, gin.H{"success": false, "error": "All defect choice codes already exist"})
		return
	}
	if err := config.DB.Create(&newDefectChoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create defect choices"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Defect choices created successfully"})
}
