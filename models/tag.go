package models

import (



	_ "github.com/go-sql-driver/mysql"
)

type Tag struct{
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	IsActive bool `json:"is_active"`
}