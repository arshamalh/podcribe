package cmd

import (
	"context"
	"podcribe/log"
	"podcribe/repo/sqlite"

	"github.com/spf13/cobra"
)

func registerMigrate(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "migrate database schema",
		Run: func(cmd *cobra.Command, args []string) {
			log.Initialize()
			migrate()
		},
	}

	root.AddCommand(cmd)
}

func migrate() {
	db, err := sqlite.New("./data.db")
	if err != nil {
		log.Gl.Fatal("unable to initialize db")
	}
	db.Migrate(context.TODO())
}
