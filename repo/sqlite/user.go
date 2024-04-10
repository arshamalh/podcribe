package sqlite

import (
	"context"
	"podcribe/entities"
	"podcribe/log"

	"go.uber.org/zap"
)

func (s *sqlite) AddUser(ctx context.Context, user *entities.User) error {
	if _, err := s.db.NewInsert().Model(user).Exec(ctx); err != nil {
		log.Gl.Error(err.Error(), zap.Int64("chatID", user.ChatID))
		return err
	}

	log.Gl.Info("new user created:", zap.Uint("ID", user.ID), zap.Int64("chatID", user.ChatID))
	return nil
}

func (s *sqlite) GetUserByID(ctx context.Context, userID uint) (*entities.User, error) {
	user := new(entities.User)
	err := s.db.
		QueryRowContext(ctx, "SELECT * FROM users WHERE id=?").
		Scan(
			&user.ID, &user.ChatID, &user.Password,
			&user.Email, &user.PhoneNumber,
			&user.CreatedAt, &user.UpdatedAt,
		)
	return user, err
}

func (s *sqlite) GetUserByChatID(ctx context.Context, userChatID int64) (*entities.User, error) {
	user := new(entities.User)
	err := s.db.
		QueryRowContext(ctx, "SELECT * FROM users WHERE chat_id=?", userChatID).
		Scan(
			&user.ID, &user.ChatID, &user.Password,
			&user.Email, &user.PhoneNumber,
			&user.CreatedAt, &user.UpdatedAt,
		)
	return user, err
}

func (s *sqlite) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user := new(entities.User)
	err := s.db.
		QueryRowContext(ctx, "SELECT * FROM users WHERE email=?", email).
		Scan(
			&user.ID, &user.ChatID, &user.Password,
			&user.Email, &user.PhoneNumber,
			&user.CreatedAt, &user.UpdatedAt,
		)
	return user, err
}

func (s *sqlite) GetUserByPhone(ctx context.Context, phone string) (*entities.User, error) {
	user := new(entities.User)
	err := s.db.
		QueryRowContext(ctx, "SELECT * FROM users WHERE phone=?", phone).
		Scan(
			&user.ID, &user.ChatID, &user.Password,
			&user.Email, &user.PhoneNumber,
			&user.CreatedAt, &user.UpdatedAt,
		)
	return user, err
}
