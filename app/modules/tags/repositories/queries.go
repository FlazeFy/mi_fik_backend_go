package repositories

import (
	"app/modules/tags/models"
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/converter"
	"app/packages/helpers/generator"
	"app/packages/helpers/response"
	"app/packages/utils/pagination"
	"database/sql"
	"math"
	"net/http"
)

func GetAllTag(page, pageSize int, path string, view string) (response.Response, error) {
	// Declaration
	var obj models.GetAllTag
	var arrobj []models.GetAllTag
	var res response.Response
	var baseTable = "tags"
	var secondTable = "dictionaries"
	var sqlStatement string

	// Nullable column
	var TagCategory sql.NullString

	// Query builder
	selectTemplate := builders.GetTemplateSelect("tag_info", nil, nil)
	firstLogicWhere := builders.GetTemplateLogic(view)
	whereActive := baseTable + firstLogicWhere
	join1 := builders.GetTemplateJoin("total", baseTable, "tag_category", secondTable, "slug_name", true)
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "tag_name")

	sqlStatement = "SELECT " + selectTemplate + " " +
		"FROM " + baseTable + " " +
		join1 +
		"WHERE " + whereActive +
		"ORDER BY " + order +
		"LIMIT ? OFFSET ?"

	// Exec
	con := database.CreateCon()
	offset := (page - 1) * pageSize
	rows, err := con.Query(sqlStatement, pageSize, offset)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.SlugName,
			&obj.TagName,
			&TagCategory)

		if err != nil {
			return res, err
		}

		obj.TagCategory = converter.CheckNullString(TagCategory)

		arrobj = append(arrobj, obj)
	}

	// Page
	total, err := builders.GetTotalCount(con, baseTable, whereActive)
	if err != nil {
		return res, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	pagination := pagination.BuildPaginationResponse(page, pageSize, total, totalPages, path)

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateNormalMsg("Tag", total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = map[string]interface{}{
			"current_page":   page,
			"data":           arrobj,
			"first_page_url": pagination.FirstPageURL,
			"from":           pagination.From,
			"last_page":      pagination.LastPage,
			"last_page_url":  pagination.LastPageURL,
			"links":          pagination.Links,
			"next_page_url":  pagination.NextPageURL,
			"path":           pagination.Path,
			"per_page":       pageSize,
			"prev_page_url":  pagination.PrevPageURL,
			"to":             pagination.To,
			"total":          total,
		}
	}

	return res, nil
}

func GetAllTagByCategory(page, pageSize int, path string, view string, category string) (response.Response, error) {
	// Declaration
	var obj models.GetAllTagByCategory
	var arrobj []models.GetAllTagByCategory
	var res response.Response
	var baseTable = "tags"
	var secondTable = "dictionaries"
	var sqlStatement string

	// Nullable column
	var TagDesc sql.NullString

	// Query builder
	selectTemplate := builders.GetTemplateGeneralSelect("info", &baseTable)
	firstLogicWhere := builders.GetTemplateLogic(view)
	whereActive := baseTable + firstLogicWhere
	whereByCategory := secondTable + ".slug_name = '" + category + "' "
	join1 := builders.GetTemplateJoin("total", baseTable, "tag_category", secondTable, "slug_name", true)
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "tag_name")

	sqlStatement = "SELECT " + selectTemplate + " " +
		"FROM " + baseTable + " " +
		join1 +
		"WHERE " + whereActive +
		"AND " + whereByCategory +
		"ORDER BY " + order +
		"LIMIT ? OFFSET ?"

	// Exec
	con := database.CreateCon()
	offset := (page - 1) * pageSize
	rows, err := con.Query(sqlStatement, pageSize, offset)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.SlugName,
			&obj.TagName,
			&TagDesc)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	// Page
	total, err := builders.GetTotalCount(con, baseTable+" "+join1, whereActive+" AND "+whereByCategory)
	if err != nil {
		return res, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	pagination := pagination.BuildPaginationResponse(page, pageSize, total, totalPages, path)

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateNormalMsg("Tag", total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = map[string]interface{}{
			"current_page":   page,
			"data":           arrobj,
			"first_page_url": pagination.FirstPageURL,
			"from":           pagination.From,
			"last_page":      pagination.LastPage,
			"last_page_url":  pagination.LastPageURL,
			"links":          pagination.Links,
			"next_page_url":  pagination.NextPageURL,
			"path":           pagination.Path,
			"per_page":       pageSize,
			"prev_page_url":  pagination.PrevPageURL,
			"to":             pagination.To,
			"total":          total,
		}
	}

	return res, nil
}
