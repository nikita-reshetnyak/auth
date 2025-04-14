package user

import (
	"context"
	"fmt"
)

func (s *serv) Update(ctx context.Context, id int64, name string, email string) (bool, error) {
	_, err := s.userRepository.Update(ctx, id, name, email)
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	return false, nil
}
