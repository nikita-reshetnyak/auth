// main is the entry point of the application.
// It initializes and runs the program.
package main

import (
	"fmt"

	"github.com/nikita-reshetnyak/auth/internal/app"
)

const grpcPort = 50051

func main() {
	fmt.Println("hello world")
	application := app.New(grpcPort)

	application.GRPCServer.Run()

}
