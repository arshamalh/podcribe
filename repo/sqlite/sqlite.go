package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"podcribe/repo"
)

// If the link was the same as database, don't download that again, using already downloaded files
// Keep track of every downloaded files

type sqlite struct {
	DB *sql.DB
}

func New() (*sqlite, error) {
	db, err := sql.Open("sqlite3", "./podcribe.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to db")

	err = createTables(db)
	if err != nil {
		return nil, err
	}

	fmt.Println("create tables")

	return &sqlite{DB: db}, nil
}

func createTables(db *sql.DB) (err error) {
	sts := `
				CREATE TABLE IF NOT EXISTS podcast(id INTEGER PRIMARY KEY, page_link varchar, podcast_link varchar, provider varchar, path varchar, referenced_count varchar, created_at timestamp);
				CREATE TABLE IF NOT EXISTS user(id INTEGER PRIMARY KEY, created_at timestamp);
				CREATE TABLE IF NOT EXISTS user_podcast(id INTEGER PRIMARY KEY, user_id int, podcast_id int);
			`
	_, err = db.Exec(sts)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlite) StorePodcast(podcast repo.Podcast) (err error) {
	stmt, err := s.DB.Prepare("INSERT INTO podcast(page_link, podcast_link, provider) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(podcast.PageLink, podcast.PodcastLink, podcast.Provider)
	if err != nil {
		return err
	}

	fmt.Println("Inserted")

	return nil
}
