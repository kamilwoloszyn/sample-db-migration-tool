package connection

import (
	"database/sql"

	"github.com/kamilwoloszyn/sample-db-migration-tool/api/models"

	_ "github.com/go-sql-driver/mysql"
)

type LisaDBSchema models.DbSchema
type ZenDBSchema models.DbSchema

type connectors interface {
	Connect()
}

func (lisa *LisaDBSchema) Connect() (db *sql.DB, err error) {
	lisaDb, lisaErr := sql.Open("mysql", "root:trythisPASS@/Lisa")
	if lisaErr != nil {
		return lisaDb, lisaErr
	}
	return lisaDb, nil
}

func (zen *ZenDBSchema) Connect() (db *sql.DB, err error) {
	zenDb, zenErr := sql.Open("mysql", "root:trythisPASS@/Zen")
	if zenErr != nil {
		return nil, zenErr
	}
	return zenDb, nil
}
