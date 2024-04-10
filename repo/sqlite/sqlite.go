package sqlite

// If the link was the same as database, don't download that again, using already downloaded files
// Keep track of every downloaded files

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	bunDebug "github.com/uptrace/bun/extra/bundebug"
)

type sqlite struct {
	db *bun.DB
}

func New(path string) (*sqlite, error) {
	sqlDB, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	db := bun.NewDB(sqlDB, sqlitedialect.New())
	db.AddQueryHook(bunDebug.NewQueryHook(
		bunDebug.WithVerbose(true),
		bunDebug.FromEnv("DEBUG"),
	))

	return &sqlite{
		db: db,
	}, nil
}
