package builders

import "app/packages/helpers/typography"

func GetTemplateSelect(name string, firstTable, secondTable *string) string {
	if name == "tag_info" {
		return "tags.slug_name, tag_name, dictionaries.dct_name as tag_category"
	} else if name == "dct_info" {
		return "slug_name, dct_name, dct_type, dct_desc, type_name"
	} else if name == "properties" {
		if firstTable == nil {
			return "created_at, created_by, updated_at, updated_by"
		} else {
			return *firstTable + ".created_at, " + *secondTable + ".username as created_by, " + *firstTable + ".updated_at, " + *secondTable + ".username as updated_by"
		}
	}
	return ""
}

func GetTemplateConcat(name, col string) string {
	if name == "value_group" {
		return "GROUP_CONCAT(" + col + " SEPARATOR ', ') as " + col
	}
	return ""
}

func GetTemplateGeneralSelect(name string, firstTable *string) string {
	ftable := typography.RemoveLastChar(*firstTable, "s")
	if name == "info" {
		return *firstTable + ".slug_name, " + ftable + "_name, " + ftable + "_desc"
	}
	return ""
}

func GetTemplateOrder(name, tableName, ext string) string {
	if name == "permanent_data" {
		return tableName + ".created_at DESC, " + tableName + "." + ext + " DESC "
	} else if name == "dynamic_data" {
		return tableName + ".updated_at DESC, " + tableName + ".created_at DESC, " + tableName + "." + ext + " DESC "
	} else if name == "most_used_normal" {
		return " COUNT(1) DESC"
	}
	return ""
}

func GetTemplateJoin(typeJoin, firstTable, firstCol, secondTable, secondCol string, nullable bool) string {
	var join = ""
	if nullable {
		join = "LEFT JOIN "
	} else {
		join = "JOIN "
	}
	if typeJoin == "same_col" {
		return join + secondTable + " USING(" + firstCol + ") "
	} else if typeJoin == "total" {
		return join + secondTable + " ON " + secondTable + "." + secondCol + " = " + firstTable + "." + firstCol + " "
	}
	return ""
}

func GetTemplateGroup(is_multi_where bool, col string) string {
	var ext = " WHERE "
	if is_multi_where {
		ext = " AND "
	}

	return ext + col + " IS NOT NULL AND trim(" + col + ") != '' GROUP BY " + col + " "
}

func GetTemplateLogic(name string) string {
	if name == "active" {
		return ".deleted_at IS NULL "
	} else if name == "trash" {
		return ".deleted_at IS NOT NULL "
	}
	return ""
}
