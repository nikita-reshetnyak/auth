package user

import (
	"github.com/nikita-reshetnyak/auth/internal/client/db"
	"github.com/nikita-reshetnyak/auth/internal/repository"
	"github.com/nikita-reshetnyak/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
) service.UserService {
	return &serv{userRepository: userRepository, txManager: txManager}
}
