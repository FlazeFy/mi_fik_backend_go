package repositories

import (
	"app/modules/auth/models"
	"app/modules/auth/validations"
	"app/packages/builders"
	"app/packages/database"
	"app/packages/helpers/auth"
	"app/packages/helpers/generator"
	"app/packages/helpers/response"
	"database/sql"
	"net/http"
	"time"
)

func PostUserAuth(username, password string) (bool, error, string) {
	var obj models.UserLogin
	var pwd string

	con := database.CreateCon()

	selectTemplate := "username, password "
	baseTable := "users"
	sqlStatement := "SELECT " + selectTemplate +
		"FROM " + baseTable +
		" WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		return false, err, "Account is not registered"
	}

	if err != nil {
		return false, err, "Something wrong. Please contact Admin"
	}

	match, err := auth.CheckPasswordHash(password, pwd)
	if !match {
		return false, err, "Username or password incorrect"
	}

	return true, nil, ""
}

func PostUserRegister(body models.UserRegister) (response.Response, error) {
	con := database.CreateCon()
	var res response.Response
	status, msg := validations.GetValidateRegister(body)

	if status {
		var baseTable = "users"
		colFirstTemplate := builders.GetTemplateGeneralSelect("user_credential", nil)
		colSecondTemplate := builders.GetTemplateGeneralSelect("user_mini_info", nil)
		colThirdTemplate := builders.GetTemplateGeneralSelect("properties", &baseTable)
		colFourthTemplate := builders.GetTemplateGeneralSelect("user_joined_info", &baseTable)
		id, err := generator.GenerateUUID(16)
		now := time.Now().Unix()

		if err != nil {
			return res, err
		}

		sqlStatement := "INSERT INTO " + baseTable + " " +
			"(id, " + colFirstTemplate + ", " + colSecondTemplate + ", valid_until, " + colThirdTemplate +
			", deleted_at, deleted_by, " + colFourthTemplate + ") " + " " +
			"VALUES (?, null, ?, ?, ?, null, ?, ?, null, ?, ?, null, null, null, null, null, null, 0)"

		cmd, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := cmd.Exec(id, body.Username, body.Email, body.Password, body.FirstName, body.LastName, body.ValidUntil, now)
		if err != nil {
			return res, err
		}

		lastInsertedId, err := result.LastInsertId()
		if err != nil {
			return res, err
		}

		res.Status = http.StatusOK
		res.Message = generator.GenerateCommandMsg("account", "register", true)
		res.Data = map[string]int64{
			"id": lastInsertedId,
			// "detail":body
		}
	} else {
		res.Status = http.StatusUnprocessableEntity
		res.Message = generator.GenerateCommandMsg("account "+msg, "register", false)
	}
	return res, nil
}
