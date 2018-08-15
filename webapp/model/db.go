package model

import (
	"database/sql"
)

var db *sql.DB

// SetDatabase registers db handler
// to the model layer for better testing
func SetDatabase(dbInstance *sql.DB) {
	db = dbInstance
}
