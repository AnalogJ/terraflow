package clean

import (
	"github.com/analogj/terraflow/pkg/config"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func Start(logger logrus.FieldLogger, configuration config.Interface) error {
	terraformPath, err := configuration.GetComponentFolder()
	if err != nil {
		return err
	}
	dotTerraformPath := filepath.Join(terraformPath, ".terraform")

	err = os.RemoveAll(dotTerraformPath)
	if err != nil {
		return err
	}
	return nil
}
