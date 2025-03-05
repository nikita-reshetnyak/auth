package authservices

import (
	"context"
	"time"

	"github.com/nikita-reshetnyak/auth/internal/domains/models"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Auth struct{}

func New() *Auth {
	return &Auth{}
}

func (a *Auth) Get(
	ctx context.Context, id int64) (
	int64,
	string,
	string,
	models.UserRole,
	*timestamppb.Timestamp,
	*timestamppb.Timestamp,
	error,
) {

	return id,
		"admin",
		"admin@admin",
		models.RoleAdmin,
		timestamppb.New(time.Now()),
		timestamppb.New(time.Now()),
		nil
}
func (a *Auth) Create(
	ctx context.Context,
	name string,
	email string,
	password string,
	password_confirm string,
	role models.UserRole,
) (int64, error) {
	return 1, nil
}
func (a *Auth) Update(ctx context.Context, id int64, name string, email string) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (a *Auth) Delete(ctx context.Context, id int64) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
