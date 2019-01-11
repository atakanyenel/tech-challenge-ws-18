package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type Ad struct {
	ImageURL   string
	Text       string
	Reason     string
	Title      string
	ProductURL string
}

func GetMainAds() []Ad {

	db, err := sql.Open("mysql", "root:example@tcp(mysql:3306)/test")
	if err != nil {
		panic(err)
	}

	var ads []Ad
	results, err := db.Query("select * from ads limit 4")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var comingAd Ad
		err = results.Scan(&comingAd.ImageURL, &comingAd.Text, &comingAd.Reason, &comingAd.Title, &comingAd.ProductURL)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		ads = append(ads, comingAd)

	}
	return ads
}
