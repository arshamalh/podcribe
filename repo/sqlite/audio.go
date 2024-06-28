package sqlite

import (
	"context"
	"podcribe/entities"
	"podcribe/log"

	"go.uber.org/zap"
)

func (s *Sqlite) AddAudio(ctx context.Context, audio *entities.Audio) error {
	if _, err := s.db.NewInsert().Model(audio).Exec(ctx); err != nil {
		log.Gl.Error(err.Error(), zap.Int64("audio.UserID", audio.UserID))
		return err
	}

	log.Gl.Info("new audio created:", zap.Int64("userID", audio.UserID), zap.Int64("audioID", audio.ID))
	return nil
}

func (s *Sqlite) UpdateAudio(ctx context.Context, audio *entities.Audio) error {
	_, err := s.db.NewUpdate().
		Model(audio).
		OmitZero().
		WherePK().
		Exec(ctx)
	return err
}
