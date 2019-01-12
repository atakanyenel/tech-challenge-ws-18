package main

import (
	"database/sql"
)

type repair struct {
	ID     int
	Status string
	Title  string
	Body   string
}

func returnRepairs(db *sql.DB) []repair {

	results, err := db.Query("select * from repairs")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var repairs []repair
	for results.Next() {
		var comingRepair repair
		err = results.Scan(&comingRepair.ID, &comingRepair.Status, &comingRepair.Title, &comingRepair.Body)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		repairs = append(repairs, comingRepair)

	}
	return repairs

}
