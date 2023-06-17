package builders

import "database/sql"

func GetTotalCount(con *sql.DB, table string, view string) (int, error) {
	var count int

	sqlStatement := "SELECT COUNT(*) FROM " + table + " WHERE " + view

	err := con.QueryRow(sqlStatement).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
