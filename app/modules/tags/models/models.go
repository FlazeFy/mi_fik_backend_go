package models

type (
	GetAllTag struct {
		SlugName    string `json:"slug_name"`
		TagName     string `json:"tag_name"`
		TagCategory string `json:"tag_category"`
	}
	GetAllTagByCategory struct {
		SlugName string `json:"slug_name"`
		TagName  string `json:"tag_name"`
		TagDesc  string `json:"tag_desc"`
	}
)
