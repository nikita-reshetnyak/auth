// main is the entry point of the application.
// It initializes and runs the program.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nikita-reshetnyak/auth/internal/app"
	"github.com/nikita-reshetnyak/auth/internal/config"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

const grpcPort = 50051

func main() {
	flag.Parse()
	fmt.Println("hello world")
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	pgConfig, err := config.NewPgConfig()
	if err != nil {
		log.Fatalf("failed to load pg config: %v", err)
	}
	application := app.New(grpcPort, pgConfig.DSN())

	application.GRPCServer.Run()

}
