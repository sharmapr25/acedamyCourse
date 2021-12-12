package models

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type Tag struct{
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	IsActive bool `json:"is_active"`
}

func(tag *Tag) CreateTag(db *sql.DB){
	tagForm, err := db.Prepare("INSERT INTO Tag(Name, CreatedAt, IsActive) values(?,?,?)")
	
	if err != nil{
		panic(err.Error())
	}
	tagForm.Exec(tag.Name, tag.CreatedAt, tag.IsActive)
	log.Printf(fmt.Sprintf("Insert: Name: %s, CreatedAt: %s, IsActive: %t", tag.Name, tag.CreatedAt, tag.IsActive));
}

