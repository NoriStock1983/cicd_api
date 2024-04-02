package dbaccess

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBAccess(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cicd")

	assert.Nil(t, err)

	defer db.Close()

	err = db.Ping()
	assert.Nil(t, err)
}
