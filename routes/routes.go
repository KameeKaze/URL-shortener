package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/KameeKaze/URL-shortener/types"
)

func RoutesHandler() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//routes
	r.Get("/", Home)
	r.Post("/url", URL)

	//start
	fmt.Println("Running on http://127.0.0.1:" + "3000")
	http.ListenAndServe(":3000", r)
}

func Home(w http.ResponseWriter, r *http.Request) {
	//set status code
	w.WriteHeader(http.StatusOK)

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, "")
}

func URL(w http.ResponseWriter, r *http.Request) {
	//set status code
	w.WriteHeader(http.StatusOK)
	//decode body data
	body := &types.URL{}
	json.NewDecoder(r.Body).Decode(&body)
	fmt.Println(body)
	w.Write([]byte("asd\n"))

}
