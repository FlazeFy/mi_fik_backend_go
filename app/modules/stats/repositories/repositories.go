package repositories

import (
	"app/modules/stats/models"
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/converter"
	"app/packages/helpers/generator"
	"app/packages/helpers/response"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

func GetMostAppearError(page, pageSize int, path string, ord string, limit string) (response.Response, error) {
	// Declaration
	var obj models.GetMostAppear
	var arrobj []models.GetMostAppear
	var res response.Response
	var baseTable = "errors"
	var mainCol = "message"
	var sqlStatement string

	// Converted column
	var totalStr string

	// Query builder
	selectTemplate := builders.GetTemplateStats(mainCol, baseTable, "most_appear", ord, nil)

	sqlStatement = selectTemplate + " LIMIT " + limit

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
			&obj.Context,
			&totalStr)

		if err != nil {
			return res, err
		}

		// Converted
		totalInt, err := strconv.Atoi(totalStr)
		if err != nil {
			return res, err
		}

		obj.Total = totalInt
		arrobj = append(arrobj, obj)
	}

	// Total
	total, err := builders.GetTotalCount(con, baseTable, nil)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}

func GetMostCreatedTagByCategory(path string, ord string, limit string) (response.Response, error) {
	// Declaration
	var obj models.GetMostAppear
	var arrobj []models.GetMostAppear
	var res response.Response
	var baseTable = "tags"
	var secondTable = "dictionaries"
	var joinArgs = "LEFT JOIN " + secondTable + " ON " + secondTable + ".slug_name = " + baseTable + ".tag_category"
	var mainCol = secondTable + ".dct_name"
	var sqlStatement string

	// Converted column
	var totalStr string

	// Nullable column
	var Context sql.NullString

	// Query builder
	selectTemplate := builders.GetTemplateStats(mainCol, baseTable, "most_appear", ord, &joinArgs)

	sqlStatement = selectTemplate + " LIMIT " + limit

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
			&Context,
			&totalStr)

		if err != nil {
			return res, err
		}

		// Converted
		totalInt, err := strconv.Atoi(totalStr)
		if err != nil {
			return res, err
		}

		obj.Context = converter.CheckNullString(Context)
		obj.Total = totalInt
		arrobj = append(arrobj, obj)
	}

	// Total
	total, err := builders.GetTotalCount(con, baseTable, nil)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}

func GetMostValidUntilUser(path string) (response.Response, error) {
	// Declaration
	var obj models.GetMostAppear
	var arrobj []models.GetMostAppear
	var res response.Response
	var baseTable = "users"

	// Converted column
	var totalStr string

	// Nullable column
	var Context sql.NullString

	// Query builder
	sqlStatement := builders.GetTemplateStats("valid_until", baseTable, "most_appear", "DESC", nil)

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
			&Context,
			&totalStr)

		if err != nil {
			return res, err
		}

		// Converted
		totalInt, err := strconv.Atoi(totalStr)
		if err != nil {
			return res, err
		}

		obj.Context = converter.CheckNullString(Context)
		obj.Total = totalInt
		arrobj = append(arrobj, obj)
	}

	// Total
	total, err := builders.GetTotalCount(con, baseTable, nil)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}

func GetMostActiveUser(page, pageSize int, path string, limit string) (response.Response, error) {
	// Declaration
	var obj models.GetMostAppear
	var arrobj []models.GetMostAppear
	var res response.Response
	var baseTable = "users_tokens"
	var secondTable = "users"
	var where = baseTable + ".context_type = 'user'"
	var joinArgs = "LEFT JOIN " + secondTable + " ON " + secondTable + ".id = " + baseTable + ".context_id" + " WHERE " + where
	var mainCol = secondTable + ".username"
	var sqlStatement string

	// Converted column
	var totalStr string

	// Query builder
	selectTemplate := builders.GetTemplateStats(mainCol, baseTable, "most_appear", "DESC", &joinArgs)

	sqlStatement = selectTemplate + " LIMIT " + limit

	fmt.Println(sqlStatement)

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
			&obj.Context,
			&totalStr)

		if err != nil {
			return res, err
		}

		// Converted
		totalInt, err := strconv.Atoi(totalStr)
		if err != nil {
			return res, err
		}

		obj.Total = totalInt
		arrobj = append(arrobj, obj)
	}

	// Total
	total, err := builders.GetTotalCount(con, baseTable, &where)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}
