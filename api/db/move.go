package db

import (
	"database/sql"
	"fmt"
)

func Merge(fromData []Schema, to *sql.DB) error {

	if query, err := to.Prepare("INSERT INTO `customers` (`customerNumber`, `customerName`, `contactLastName`, `contactFirstName`, `phone`, `addressLine1`, `addressLine2`, `city`, `state`, `postalCode`, `country`, `salesRepEmployeeNumber`, `creditLimit`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"); err == nil {
		for _, item := range fromData {
			if respond, errQuery := query.Exec(item.CustomerNumber, item.CustomerName, item.ContactLastName, item.ContactFirstName, item.Phone, item.AddressLine1, item.AddressLine2, item.City, item.State, item.PostalCode, item.Country, item.SalesRepEmployeeNumber, item.CreditLimit); errQuery == nil {
				if id, errFetchID := respond.LastInsertId(); errFetchID == nil {
					fmt.Printf("Inserted id: %d\n", id)
				} else {
					return errFetchID
				}
			} else {
				return errQuery
			}
		}
	} else {
		return err
	}
}
