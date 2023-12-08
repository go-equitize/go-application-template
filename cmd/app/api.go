package app

import (
	"log"

	"github.com/urfave/cli/v2"

	"go-template/internal/app/api/component"
	"go-template/internal/app/api/server"
	"go-template/internal/pkg/config"
	"go-template/internal/pkg/migration"
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
			&cli.BoolFlag{
				Name:    "auto_migration",
				Aliases: []string{"a"},
				Value:   false,
			},
			&cli.StringFlag{
				Name:    "migration_dir",
				Aliases: []string{"m"},
				Value:   "file://./migrations/mysql",
			},
		},
		Action: func(c *cli.Context) error {
			configFilePath := c.String("config")

			err := config.Load(configFilePath)
			if err != nil {
				return err
			}

			// auto migrations
			if c.Bool("auto_migration") {
				log.Println("-------- Run migration --------")
				m, err := migration.NewMigration(c.String("migration_dir"))
				if err != nil {
					log.Println("Can not create migration " + err.Error())
					return err
				}
				err = m.MigrateUp(0)
				if err != nil && err.Error() != "no change" {
					return err
				}
				sourceErr, dbErr := m.Close()
				if sourceErr != nil {
					return sourceErr
				}
				if dbErr != nil {
					return dbErr
				}
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
