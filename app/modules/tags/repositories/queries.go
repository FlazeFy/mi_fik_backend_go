package repositories

import (
	"app/modules/tags/models"
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/converter"
	"app/packages/helpers/response"
	"app/packages/utils/pagination"
	"database/sql"
	"math"
	"net/http"
)

func GetAllTag(page, pageSize int, path string, view string) (response.Response, error) {
	var obj models.GetAllTag
	var arrobj []models.GetAllTag
	var res response.Response
	var tableName = "tags"
	var whereActive string
	var sqlStatement string

	selectTemplate := builders.GetTemplateSelect("tag_info")
	if view == "active" {
		whereActive = tableName + ".deleted_at IS NULL "
	} else {
		whereActive = tableName + ".deleted_at IS NOT NULL "
	}
	order := "tags.created_at DESC, tags.id DESC "

	var TagCategory sql.NullString

	con := database.CreateCon()

	offset := (page - 1) * pageSize

	if view == "active" {
		sqlStatement = "SELECT " + selectTemplate + " FROM " + tableName + " " +
			"LEFT JOIN dictionaries ON dictionaries.slug_name = tags.tag_category " +
			"WHERE " + whereActive +
			"ORDER BY " + order +
			"LIMIT ? OFFSET ?"
	} else {
		sqlStatement = "SELECT " + selectTemplate + " FROM " + tableName + " " +
			"LEFT JOIN dictionaries ON dictionaries.slug_name = tags.tag_category " +
			"WHERE " + whereActive +
			"ORDER BY " + order +
			"LIMIT ? OFFSET ?"
	}

	rows, err := con.Query(sqlStatement, pageSize, offset)
	defer rows.Close()

	if err != nil {
		return res, err
	}

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

	total, err := builders.GetTotalCount(con, tableName, whereActive)
	if err != nil {
		return res, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	pagination := pagination.BuildPaginationResponse(page, pageSize, total, totalPages, path)

	res.Status = http.StatusOK
	res.Message = "Tag Found"
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

	return res, nil
}
