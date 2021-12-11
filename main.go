package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sample/acedamyCourse/database"
	"sample/acedamyCourse/models"
	"sample/acedamyCourse/utility"
)

var db *sql.DB

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home page endpoint hit")
}

func createTag(w http.ResponseWriter, r *http.Request){
	var tag *models.Tag
	err := json.NewDecoder(r.Body).Decode(&tag)
	if err != nil{
		utility.NewJSONWriter(w).Write(models.Response{
			Error: err.Error(),
			Message: "Error decoding request body",
		}, http.StatusBadRequest)
		return
	}

	database.CreateTag(db, tag)

	utility.NewJSONWriter(w).Write(models.Response{
		Error: "",
		Message: "Tag created successfully",
	}, http.StatusOK)
}

func handleRequests(){
	http.HandleFunc("/", homePage);
	http.HandleFunc("/tag", createTag)
	log.Fatal(http.ListenAndServe(":8081", nil))
}



func main(){
	var err error
	db, err = database.DBConnection();
	database.Setup(db)
	if err == nil{
		handleRequests()
	}

}

