package models

type (
	GetAllDictionary struct {
		SlugName  string `json:"slug_name"`
		DctName   string `json:"dct_name"`
		DctDesc   string `json:"dct_desc"`
		DctType   string `json:"dct_type"`
		TypeName  string `json:"type_name"`
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
		UpdatedAt string `json:"updated_at"`
		UpdatedBy string `json:"updated_by"`
	}
)
