package sqlz

import (
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"testing"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func TestMapRows(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Alice").
		AddRow(2, "Bob")

	mock.ExpectQuery("SELECT .* FROM users").WillReturnRows(rows)

	result, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		t.Fatal(err)
	}

	var users []User
	err = Select(result, &users)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(users)

	if len(users) != 2 || users[0].Name != "Alice" {
		t.Error("Unexpected result:", users)
	}
}
