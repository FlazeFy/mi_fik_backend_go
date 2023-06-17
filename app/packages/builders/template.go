package builders

func GetTemplateSelect(name string) string {
	if name == "tag_info" {
		return "tags.slug_name, tag_name, dictionaries.dct_name as tag_category"
	}
	return ""
}
