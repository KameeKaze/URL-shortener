package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/KameeKaze/URL-shortener/types"
	"github.com/gorilla/mux"
)

func RoutesHandler() {
	r := mux.NewRouter()

	//routes
	r.PathPrefix("/css").Handler(http.StripPrefix("/css", http.FileServer(http.Dir("templates/css")))).Methods("GET")
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/url", URL).Methods("POST")

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

func URL(w http.ResponseWriter, r *http.Request) {
	//set status code
	w.WriteHeader(http.StatusOK)
	//decode body data
	body := &types.URL{}
	json.NewDecoder(r.Body).Decode(&body)
	w.Write([]byte("asd\n"))
}
