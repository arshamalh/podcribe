package sqlite_test

import (
	"context"
	"fmt"
	"os"
	"podcribe/entities"
	"podcribe/repo"
	"podcribe/repo/sqlite"
	"testing"
)

var GlobalDB repo.DB

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	code := m.Run()

	if err := teardown(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(code)
}

func setup() error {
	fmt.Println("setting up")
	db, err := sqlite.New("data.db")
	if err != nil {
		return err
	}
	ctx := context.Background()

	// Migrate
	db.Migrate(ctx)

	// Seed
	if err := db.AddUser(ctx, &entities.User{
		ChatID: 10,
	}); err != nil {
		return err
	}

	GlobalDB = db
	return nil
}

func teardown() error {
	fmt.Println("tearing down")
	return os.Remove("data.db")
}
