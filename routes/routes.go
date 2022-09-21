package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KameeKaze/URL-shortener/types"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func RoutesHandler() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//routes
	r.Get("/", Home)

	//start
	fmt.Println("Running on http://127.0.0.1:" + "3000")
	http.ListenAndServe(":3000", r)
}

func Home(w http.ResponseWriter, r *http.Request) {
	createHttpResponse(w, http.StatusOK, "Simple URL shortener")
}

func createHttpResponse(w http.ResponseWriter, statusCode int, text string) {
	//set status code
	w.WriteHeader(statusCode)
	//create json
	r, _ := json.Marshal(types.HTTPResponse{
		Msg: text,
	})
	//send data
	w.Write([]byte(r))
}
