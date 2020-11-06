package db

// Schema of database, which is the same for Zen and Lisa.
type Schema struct {
	CustomerNumber         uint16
	CustomerName           string
	ContactLastName        string
	ContactFirstName       string
	Phone                  string
	AddressLine1           string
	AddressLine2           string
	City                   string
	State                  string
	PostalCode             string
	Country                string
	SalesRepEmployeeNumber int
	CreditLimit            float32
}

// Data is a struct which pass into Connect function.
type Data struct {
	Driver     string
	DbUser     string
	DbPassword string
	DbName     string
}
