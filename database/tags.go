package database

import (
	"database/sql"
	"fmt"
	"log"
	"sample/acedamyCourse/models"

	_ "github.com/go-sql-driver/mysql"
)


type TagStore struct{
	db *sql.DB
}

func NewTagStore(db *sql.DB) *TagStore{
	return &TagStore{db: db}
}

func CreateTag(db *sql.DB, tag *models.Tag){
	tagForm, err := db.Prepare("INSERT INTO Tag(Name, CreatedAt, IsActive) values(?,?,?)")

	if err != nil{
		panic(err.Error())
	}
	tagForm.Exec(tag.Name, tag.CreatedAt, tag.IsActive)
	log.Printf(fmt.Sprintf("Insert: Name: %s, CreatedAt: %s, IsActive: %t", tag.Name, tag.CreatedAt, tag.IsActive));
}

