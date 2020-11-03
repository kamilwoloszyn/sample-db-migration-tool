package models

type DbSchema struct {
	CustomerNumber         uint16
	CustomerName           string
	ContactLastName        string
	ContactFirstName       string
	Phone                  string
	AddressLine2           string
	City                   string
	State                  string
	Country                string
	SalesRepEmployeeNumber int
	CreditLimit            float32
}
