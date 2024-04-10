package sqlite

import (
	"context"
	"podcribe/entities"
	"podcribe/log"
)

func (s sqlite) Migrate(ctx context.Context) {
	s.CreateTablesIfNotExists(ctx)
}

func (s sqlite) CreateTablesIfNotExists(ctx context.Context) {
	if _, err := s.db.NewCreateTable().
		Model(new(entities.User)).
		IfNotExists().
		Exec(ctx); err != nil {
		log.Gl.Fatal(err.Error())
	}
}
