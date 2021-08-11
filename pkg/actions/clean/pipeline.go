package clean

import (
	"github.com/analogj/terraflow/pkg/config"
	"github.com/sirupsen/logrus"
	"os"
)

func Start(logger logrus.FieldLogger, configuration config.Interface) error {

	err := os.RemoveAll(".terraform")
	if err != nil {
		return err
	}
	return nil
}
