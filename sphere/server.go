package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var user struct {
	Name    string
	Surname string
	ID      int
}

func startServer() {
	user.Name = "Example"
	user.Surname = "User"
	user.ID = 1

	r := gin.Default()
	db, err := sql.Open("mysql", "root:example@tcp(mysql:3306)/local")
	if err != nil {
		panic(err)
	}
	r.Use(static.Serve("/", static.LocalFile("./views/", false)))
	r.LoadHTMLGlob("views/*.html")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(200, "home.html", gin.H{"ads": GetMainAds(), "user": user})
	})

	r.GET("/line-chart", func(c *gin.Context) {
		c.HTML(http.StatusOK, "usage-by-day.html", gin.H{})
	})

	r.GET("/insert", func(c *gin.Context) {
		receivedMessage := socketData{
			Name: 2, Time: time.Now()}
		stmt, err := db.Prepare("INSERT measurements SET socket_id=?,time=?")
		checkErr(err)

		res, err := stmt.Exec(receivedMessage.Name, receivedMessage.Time)
		fmt.Println(res)
		fmt.Println(err)
		c.JSON(200, "ok")
	})

	r.GET("/by-type", func(c *gin.Context) {

		c.JSON(200, returnByType(db))
	})

	r.GET("/bar-chart", func(c *gin.Context) {
		results, err := db.Query("select count(socket_id),date(time) from measurements group by date(time)")
		if err != nil {
			panic(err.Error())
		}
		dataArray := make(map[string]int)

		var time string
		var count int
		for results.Next() {
			err = results.Scan(&count, &time)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			dataArray[time] = count

		}

		c.JSON(200, gin.H{"values": dataArray})
	})

	r.GET(("/usage-by-day"), func(c *gin.Context) {
		results, err := db.Query("select s.type,count(*),date(time) as day from measurements m , sockets s where s.socket_id=m.socket_id  group by day,type")

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		var typex string
		var date string
		var count int
		var dates []string
		dataArray := make(map[string]map[string]int)
		for results.Next() {

			// for each row, scan the result into our tag composite object
			err = results.Scan(&typex, &count, &date)
			dates = append(dates, date)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			inner, ok := dataArray[typex]
			if !ok {
				inner = make(map[string]int)
				dataArray[typex] = inner
			}
			inner[date] = count
			// and then print out the tag's Name attribute

		}
		dates = removeDuplicates(dates)

		for t := range dataArray {
			for _, v := range dates {
				if _, ok := dataArray[t][v]; !ok {
					dataArray[t][v] = 0
				}
			}
		}
		c.JSON(200, gin.H{"labels": dates, "values": dataArray})
	})

	socketRoutes := r.Group("sockets")
	{
		socketRoutes.GET("/", func(c *gin.Context) {
			type socket struct {
				ID     int
				Typex  string
				Status string
			}
			results, err := db.Query("select * from sockets")
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			var sockets []socket
			for results.Next() {
				var s socket
				err = results.Scan(&s.ID, &s.Typex, &s.Status)
				if err != nil {
					panic(err.Error()) // proper error handling instead of panic in your app
				}
				sockets = append(sockets, s)

			}

			c.JSON(200, sockets)
		})
	}

	r.GET("/ads", func(c *gin.Context) {
		//c.HTML(200, "ads.html", GetMainAds())
		c.HTML(200, "index.html", gin.H{"ads": SimAds()})
	})

	r.GET("/charts", func(c *gin.Context) { c.HTML(200, "charts.html", gin.H{"user": user}) })
	r.GET("/tables", func(c *gin.Context) { c.HTML(200, "tables.html", gin.H{"user": user}) })
	r.GET("/compare", func(c *gin.Context) { c.HTML(200, "compare.html", gin.H{"user": user}) })
	r.GET("/notif", func(c *gin.Context) { c.HTML(200, "notif.html", gin.H{"user": user, "data": returnNotifs(db)}) })
	r.GET("/promotions", func(c *gin.Context) {
		c.HTML(200, "promotions.html", gin.H{"user": user, "data": returnPromotions(db)})
	})
	r.GET("/repair", func(c *gin.Context) { c.HTML(200, "repair.html", gin.H{"user": user, "data": returnRepairs(db)}) })
	r.GET("/test", func(c *gin.Context) { c.HTML(200, "test.html", gin.H{"user": user, "data": returnNotifs(db)}) })
	simRoutes(r)
	r.Run(":4000")
}

func simRoutes(r *gin.Engine) {
	sim := r.Group("/sim")
	{
		sim.GET("/by-type", func(c *gin.Context) {
			c.JSON(200, map[string]int{
				"lighting": 4,
				"charging": 15,
			})
		})

		sim.GET("/usage-by-day", func(c *gin.Context) {
			type dayUsage struct {
				Typex string
				Count int
			}

			dataArray := map[string]map[string]int{
				"charging": {
					"2019-01-07": 6,
					"2019-01-08": 61,
				},
				"lighting": {
					"2019-01-07": 1,
				},
			}
			labels := []string{"2019-01-07", "2019-01-08"}
			for t := range dataArray {
				for _, v := range labels {
					if _, ok := dataArray[t][v]; !ok {
						dataArray[t][v] = 0
					}
				}
			}
			c.JSON(200, gin.H{"labels": labels, "values": dataArray})
		})
		sim.GET("/sockets", func(c *gin.Context) {
			type socket struct {
				ID     int
				Typex  string
				Status string
			}
			sockets := []socket{
				{1, "entertainment", "ACTIVE"},
				{2, "LIGHT", "DEACTIVE"},
				{3, "COOKING", "ACTIVE"},
				{4, "ENTERTAINMENT", "FAIL"},
			}

			c.JSON(200, sockets)
		})
		sim.GET("/ads", func(c *gin.Context) {
			c.HTML(200, "ads.html", SimAds())
		})
	}

}

//select s.type,count(status),date(time) as day from measurements m , sockets s where s.socket_id=m.socket_id and m.status=1 group by day,type

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func returnByType(db *sql.DB) map[string]int {
	results, err := db.Query("select count(*),type from measurements,sockets where sockets.socket_id=measurements.socket_id group by type")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	data := make(map[string]int)

	for results.Next() {

		var typex string
		var count int
		// for each row, scan the result into our tag composite object
		err = results.Scan(&count, &typex)
		data[typex] = count
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
	}
	return data
}

func returnHandler(route string, data interface{}) func(c *gin.Context) {

	return func(c *gin.Context) {
		c.HTML(200, fmt.Sprintf("%v.html", route), gin.H{"user": user, route: data})
	}
}
