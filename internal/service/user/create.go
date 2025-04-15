package user

import (
	"context"
	"fmt"

	"github.com/nikita-reshetnyak/auth/internal/model"
)

func (s *serv) Create(
	ctx context.Context,
	name string,
	email string,
	password string,
	password_confirm string,
	role model.Role,
) (int64, error) {
	userId, err := s.userRepository.Create(ctx, name, email, password, password_confirm, role)
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}
	return userId, nil
}
