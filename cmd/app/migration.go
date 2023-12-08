package app

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"

	"go-template/internal/pkg/config"
	"go-template/internal/pkg/migration"
)

func MigrationCommand() *cli.Command {
	return &cli.Command{
		Name:    "migration",
		Aliases: []string{},
		Usage:   "Run migration",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "internal/pkg/config/file/default.yaml",
				Usage:   "Configuration file",
			},
			&cli.StringFlag{
				Name:  "migration_dir",
				Value: "file://./migrations/mysql",
			},
			&cli.BoolFlag{
				Name:  "down",
				Value: false,
			},
		},
		Action: func(c *cli.Context) error {
			err := config.Load(c.String("config"))
			if err != nil {
				return err
			}

			down := c.Bool("down")

			m, err := migration.NewMigration(c.String("migration_dir"))
			if err != nil {
				log.Println("Can not create migration " + err.Error())
				return err
			}

			if down {
				err = m.MigrateDown(0)
			} else {
				err = m.MigrateUp(0)
			}

			if err != nil && err.Error() != "no change" {
				fmt.Println(err)
				return err
			}
			sourceErr, dbErr := m.Close()
			if sourceErr != nil {
				return sourceErr
			}
			if dbErr != nil {
				return dbErr
			}
			return nil
		},
	}
}
