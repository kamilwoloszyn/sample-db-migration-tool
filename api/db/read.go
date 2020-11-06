package db

import (
	"database/sql"
)

type DataReader interface {
	GetData()
}

func GetData(openedDB *sql.DB, userSchema []Schema) error {
	rows, err := openedDB.Query("select * from customers")
	if err == nil {
		var i int
		for rows.Next() {
			if err := rows.Scan(userSchema[i].CustomerNumber, userSchema[i].CustomerName, userSchema[i].ContactLastName, userSchema[i].ContactFirstName, userSchema[i].Phone, userSchema[i].AddressLine1, userSchema[i].AddressLine2, userSchema[i].City, userSchema[i].State, userSchema[i].PostalCode, userSchema[i].Country, userSchema[i].SalesRepEmployeeNumber, userSchema[i].CreditLimit); err == nil {
				return nil
			}
			return err

		}
	}
	return err
}
