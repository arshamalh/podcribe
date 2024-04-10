package repo

import (
	"context"
	"podcribe/entities"
)

type DB interface {
	User
}

type User interface {
	AddUser(ctx context.Context, user *entities.User) error
	GetUserByID(ctx context.Context, userID uint) (*entities.User, error)
	GetUserByChatID(ctx context.Context, userChatID int64) (*entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*entities.User, error)
}
