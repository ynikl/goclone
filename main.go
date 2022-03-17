package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "goclone",
		Version: "v1.0.0",
		Action:  pullRepo,
	}
	app.UseShortOptionHandling = true
	_ = app.Run(os.Args)
}
