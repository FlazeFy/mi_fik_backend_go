package repositories

import (
	"app/modules/systems/models"
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

func GetAllDictionary(page, pageSize int, path string, view string) (response.Response, error) {
	// Declaration
	var obj models.GetAllDictionary
	var arrobj []models.GetAllDictionary
	var res response.Response
	var baseTable = "dictionaries"
	var secondTable = "dictionaries_types"
	var thirdTable = "admins"
	var sqlStatement string

	// Nullable column
	var DctDesc sql.NullString
	var UpdatedAt sql.NullString
	var UpdatedBy sql.NullString

	// Query builder
	selectTemplate := builders.GetTemplateSelect("dct_info", nil, nil)
	extSelect := builders.GetTemplateSelect("properties", &baseTable, &thirdTable)
	firstLogicWhere := builders.GetTemplateLogic(view)
	whereActive := baseTable + firstLogicWhere
	order := builders.GetTemplateOrder("dynamic_data", baseTable, "dct_name")
	join1 := builders.GetTemplateJoin("total", baseTable, "dct_type", secondTable, "app_code", false)
	join2 := builders.GetTemplateJoin("total", baseTable, "created_by", thirdTable, "id", true)

	sqlStatement = "SELECT " + selectTemplate + ", " + extSelect + " " +
		"FROM " + baseTable + " " +
		join1 + join2 +
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
			&obj.DctName,
			&obj.DctType,
			&DctDesc,
			&obj.TypeName,
			&obj.CreatedAt,
			&obj.CreatedBy,
			&UpdatedAt,
			&UpdatedBy,
		)

		if err != nil {
			return res, err
		}

		obj.DctDesc = converter.CheckNullString(DctDesc)
		obj.UpdatedAt = converter.CheckNullString(UpdatedAt)
		obj.UpdatedBy = converter.CheckNullString(UpdatedBy)

		arrobj = append(arrobj, obj)
	}

	// Page
	total, err := builders.GetTotalCount(con, baseTable, &whereActive)
	if err != nil {
		return res, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	pagination := pagination.BuildPaginationResponse(page, pageSize, total, totalPages, path)

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("dictionary", total)
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

func GetAllRecentColor(view string) (response.Response, error) {
	// Declaration
	var obj models.GetAllColor
	var arrobj []models.GetAllColor
	var res response.Response
	var baseTable = "dictionaries"
	var sqlStatement string

	// Nullable column
	var DctColor sql.NullString

	// Query Builder
	firstConcat := builders.GetTemplateConcat("value_group", "dct_name")
	selectTemplate := firstConcat + ", dct_color "
	firstLogicWhere := builders.GetTemplateLogic(view)
	whereActive := baseTable + firstLogicWhere
	order := builders.GetTemplateOrder("most_used_normal", baseTable, "dct_name")
	group := builders.GetTemplateGroup(true, "dct_color")

	sqlStatement = "SELECT " + selectTemplate +
		"FROM " + baseTable + " " +
		"WHERE " + whereActive + " " +
		group +
		"ORDER BY " + order

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.DctName,
			&DctColor,
		)

		if err != nil {
			return res, err
		}

		obj.DctColor = converter.CheckNullString(DctColor)

		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("dictionary color", len(arrobj))
	if len(arrobj) == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}
