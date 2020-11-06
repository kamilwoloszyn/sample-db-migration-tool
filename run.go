package main

import (
	"fmt"

	"github.com/kamilwoloszyn/sample-db-migration-tool/api/db"
)

// Main function contians Zen and Lisa struct got from DB model. Zen and lisa has the same database layouts.
func main() {
	fmt.Print("Starting ...")
	LisaDbInfo= db.Data{
		"mysql"
		"root"
		"trythisPASS"
		"Lisa"
	}

	ZenDbInfo:= db.Data{
		"mysql"
		"root"
		"trythisPASS"
		"Zen"
	}

	Lisa := db.DbSchema{}
	Zen := db.DbSchema{}

	zenDb, zenDbErr := ZenDbInfo.Connect()
	lisaDb, lisaDbErr := LisaDbInfo.Connect()

	if zenDbErr == nil && lisaDbErr == nil {
		defer zenDb.Close()
		defer lisaDb.Close()

	} else {
		fmt.Printf(" Got state \n Lisa db: %s \n Zen db: %s", lisaDbErr, zenDbErr)
		panic("Cannot connect to database!")
	}

}
