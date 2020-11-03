package connection

import (
	"database/sql"

	"github.com/kamilwoloszyn/sample-db-migration-tool/api/models"

	_ "github.com/go-sql-driver/mysql"
)

// LisaDBSchema extends models.DbSchema
// Extends main model DbSchema due to Golang not allow us to extend
// 'no local' structs. Designed for Lisa.
type LisaDBSchema models.DbSchema

// ZenDBSchema extends models.DbSchema
// Extends main model DbSchema due to Golang not allow us to extend
// 'no local' structs. Designed for Zen.
type ZenDBSchema models.DbSchema

type connectors interface {
	Connect()
}

// Connect () (db *sql.DB, err error)
// Function that extends struct LisaDBSchema which extends main model.
// It returns lisaDb no matter if error occured. Please check always
// if error is equal nil.
func (lisa *LisaDBSchema) Connect() (db *sql.DB, err error) {
	lisaDb, lisaErr := sql.Open("mysql", "root:trythisPASS@/Lisa")
	if lisaErr != nil {
		return lisaDb, lisaErr
	}
	return lisaDb, nil
}

// Connect () (db *sql.DB, err error)
// Same as above. Only with difference this is designed for Zen db.
func (zen *ZenDBSchema) Connect() (db *sql.DB, err error) {
	zenDb, zenErr := sql.Open("mysql", "root:trythisPASS@/Zen")
	if zenErr != nil {
		return nil, zenErr
	}
	return zenDb, nil
}
