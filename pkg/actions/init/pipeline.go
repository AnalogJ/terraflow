package init

import (
	"fmt"
	"github.com/analogj/go-util/utils"
	"github.com/analogj/terraflow/pkg"
	"github.com/analogj/terraflow/pkg/config"
	"github.com/sirupsen/logrus"
	"path"
)

func Start(logger logrus.FieldLogger, configuration config.Interface) error {
	//environmentName := configuration.GetString("environment")
	componentName := configuration.GetString("component")

	logger.Info("init terraform project")

	terraformPath := path.Join("components", componentName)

	if !pkg.FolderExists(terraformPath) {
		return fmt.Errorf("component directory is missing: %s", terraformPath)
	}

	// run the terraform commands necessary.
	cmdInit := []string{
		"terraform", "init",
		"-force-copy",
		"-get-plugins=true",
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
	cmdInit = append(cmdInit, terraformPath)

	return utils.CmdExec(cmdInit[0], cmdInit[1:], "", []string{}, "--> ")

}
