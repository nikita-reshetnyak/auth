package grpcapp

import (
	"fmt"
	"log"
	"net"

	authgrpc "github.com/nikita-reshetnyak/auth/internal/grpc"
	"google.golang.org/grpc"
)

type App struct {
	grpcServer *grpc.Server
	port       int
}

func New(auth authgrpc.Auth, port int) *App {
	server := grpc.NewServer()
	authgrpc.Register(server, auth)
	return &App{
		grpcServer: server,
		port:       port,
	}
}
func (a *App) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := a.grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
