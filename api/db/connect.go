package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connect () (db *sql.DB, err error)
// Function that extends struct LisaDBSchema which extends main model.
// It returns lisaDb no matter if error occured. Please check always
// if error is equal nil.
func (d Data) Connect() (db *sql.DB, err error) {
	lisaDb, lisaErr := sql.Open(d.Driver, "root:trythisPASS@/Lisa")
	if lisaErr != nil {
		return lisaDb, lisaErr
	}
	return lisaDb, nil
}
