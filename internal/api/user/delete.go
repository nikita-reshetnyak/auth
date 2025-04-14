package user_api

import (
	"context"

	v1 "github.com/nikita-reshetnyak/auth/gen/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementaion) Delete(ctx context.Context, req *v1.DeleteRequest) (*emptypb.Empty, error) {
	_, err := i.userService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
