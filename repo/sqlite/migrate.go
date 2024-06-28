package sqlite

import (
	"context"
	"podcribe/entities"
	"podcribe/log"
)

func (s Sqlite) Migrate(ctx context.Context) {
	s.CreateTablesIfNotExists(ctx)
}

func (s Sqlite) CreateTablesIfNotExists(ctx context.Context) {
	if _, err := s.db.NewCreateTable().
		Model(new(entities.User)).
		IfNotExists().
		Exec(ctx); err != nil {
		log.Gl.Fatal(err.Error())
	}

	if _, err := s.db.NewCreateTable().
		Model(new(entities.Invoice)).
		IfNotExists().
		Exec(ctx); err != nil {
		log.Gl.Fatal(err.Error())
	}

	if _, err := s.db.NewCreateTable().
		Model(new(entities.Charge)).
		IfNotExists().
		Exec(ctx); err != nil {
		log.Gl.Fatal(err.Error())
	}

	if _, err := s.db.NewCreateTable().
		Model(new(entities.TIRTCharge)).
		IfNotExists().
		Exec(ctx); err != nil {
		log.Gl.Fatal(err.Error())
	}

	if _, err := s.db.NewCreateTable().
		Model(new(entities.TONCharge)).
		IfNotExists().
		Exec(ctx); err != nil {
		log.Gl.Fatal(err.Error())
	}

	if _, err := s.db.NewCreateTable().
		Model(new(entities.Audio)).
		IfNotExists().
		Exec(ctx); err != nil {
		log.Gl.Fatal(err.Error())
	}
}
