package dbaccess

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBAccess() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cicd")

	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}
