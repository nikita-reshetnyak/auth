package repository

import (
	"context"

	"github.com/nikita-reshetnyak/auth/internal/model"
)

type UserRepository interface {
	Create(
		ctx context.Context,
		name string,
		email string,
		password string,
		password_confirm string,
		role model.Role,
	) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, id int64, name string, email string) (bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
