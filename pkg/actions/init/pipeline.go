package init

import (
	"fmt"
	"github.com/analogj/go-util/utils"
	"github.com/analogj/terraflow/pkg/config"
	"github.com/sirupsen/logrus"
	"strings"
)

func Start(logger logrus.FieldLogger, configuration config.Interface) error {
	logger.Info("init terraform project")

	terraformPath, err := configuration.GetComponentFolder()
	if err != nil {
		return err
	}

	// run the terraform commands necessary.
	cmdInit := []string{
		"terraform",
		fmt.Sprintf("-chdir=%s", terraformPath),
		"init",
		"-force-copy",
		"-input=false",
	}

	backendConfig, err := configuration.BackendConfig()
	if err != nil {
		return err
	}
	if len(backendConfig) > 0 {
		logger.Infof("Backend Configured: %v", backendConfig)
		for key, val := range backendConfig {
			cmdInit = append(cmdInit, "-backend=true")
			cmdInit = append(cmdInit, fmt.Sprintf("-backend-config=\"%s=%s\"", key, val))
		}
	}

	logger.Infof("Terraform Cmd: %s", strings.Join(cmdInit, " "))
	return utils.CmdExec(cmdInit[0], cmdInit[1:], "", nil, "--> ")

}
