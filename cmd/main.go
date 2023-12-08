package main

import (
	"log"
	"os"
	"reflect"

	"github.com/urfave/cli/v2"

	"go-template/cmd/app"
)

func main() {
	cmd := &cli.App{
		Name: "Your application name",
		Commands: []*cli.Command{
			app.APIServerCommand(),
			app.MigrationCommand(),
		},
	}
	err := cmd.Run(os.Args)
	if err != nil && !reflect.ValueOf(&err).IsNil() {
		log.Fatal(err)
	}
}
