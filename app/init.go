package app

import (
	"database/sql"
	"log"
	"net/http"
	"sample/acedamyCourse/database"
	"github.com/gorilla/mux"
)

type App struct {
	DB *sql.DB
	Router *mux.Router
}

func(a *App) Initialize(){
	db, err := database.DBConnection();
	if err == nil{
		database.Setup(db)
		a.DB = db
		a.Router = mux.NewRouter()
		a.InitializeRoutes()
	}
}

func(a *App) Run(){
	log.Fatal(http.ListenAndServe(":8081", a.Router))
}

