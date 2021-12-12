package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sample/acedamyCourse/models"
	"sample/acedamyCourse/utility"
)


func CreateTagApi(db *sql.DB, w http.ResponseWriter, r *http.Request){
	tag := models.Tag{}
	err := json.NewDecoder(r.Body).Decode(&tag)
	if err != nil{
		utility.NewJSONWriter(w).Write(models.Response{
			Error: err.Error(),
			Message: "Error decoding request body",
		}, http.StatusBadRequest)
		return
	}


	tag.CreateTag(db)

	utility.NewJSONWriter(w).Write(models.Response{
		Error: "",
		Message: "Tag created successfully",
	}, http.StatusOK)
}