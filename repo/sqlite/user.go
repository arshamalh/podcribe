package sqlite

import (
	"context"
	"podcribe/entities"
	"podcribe/log"

	"go.uber.org/zap"
)

func (s *Sqlite) AddUser(ctx context.Context, user *entities.User) error {
	if _, err := s.db.NewInsert().Model(user).Exec(ctx); err != nil {
		log.Gl.Error(err.Error(), zap.Int64("chatID", user.ChatID))
		return err
	}

	log.Gl.Info("new user created:", zap.Int64("ID", user.ID), zap.Int64("chatID", user.ChatID))
	return nil
}

func (s *Sqlite) GetUserByID(ctx context.Context, userID uint) (*entities.User, error) {
	user := new(entities.User)
	err := s.db.
		QueryRowContext(ctx, "SELECT id, chat_id, email, phone_number, tf_name, tl_name, balance, created_at, updated_at FROM users WHERE id=?").
		Scan(
			&user.ID, &user.ChatID, &user.Email,
			&user.PhoneNumber, &user.TFName,
			&user.TLName, &user.Balance,
			&user.CreatedAt, &user.UpdatedAt,
		)
	return user, err
}

func (s *Sqlite) GetUserByChatID(ctx context.Context, userChatID int64) (*entities.User, error) {
	user := new(entities.User)
	err := s.db.
		QueryRowContext(ctx, "SELECT id, chat_id, email, phone_number, tf_name, tl_name, balance, created_at, updated_at FROM users WHERE chat_id=?", userChatID).
		Scan(
			&user.ID, &user.ChatID, &user.Email,
			&user.PhoneNumber, &user.TFName,
			&user.TLName, &user.Balance,
			&user.CreatedAt, &user.UpdatedAt,
		)
	return user, err
}

func (s *Sqlite) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user := new(entities.User)
	err := s.db.
		QueryRowContext(ctx, "SELECT id, chat_id, email, phone_number, tf_name, tl_name, balance, created_at, updated_at FROM users WHERE email=?", email).
		Scan(
			&user.ID, &user.ChatID, &user.Email,
			&user.PhoneNumber, &user.TFName,
			&user.TLName, &user.Balance,
			&user.CreatedAt, &user.UpdatedAt,
		)
	return user, err
}

func (s *Sqlite) GetUserByPhone(ctx context.Context, phone string) (*entities.User, error) {
	user := new(entities.User)
	err := s.db.
		QueryRowContext(ctx, "SELECT id, chat_id, email, phone_number, tf_name, tl_name, balance, created_at, updated_at FROM users WHERE phone=?", phone).
		Scan(
			&user.ID, &user.ChatID, &user.Email,
			&user.PhoneNumber, &user.TFName,
			&user.TLName, &user.Balance,
			&user.CreatedAt, &user.UpdatedAt,
		)
	return user, err
}
