package authservices

import (
	"context"
	"fmt"

	"github.com/nikita-reshetnyak/auth/internal/domains/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Auth struct {
	user_i UserI
}
type UserI interface {
	Create(
		ctx context.Context,
		name string,
		email string,
		password string,
		password_confirm string,
		role models.UserRole,
	) (int64, error)
	Get(
		ctx context.Context, id int64) (
		models.User,
		error,
	)
	Update(ctx context.Context, id int64, name string, email string) (*emptypb.Empty, error)
	Delete(ctx context.Context, id int64) (*emptypb.Empty, error)
}

func New(user_i UserI) *Auth {
	return &Auth{user_i: user_i}
}

func (a *Auth) Get(
	ctx context.Context, id int64) (
	models.User,
	error,
) {
	user, err := a.user_i.Get(ctx, id)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (a *Auth) Create(
	ctx context.Context,
	name string,
	email string,
	password string,
	password_confirm string,
	role models.UserRole,
) (int64, error) {
	userId, err := a.user_i.Create(ctx, name, email, password, password_confirm, role)
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}
	return userId, nil
}
func (a *Auth) Update(ctx context.Context, id int64, name string, email string) (*emptypb.Empty, error) {
	_, err := a.user_i.Update(ctx, id, name, email)
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("%w", err)
	}
	return &emptypb.Empty{}, nil
}
func (a *Auth) Delete(ctx context.Context, id int64) (*emptypb.Empty, error) {
	_, err := a.user_i.Delete(ctx, id)
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("%w", err)
	}
	return &emptypb.Empty{}, nil
}
