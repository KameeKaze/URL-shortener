package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func RoutesHandler() {
	r := mux.NewRouter()

	//routes
	r.PathPrefix("/css").Handler(http.StripPrefix("/css", http.FileServer(http.Dir("templates/css/")))).Methods("GET")
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/url", ShortURL).Methods("POST")

	fmt.Println("Running on http://127.0.0.1:" + "2222")
	http.ListenAndServe(":2222", r)

}

func Home(w http.ResponseWriter, r *http.Request) {
	//set status code
	w.WriteHeader(http.StatusOK)

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, "")
}

func CSS(w http.ResponseWriter, r *http.Request) {
	//set status code
	w.WriteHeader(http.StatusOK)

	tmpl, _ := template.ParseFiles("templates/css/styles.css")
	tmpl.Execute(w, "")
}

func ShortURL(w http.ResponseWriter, r *http.Request) {
	//set status code
	w.WriteHeader(http.StatusOK)
	//decode body data
	URL := r.FormValue("url")

	//parse url
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		w.Write([]byte("invalid url\n"))
		return
	}

	w.Write([]byte(URL))
}
