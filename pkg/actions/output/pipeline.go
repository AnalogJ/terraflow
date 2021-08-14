package output

import (
	"github.com/analogj/go-util/utils"
	"github.com/analogj/terraflow/pkg/config"
	"github.com/sirupsen/logrus"
)

func Start(logger logrus.FieldLogger, configuration config.Interface) error {

	logger.Info("outputs for local terraform.tfstate file")

	// run the terraform commands necessary.
	cmdOutput := []string{
		"terraform", "output",
		"-no-color",
		"-json",
	}

	return utils.CmdExec(cmdOutput[0], cmdOutput[1:], "", nil, "--> ")

}
