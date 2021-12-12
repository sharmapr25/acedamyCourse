package app

import (
	"database/sql"
	"net/http"
	"sample/acedamyCourse/api"
)

func(a *App) InitializeRoutes(){
	a.Post("/tags", a.handleRequests(api.CreateTagApi))
}

func(a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)){
	a.Router.HandleFunc(path, f).Methods("POST")
}

type RequestHandlerFunc func(db *sql.DB, w http.ResponseWriter, r *http.Request)

func(a *App) handleRequests(handler RequestHandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}