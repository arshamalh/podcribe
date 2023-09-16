package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"podcribe/repo"
	"strconv"
	"time"
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
	stmt, err := s.DB.Prepare("INSERT INTO podcast(page_link, podcast_link, provider, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(podcast.PageLink, podcast.Mp3Link, podcast.Provider, time.Now())
	if err != nil {
		return err
	}

	lastId, _ := result.LastInsertId()
	fmt.Println("Inserted", strconv.FormatInt(lastId, 10))

	return nil
}

func (s *sqlite) GetPodcastByPageLink(pageLink string) (podcast repo.Podcast, err error) {
	err = s.DB.QueryRow("SELECT id, podcast_link, provider FROM podcast WHERE page_link=?", pageLink).Scan(&podcast.Id, &podcast.Mp3Link, &podcast.Provider)
	if err != nil {
		return podcast, err
	}

	return podcast, err
}

func (s *sqlite) IncreasePodcastReferencedCount(podcastId int) (err error) {
	stmt, err := s.DB.Prepare("Update podcast  SET referenced_count = referenced_count + 1 WHERE id = ?")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(podcastId)
	if err != nil {
		return err
	}

	row, _ := result.RowsAffected()
	fmt.Println("Updated", strconv.FormatInt(row, 10))

	return nil
}
