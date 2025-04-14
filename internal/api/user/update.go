package user_api

import (
	"context"

	v1 "github.com/nikita-reshetnyak/auth/gen/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementaion) Update(ctx context.Context, req *v1.UpdateRequest) (*emptypb.Empty, error) {
	_, err := i.userService.Update(ctx, req.Id, req.Name.String(), req.Email.String())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
