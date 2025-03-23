package models

type ViewPhoneWithDeductions struct {
	BrandID      string  `json:"brandId"`
	BrandName    string  `json:"brandName"`
	ModelID      string  `json:"modelId"`
	ModelName    string  `json:"modelName"`
	StorageID    string  `json:"storageId"`
	StorageValue string  `json:"storageValue"`
	PhoneID      string  `json:"phoneId"`
	PhoneName    string  `json:"phoneName"`
	DefectID     string  `json:"defectId"`
	Index        int     `json:"index"`
	DefectName   string  `json:"defectName"`
	ConfigID     string  `json:"configId"`
	ChoiceID     string  `json:"choiceId"`
	ChoiceName   string  `json:"choiceName"`
	Price        float64 `json:"price"`
	MinPrice     float64 `json:"minPrice"`
	Deduction    float64 `json:"deduction"`
}

func (b *ViewPhoneWithDeductions) TableName() string {
	return "view_phone_with_deductions"
}
