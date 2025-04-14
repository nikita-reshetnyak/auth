package user_api

import (
	"context"

	v1 "github.com/nikita-reshetnyak/auth/gen/auth_v1"
	"github.com/nikita-reshetnyak/auth/internal/converter"
)

func (i *Implementaion) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	id, err := i.userService.Create(ctx, req.Name, req.Email, req.Password, req.PasswordConfirm, converter.FromProtoRole(req.Role))
	if err != nil {
		return nil, err
	}
	return &v1.CreateResponse{Id: id}, nil
}
