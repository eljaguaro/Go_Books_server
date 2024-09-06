package repository

import (
	"goserv/pkg/models"
	"testing"
)

const connStr = "postgres://postgres2:David2004@localhost:5432/David"

func TestBooksCRUD(t *testing.T) {
	db, err := New(connStr)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = db.GetBooks()
	if err != nil {
		t.Fatal(err.Error())
	}

	book := models.Book{AuthorID: 1, GenreID: 1, Name: "Идиот", Price: 150}
	id, err := db.NewBook(book)
	if err != nil || id == 0 {
		t.Fatal(err.Error())
	}

	_, err = db.GetBookByID(id)
	if err != nil {
		t.Fatal(err.Error())
	}

}
