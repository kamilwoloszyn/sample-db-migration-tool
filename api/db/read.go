package db

import (
	"database/sql"
)

func GetData(openedDB *sql.DB) ([]Schema, error) {

	size, errCount := CountData("SELECT COUNT(*) FROM customers", openedDB)
	if errCount == nil && size > 0 {
		if rows, errQuery := openedDB.Query("select * from customers"); errQuery == nil {
			var userSchema []Schema = make([]Schema, size+1)
			var i int
			for rows.Next() {
				if err := rows.Scan(&userSchema[i].CustomerNumber, &userSchema[i].CustomerName, &userSchema[i].ContactLastName, &userSchema[i].ContactFirstName, &userSchema[i].Phone, &userSchema[i].AddressLine1, &userSchema[i].AddressLine2, &userSchema[i].City, &userSchema[i].State, &userSchema[i].PostalCode, &userSchema[i].Country, &userSchema[i].SalesRepEmployeeNumber, &userSchema[i].CreditLimit); err == nil {
					i++
				} else {
					return userSchema, err
				}

			}
		} else {
			return nil, errQuery
		}

	}

	return nil, errCount

}

func CountData(queryStr string, openedDB *sql.DB) (int, error) {
	var result int
	query, err := openedDB.Query(queryStr)
	if query != nil && err == nil {
		query.Next()
		if errorScan := query.Scan(&result); errorScan == nil {
			return result, nil
		} else {
			return 0, errorScan
		}
	}
	return 0, err
}
