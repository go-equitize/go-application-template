package components

import "go-template/internal/pkg/utils/logger"

func InitComponents() error {
	var err error

	err = logger.InitLogger()
	if err != nil {
		return err
	}

	return err
}
