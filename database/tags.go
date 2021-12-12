package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


type TagStore struct{
	db *sql.DB
}

func NewTagStore(db *sql.DB) *TagStore{
	return &TagStore{db: db}
}
