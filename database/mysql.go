package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


func DBConnection() (*sql.DB, error){
	dbUser := "root"
	dbPass := ""
	dbName := "acedamy"
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
		return nil, err
	}
	log.Println("Connected to database")
	return db, nil
}

func Setup(db *sql.DB){
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS Tag(Name TEXT NOT NULL, CreatedAt TEXT NOT NULL, IsActive TEXT NOT NULL)")
	if err != nil{
		panic(err)
	}
	log.Println("Table Tag has been created")
}