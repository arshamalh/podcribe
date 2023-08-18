package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// If the link was the same as database, don't download that again, using already downloaded files
// Keep track of every downloaded files

type sqlite struct {
	DB *sql.DB
}

func New() (*sqlite, error) {
	//TODO: fix the .db name file and path
	sqliteDb, err := sql.Open("sqlite3", "./mydb.db")
	if err != nil {
		return nil, err
	}

	err = sqliteDb.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to db")

	return &sqlite{DB: sqliteDb}, nil
}
