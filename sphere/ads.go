package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Ad struct {
	ImageURL   string
	Text       string
	Reason     string
	Usage      int
	Title      string
	ProductURL string
}

func GetMainAds() []Ad {

	db, err := sql.Open("mysql", "root:example@tcp(mysql:3306)/local")
	if err != nil {
		panic(err)
	}

	var ads []Ad
	results, err := db.Query("select * from ads")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var comingAd Ad
		var id int
		err = results.Scan(&id, &comingAd.ImageURL, &comingAd.Text, &comingAd.Reason, &comingAd.Usage, &comingAd.Title, &comingAd.ProductURL)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		ads = append(ads, comingAd)

	}
	socketsUsage := returnByType(db)

	var returnedads []Ad
	for k, v := range socketsUsage {
		for _, a := range ads {
			if a.Reason == k && a.Usage < v {

				returnedads = append(returnedads, a)
			}
		}
	}
	return returnedads

}

func SimAds() []Ad {
	ads := []Ad{
		{"https://images-na.ssl-images-amazon.com/images/I/51B5x6fi4HL._SL1500_.jpg", "This is the new PS4.", "Entertainment", 14, "Playstation 4", "https://www.amazon.de/dp/B07KMV94JF/ref=asc_df_B07KMV94JF57979247/?tag=googshopde-21&creative=22434&creativeASIN=B07KMV94JF&linkCode=df0&hvadid=308847264890&hvpos=1o1&hvnetw=g&hvrand=16041305715093971334&hvpone=&hvptwo=&hvqmt=&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=9042512&hvtargid=pla-647432777687&th=1&psc=1&tag=&ref=&adgrpid=64245981791&hvpone=&hvptwo=&hvadid=308847264890&hvpos=1o1&hvnetw=g&hvrand=16041305715093971334&hvqmt=&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=9042512&hvtargid=pla-647432777687"},
		{"https://images-na.ssl-images-amazon.com/images/I/511F7fH4BbL._SL1000_.jpg", "A smart Philip Hue light", "Lighting", 7, "Philips HF3520/01", "https://www.amazon.de/dp/B008LR3KD8/ref=asc_df_B008LR3KD857981926/?tag=googshopde-21&creative=22398&creativeASIN=B008LR3KD8&linkCode=df0&hvadid=232067132146&hvpos=1o1&hvnetw=g&hvrand=3860014294753512637&hvpone=&hvptwo=&hvqmt=&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=9042512&hvtargid=pla-386142310710&th=1&psc=1"},
		{"https://images-na.ssl-images-amazon.com/images/I/61maXsV3t9L._SL1500_.jpg", "An economic refrigerator for your kitchen", "Cooking", 2, "Bomann KG 320.1 Kühl-Gefrier-Kombination", "https://www.amazon.de/Bomann-KG-320-1-K%C3%BChl-Gefrier-Kombination-Gefrierteil/dp/B06Y2TVW9Y/ref=sr_1_8?s=appliances&ie=UTF8&qid=1547224368&sr=1-8&keywords=refrigerator"},
		{"https://images-na.ssl-images-amazon.com/images/I/91wEhBwdFfL._SL1500_.jpg", "A new TV for your home", "Entertainment", 10, "Dyon Enter 32 Pro-X 80 cm (32 Zoll) Fernseher", "https://www.amazon.de/Enter-Pro-X-Fernseher-Triple-Energieklasse/dp/B07CZ1YWVG/ref=sr_1_8?s=ce-de&ie=UTF8&qid=1547224470&sr=1-8&keywords=tv"},
	}

	sockets := map[string]int{
		"Lighting":      8,
		"Charging":      15,
		"Entertainment": 11,
	}
	var returnedads []Ad
	for k, v := range sockets {
		for _, a := range ads {
			if a.Reason == k && a.Usage < v {

				returnedads = append(returnedads, a)
			}
		}
	}
	fmt.Println(returnedads)
	return ads
}
