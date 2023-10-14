package app

import (
	"github.com/urfave/cli/v2"

	"go-template/internal/app/api/component"
	"go-template/internal/app/api/server"
	"go-template/internal/pkg/config"
)

func APIServerCommand() *cli.Command {
	return &cli.Command{
		Name:    "api",
		Aliases: []string{},
		Usage:   "Run api server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "internal/pkg/config/file/default.yaml",
				Usage:   "Configuration file",
			},
		},
		Action: func(c *cli.Context) error {
			configFilePath := c.String("config")

			err := config.Load(configFilePath)
			if err != nil {
				return err
			}

			err = component.InitComponents()
			if err != nil {
				return err
			}

			s := server.NewAPIServer()
			return s.Run()
		},
	}
}
