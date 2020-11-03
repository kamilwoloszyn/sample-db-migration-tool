package main

import (
	"fmt"

	"github.com/kamilwoloszyn/sample-db-migration-tool/api/connection"
)

// Main function contians Zen and Lisa struct got from DB model. Zen and lisa has the same database layouts.
func main() {
	fmt.Print("Starting ...")
	Lisa := connection.LisaDBSchema{}
	Zen := connection.ZenDBSchema{}
	zenDb, zenDbErr := Zen.Connect()
	lisaDb, lisaDbErr := Lisa.Connect()

	if zenDbErr == nil && lisaDbErr == nil {
		defer zenDb.Close()
		defer lisaDb.Close()

	} else {
		fmt.Printf(" Got state \n Lisa db: %s \n Zen db: %s", lisaDbErr, zenDbErr)
		panic("Cannot connect to database!")
	}

}
