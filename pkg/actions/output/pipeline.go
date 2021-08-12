package output

import (
	"github.com/analogj/go-util/utils"
	"github.com/analogj/terraflow/pkg/config"
	"github.com/sirupsen/logrus"
)

func Start(logger logrus.FieldLogger, configuration config.Interface) error {

	logger.Info("outputs for local terraform.tfstate file")

	// run the terraform commands necessary.
	cmdInit := []string{
		"terraform", "output",
		"-no-color",
		"-json",
	}

	return utils.CmdExec(cmdInit[0], cmdInit[1:], "", []string{}, "--> ")

}
