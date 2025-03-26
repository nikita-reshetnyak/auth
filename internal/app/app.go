package app

import (
	grpcapp "github.com/nikita-reshetnyak/auth/internal/app/grpc"
	authservices "github.com/nikita-reshetnyak/auth/internal/services/auth"
	postgres_strg "github.com/nikita-reshetnyak/auth/internal/storage/postgres"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(grpcport string, configPath string) *App {
	pool, err := postgres_strg.New(configPath)
	if err != nil {
		panic(err)
	}
	authService := authservices.New(pool)
	grpcApp := grpcapp.New(authService, grpcport)
	return &App{GRPCServer: grpcApp}
}
