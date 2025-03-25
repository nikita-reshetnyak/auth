package authgrpc

import (
	"context"

	v1 "github.com/nikita-reshetnyak/auth/gen/auth_v1"
	models "github.com/nikita-reshetnyak/auth/internal/domains/models"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Auth interface {
	Create(
		ctx context.Context,
		name string,
		email string,
		password string,
		password_confirm string,
		role models.UserRole,
	) (int64, error)
	Get(context.Context, int64) (models.User, error)
	Update(ctx context.Context, id int64, name string, email string) (*emptypb.Empty, error)
	Delete(ctx context.Context, id int64) (*emptypb.Empty, error)
}
type serverApi struct {
	v1.UnimplementedAuthV1Server
	auth Auth
}

func Register(grpcServer *grpc.Server, auth Auth) {
	v1.RegisterAuthV1Server(grpcServer, &serverApi{auth: auth})
}
func (s *serverApi) Get(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	user, err := s.auth.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      v1.Role(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
func (s *serverApi) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	id, err := s.auth.Create(ctx, req.Name, req.Email, req.Password, req.PasswordConfirm, models.UserRole(req.Role))
	if err != nil {
		return nil, err
	}
	return &v1.CreateResponse{Id: id}, nil
}
func (s *serverApi) Update(ctx context.Context, req *v1.UpdateRequest) (*emptypb.Empty, error) {
	empty, err := s.auth.Update(ctx, req.Id, req.Name.String(), req.Email.String())
	if err != nil {
		return nil, err
	}
	return empty, nil
}
func (s *serverApi) Delete(ctx context.Context, req *v1.DeleteRequest) (*emptypb.Empty, error) {
	empty, err := s.auth.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return empty, nil
}
