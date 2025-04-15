package app

import (
	"context"
	"log"

	user_api "github.com/nikita-reshetnyak/auth/internal/api/user"
	"github.com/nikita-reshetnyak/auth/internal/client/db"
	"github.com/nikita-reshetnyak/auth/internal/client/db/pg"
	"github.com/nikita-reshetnyak/auth/internal/client/db/transaction"
	"github.com/nikita-reshetnyak/auth/internal/closer"
	"github.com/nikita-reshetnyak/auth/internal/config"
	"github.com/nikita-reshetnyak/auth/internal/repository"
	user_repo "github.com/nikita-reshetnyak/auth/internal/repository/user"
	"github.com/nikita-reshetnyak/auth/internal/service"
	userService "github.com/nikita-reshetnyak/auth/internal/service/user"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GrpcConfig

	dbClient       db.Client
	txManager      db.TxManager
	userRepository repository.UserRepository

	userService service.UserService

	userImpl *user_api.Implementaion // why with *
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPgConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GrpcConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}
func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = user_repo.NewRepository(s.DBClient(ctx))
	}
	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}
func (s *serviceProvider) NoteImpl(ctx context.Context) *user_api.Implementaion {
	if s.userImpl == nil {
		s.userImpl = user_api.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}
