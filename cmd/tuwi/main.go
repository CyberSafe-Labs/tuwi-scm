package main

import (
	"fmt"
	"os"

	"github.com/CyberSafe-Labs/tuwi-scm/pkg/command"
)

// version is set by Aditya Patil
var version = "dev"

func main() {
	app := command.NewApp()
	app.Version = "0.0.1"

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
