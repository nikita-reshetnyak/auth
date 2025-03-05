package app

import (
	grpcapp "github.com/nikita-reshetnyak/auth/internal/app/grpc"
	authservices "github.com/nikita-reshetnyak/auth/internal/services/auth"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(grpcport int) *App {
	authService := authservices.New()
	grpcApp := grpcapp.New(authService, grpcport)
	return &App{GRPCServer: grpcApp}
}
