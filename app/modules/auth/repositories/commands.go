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
	"fmt"
	"net/http"
)

func PostUserAuth(username, password string) (string, error, string) {
	status, msg := validations.GetValidateLogin(username, password)
	if status {
		// Declaration
		var obj models.UserLogin
		var pwd string
		var id string

		// Exec
		selectTemplate := builders.GetTemplateSelect("auth", nil, nil)
		baseTable := "users"
		sqlStatement := "SELECT id, " + selectTemplate + " " +
			"FROM " + baseTable +
			" WHERE username = ?"

		con := database.CreateCon()
		err := con.QueryRow(sqlStatement, username).Scan(
			&id, &obj.Username, &pwd,
		)

		if err == sql.ErrNoRows {
			return "", nil, "Account is not registered"
		} else if err != nil {
			return "", err, "Something went wrong. Please contact Admin"
		}

		match, err := auth.CheckPasswordHash(password, pwd)
		if !match {
			return "", nil, "Password incorrect"
		}

		if err != nil {
			return "", err, "Something went wrong. Please contact Admin"
		}

		return id, nil, ""
	} else {
		return "", nil, msg
	}
}

func PostUserRegister(body models.UserRegister) (response.Response, error) {
	var res response.Response
	status, msg := validations.GetValidateRegister(body)

	if status {
		// Declaration
		var baseTable = "users"
		id, err := generator.GenerateUUID(16)
		if err != nil {
			return res, err
		}

		createdAt := generator.GenerateTimeNow("timestamp")
		hashPass := auth.GenerateHashPassword(body.Password)

		// Query builder
		colFirstTemplate := builders.GetTemplateSelect("user_credential", nil, nil)
		colSecondTemplate := builders.GetTemplateSelect("user_mini_info", nil, nil)
		colThirdTemplate := builders.GetTemplateSelect("properties", &baseTable, nil)
		colFourthTemplate := builders.GetTemplateSelect("user_joined_info", &baseTable, nil)

		if err != nil {
			return res, err
		}

		sqlStatement := "INSERT INTO " + baseTable + " " +
			"(id, " + colFirstTemplate + ", " + colSecondTemplate + ", valid_until, " + colThirdTemplate +
			", deleted_at, deleted_by, " + colFourthTemplate + ", language_code) " + " " +
			"VALUES (?, ?, ?, ?, null, ?, ?, null, ?, ?, null, null, null, null, null, null, 0, 'en')"

		// Exec
		con := database.CreateCon()
		cmd, err := con.Prepare(sqlStatement)
		defer cmd.Close()

		if err != nil {
			return res, err
		}

		result, err := cmd.Exec(id, body.Username, body.Email, hashPass, body.FirstName, body.LastName, body.ValidUntil, createdAt)
		if err != nil {
			return res, err
		}

		rowsAffected, _ := result.RowsAffected()
		resultStr := fmt.Sprintf("%d", rowsAffected)

		// Response
		res.Status = http.StatusOK
		res.Message = generator.GenerateCommandMsg("account", "register", true)
		res.Data = map[string]string{"last_inserted_id": id, "result": resultStr + " rows affected"}
	} else {
		res.Status = http.StatusUnprocessableEntity
		res.Message = generator.GenerateCommandMsg("account. "+msg, "register", false)
		res.Data = map[string]string{"result": "0 rows affected"}
	}
	return res, nil
}

func PostAccessToken(body models.UserToken) error {
	// Declaration
	var baseTable = "users_tokens"
	id, err := generator.GenerateUUID(16)
	if err != nil {
		return err
	}
	createdAt := generator.GenerateTimeNow("timestamp")

	// Query builder
	colFirstTemplate := builders.GetTemplateSelect("user_access", nil, nil)

	sqlStatement := "INSERT INTO " + baseTable + " " +
		"(id, " + colFirstTemplate + ", token, last_used_at, created_at) " + " " +
		"VALUES (?, ?, ?, ?, null, ?)"

	// Exec
	con := database.CreateCon()
	cmd, err := con.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	result, err := cmd.Exec(id, body.ContextType, body.ContextId, body.Token, createdAt)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return err
	}
	return nil
}
