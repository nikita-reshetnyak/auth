package user_api

import (
	v1 "github.com/nikita-reshetnyak/auth/gen/auth_v1"
	"github.com/nikita-reshetnyak/auth/internal/service"
)

type Implementaion struct {
	v1.UnimplementedAuthV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementaion {
	return &Implementaion{
		userService: userService,
	}
}
