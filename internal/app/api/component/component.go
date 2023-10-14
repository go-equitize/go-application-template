package component

import "go-template/internal/pkg/util/logger"

func InitComponents() error {
	var err error

	err = logger.InitLogger()
	if err != nil {
		return err
	}

	return err
}