package user_api

import (
	"context"

	v1 "github.com/nikita-reshetnyak/auth/gen/auth_v1"
	"github.com/nikita-reshetnyak/auth/internal/converter"
)

func (i *Implementaion) Get(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return converter.ToUserFromService(user), nil
}
