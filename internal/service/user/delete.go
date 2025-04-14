package user

import (
	"context"
	"fmt"
)

func (s *serv) Delete(ctx context.Context, id int64) (bool, error) {
	_, err := s.userRepository.Delete(ctx, id)
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	return false, nil
}
