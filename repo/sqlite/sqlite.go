package sqlite

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
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
	db, err := sql.Open("sqlite3", "podcribe")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to db")

	err = runMigration(db)
	if err != nil {
		return nil, err
	}

	fmt.Println("create tables")

	return &sqlite{DB: db}, nil
}

func runMigration(db *sql.DB) (err error) {
	dbDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		fmt.Printf("Instance file error: %v\n", err)
		return err
	}

	fileSource, err := (&file.File{}).Open("./repo/migration")
	if err != nil {
		fmt.Printf("opening file error: %v\n", err)
		return err
	}

	m, err := migrate.NewWithInstance("file", fileSource, "podcribe", dbDriver)
	if err != nil {
		fmt.Printf("migrate error: %v\n", err)
		return err
	}

	if err = m.Up(); err != nil {
		fmt.Printf("migrate up error: %v\n", err)
	}

	fmt.Printf("migrate up done with success")

	return nil
}

func (s *sqlite) StorePodcast(podcast repo.Podcast) (err error) {
	stmt, err := s.DB.Prepare("INSERT INTO podcasts(page_link, mp3_link, provider, created_at) VALUES (?, ?, ?, ?)")
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
	err = s.DB.QueryRow("SELECT id, mp3_link, provider FROM podcasts WHERE page_link=?", pageLink).Scan(&podcast.Id, &podcast.Mp3Link, &podcast.Provider)
	if err != nil {
		return podcast, err
	}

	fmt.Println("podcast by page link", podcast.Id)

	return podcast, err
}

func (s *sqlite) IncreasePodcastReferencedCount(podcastId int) (err error) {
	stmt, err := s.DB.Prepare("Update podcasts  SET referenced_count = referenced_count + 1 WHERE id = ?")
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
