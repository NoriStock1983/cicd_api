package dbaccess

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBAccess() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cicd")

	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return db, nil
	}
}
