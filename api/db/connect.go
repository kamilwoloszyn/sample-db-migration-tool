package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Connect () (db *sql.DB, err error)
// Function that extends struct LisaDBSchema which extends main model.
// It returns lisaDb no matter if error occured. Please check always
// if error is equal nil.
func (d Data) Connect() (db *sql.DB, err error) {
	userDb, userErr := sql.Open(d.Driver, d.ToString())
	if userErr != nil {
		return nil, userErr
	}
	return userDb, nil
}

// ToString () returns a string in connection format. See above.
func (d Data) ToString() string {
	return fmt.Sprint(d.DbUser + ":" + d.DbPassword + "@/" + d.DbName)
}
