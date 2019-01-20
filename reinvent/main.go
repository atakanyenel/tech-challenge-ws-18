package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Ad struct {
	ID         int
	ImageURL   string
	Text       string
	Reason     string
	Usage      int
	Title      string
	ProductURL string
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:example@tcp(mysql:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM ads")
	if err != nil {
		panic(err.Error())
	}
	comingAd := Ad{}
	res := []Ad{}
	for selDB.Next() {

		err = selDB.Scan(&comingAd.ID, &comingAd.ImageURL, &comingAd.Text, &comingAd.Reason, &comingAd.Usage, &comingAd.Title, &comingAd.ProductURL)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, comingAd)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM ads WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	comingAd := Ad{}
	for selDB.Next() {

		err = selDB.Scan(&comingAd.ID, &comingAd.ImageURL, &comingAd.Text, &comingAd.Reason, &comingAd.Usage, &comingAd.Title, &comingAd.ProductURL)
		if err != nil {
			panic(err.Error())
		}

	}
	tmpl.ExecuteTemplate(w, "Show", comingAd)
	defer db.Close()
}

func new(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM ads WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	comingAd := Ad{}
	for selDB.Next() {

		err = selDB.Scan(&comingAd.ID, &comingAd.ImageURL, &comingAd.Text, &comingAd.Reason, &comingAd.Usage, &comingAd.Title, &comingAd.ProductURL)
		if err != nil {
			panic(err.Error())
		}
	}
	tmpl.ExecuteTemplate(w, "Edit", comingAd)
	defer db.Close()
}

func insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		image := r.FormValue("image")
		text := r.FormValue("text")
		reason := r.FormValue("reason")
		usage := r.FormValue("usage")
		title := r.FormValue("title")
		product := r.FormValue("product")
		insForm, err := db.Prepare("INSERT `ads` (`image_url`,`text`,`type`,`usage`,`title`,`product_url`) VALUES(?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(image, text, reason, usage, title, product)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		image := r.FormValue("image")
		text := r.FormValue("text")
		reason := r.FormValue("reason")
		usage := r.FormValue("usage")
		title := r.FormValue("title")
		product := r.FormValue("product")
		id := r.FormValue("uid")

		insForm, err := db.Prepare("UPDATE `ads` SET `image_url`=?, `text`=?,`type`=?,`usage`=?,`title`=?,`product_url`=? WHERE `id`=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(image, text, reason, usage, title, product, id)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM ads WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
func getAdsAPI(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM ads")
	if err != nil {
		panic(err.Error())
	}
	comingAd := Ad{}
	res := []Ad{}
	for selDB.Next() {

		err = selDB.Scan(&comingAd.ID, &comingAd.ImageURL, &comingAd.Text, &comingAd.Reason, &comingAd.Usage, &comingAd.Title, &comingAd.ProductURL)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, comingAd)
	}
	js, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func main() {
	log.Println("Server started on: http://localhost:5000")
	http.HandleFunc("/", index)
	http.HandleFunc("/show", show)
	http.HandleFunc("/new", new)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/api/ads", getAdsAPI)
	http.ListenAndServe(":5000", nil)
}
