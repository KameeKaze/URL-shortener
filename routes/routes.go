package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"

	"github.com/KameeKaze/URL-shortener/db"
	"github.com/KameeKaze/URL-shortener/utils"
	"github.com/gorilla/mux"
)

type link struct {
	URL string
}

func RoutesHandler() {
	r := mux.NewRouter()

	//routes
	r.PathPrefix("/css").Handler(http.StripPrefix("/css", http.FileServer(http.Dir("templates/css/")))).Methods("GET")
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/", ShortURL).Methods("POST")
	r.HandleFunc("/{short}", GetURL).Methods("GET")

	fmt.Println("Running on http://127.0.0.1:" + os.Getenv("PORT"))
	http.ListenAndServe(":2000", r)

}

func Home(w http.ResponseWriter, r *http.Request) {
	//set status code
	w.WriteHeader(http.StatusOK)

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, "")
}

func ShortURL(w http.ResponseWriter, r *http.Request) {
	//decode body data
	URL := r.FormValue("url")

	//parse url
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		w.Write([]byte("invalid url\n"))
		return
	}
	// generate random string for uri
	URI := utils.RandStringBytes()

	// save url in database
	err = db.Redis.SetURL(URI, URL)
	if err != nil {
		w.Write([]byte("database error\n"))
		fmt.Println(err)
		return
	}
	// display generated URI
	tmpl, _ := template.ParseFiles("templates/url.html")
	tmpl.Execute(w, link{URL: URI})
}

func GetURL(w http.ResponseWriter, r *http.Request) {
	// get short link
	URI := mux.Vars(r)["short"]
	// get the redirect
	URL, err := db.Redis.GetURL(URI)
	if err != nil {
		w.Write([]byte("database error\n"))
		fmt.Println(err)
		return
	}
	// redirect to URL
	http.Redirect(w, r, URL, http.StatusMovedPermanently)
}
