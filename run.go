package main

import (
	"fmt"

	"github.com/kamilwoloszyn/sample-db-migration-tool/api/db"
)

// Main function contians connection information about Lisa and Zen. Zen and lisa has the same database layouts.
func main() {
	fmt.Print("Starting ...")
	LisaDbInfo := db.Data{
		"mysql",
		"root",
		"trythisPASS",
		"Lisa",
	}

	ZenDbInfo := db.Data{
		"mysql",
		"root",
		"trythisPASS",
		"Zen",
	}

	zenDb, zenDbErr := ZenDbInfo.Connect()
	lisaDb, lisaDbErr := LisaDbInfo.Connect()

	if zenDbErr == nil && lisaDbErr == nil {
		if Lisa, err := db.GetData(lisaDb); err == nil {
			if mergeErr := db.Merge(Lisa, zenDb); err == nil {
				fmt.Println("Done!")
			} else {
				fmt.Printf("Found errors during merge: %s", string(mergeErr))
			}
		} else {
			fmt.Printf("Found errors during fetching data: %s", string(err))
		}

	} else {
		fmt.Printf(" Got state \n Lisa db: %s \n Zen db: %s", lisaDbErr, zenDbErr)
	}
	if zenDb != nil {
		defer zenDb.Close()
	}
	if lisaDb != nil {
		defer lisaDb.Close()
	}

}
