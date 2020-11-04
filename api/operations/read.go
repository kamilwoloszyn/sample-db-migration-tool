package operations

import (
	"database/sql"

	"github.com/kamilwoloszyn/sample-db-migration-tool/api/connection"
)

type reader interface {
	GetData() error
}

func (lisa *connection.LisaDBSchema) GetData(openedDB *sql.DB) error {
	if rows, err := openedDB.Query("select * from customers"); err == nil {
		
	}
}

func (zen *connection.ZenDBSchema) GetData(openedDB *sql.DB) error {

}
