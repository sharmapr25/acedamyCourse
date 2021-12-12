package models

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestShouldCreateTag(t *testing.T){
	db, mock, err := sqlmock.New()
	if err != nil{
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectPrepare("INSERT").ExpectExec().WithArgs("Tag-2", "1639225436", true).WillReturnResult(sqlmock.NewResult(1,1));

	tag := &Tag{
	 	Name: "Tag-2",
		CreatedAt: "1639225436",
		IsActive: true,
	}

	tag.CreateTag(db);

	if err := mock.ExpectationsWereMet(); err != nil{
		t.Errorf("There are unfulfilled expectation: %s", err)
	}
}