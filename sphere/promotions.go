package main

import (
	"database/sql"
)

type promotion struct {
	ID    int
	Title string
	Body  string
}

func returnPromotions(db *sql.DB) []promotion {

	results, err := db.Query("select * from promotions")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var promotions []promotion
	for results.Next() {
		var comingNotif promotion
		err = results.Scan(&comingNotif.ID, &comingNotif.Title, &comingNotif.Body)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		promotions = append(promotions, comingNotif)
	}
	return promotions

}
