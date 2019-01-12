package main

import (
	"database/sql"
)

type notif struct {
	ID    int
	Title string
	Body  string
}

func returnNotifs(db *sql.DB) []notif {

	results, err := db.Query("select * from notifs")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var repairs []notif
	for results.Next() {
		var comingNotif notif
		err = results.Scan(&comingNotif.ID, &comingNotif.Title, &comingNotif.Body)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		repairs = append(repairs, comingNotif)
	}
	return repairs

}
