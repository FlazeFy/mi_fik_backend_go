package converter

import (
	"database/sql"
)

func CheckNullString(data sql.NullString) string {
	var res string
	if data.Valid {
		res = data.String
	} else {
		res = ""
	}

	return res
}
