package repositories

import (
	"app/modules/users/models"
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/converter"
	"app/packages/helpers/generator"
	"app/packages/helpers/response"
	"database/sql"
	"net/http"
)

func GetMyProfile(token string) (response.Response, error) {
	// Declaration
	var obj models.GetUserProfile
	var arrobj []models.GetUserProfile
	var res response.Response
	var baseTable = "users"
	var secondTable = "users_tokens"
	var sqlStatement string

	// Nullable column
	var lastName sql.NullString
	var imageURL sql.NullString
	var role sql.NullString
	var acceptedAt sql.NullString
	var updatedAt sql.NullString

	// Query builder
	firstSelectTemplate := builders.GetTemplateSelect("user_credential", nil, nil)
	secondSelectTemplate := builders.GetTemplateSelect("user_mini_info", nil, nil)
	firstLogicWhere := builders.GetTemplateLogic("active")
	whereLogic := baseTable + firstLogicWhere
	whereMine := builders.GetWhereMine(token) + " LIMIT 1"
	firstJoin := builders.GetTemplateJoin("total", "users", "id", secondTable, "context_id", false)

	sqlStatement = "SELECT " + firstSelectTemplate + ", " + secondSelectTemplate + ", valid_until, accepted_at, updated_at " +
		"FROM " + baseTable + " " +
		firstJoin +
		"WHERE " + whereLogic + " AND " + whereMine

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
			&obj.Username,
			&obj.Email,
			&obj.Password,
			&imageURL,
			&obj.FirstName,
			&lastName,
			&role,
			&obj.ValidUntil,
			&acceptedAt,
			&updatedAt,
		)
		if err != nil {
			return res, err
		}

		obj.ImageUrl = converter.CheckNullString(imageURL)
		obj.LastName = converter.CheckNullString(lastName)
		obj.Role = converter.CheckNullString(role)
		obj.AcceptedAt = converter.CheckNullString(acceptedAt)
		obj.UpdatedAt = converter.CheckNullString(updatedAt)

		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("User", len(arrobj))
	res.Data = arrobj

	return res, nil
}
