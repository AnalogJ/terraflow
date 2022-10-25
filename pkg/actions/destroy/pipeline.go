package destroy

import (
	"fmt"
	"github.com/analogj/go-util/utils"
	"github.com/analogj/terraflow/pkg"
	cleanAction "github.com/analogj/terraflow/pkg/actions/clean"
	initAction "github.com/analogj/terraflow/pkg/actions/init"
	"github.com/analogj/terraflow/pkg/config"
	"github.com/sirupsen/logrus"
	"strings"
)

func Start(logger logrus.FieldLogger, configuration config.Interface) error {
	err := cleanAction.Start(logger, configuration)
	if err != nil {
		return err
	}

	err = initAction.Start(logger, configuration)
	if err != nil {
		return err
	}

	logger.Infof("destroy '%s' component in '%s' environment", configuration.GetString("component"), configuration.GetString("environment"))

	terraformPath, err := configuration.GetComponentFolder()
	if err != nil {
		return err
	}

	cmdDestroy := []string{
		"terraform",
		fmt.Sprintf("-chdir=%s", terraformPath),
		"destroy",
		"-no-color",
		"-auto-approve",
	}

	tfVars, err := pkg.TFVarFiles(configuration.GetString("environment"), configuration.GetString("component"))
	if err != nil {
		return err
	}
	cmdDestroy = append(cmdDestroy, tfVars...)

	if configuration.IsSet("var") {
		for _, val := range configuration.GetStringSlice("var") {
			cmdDestroy = append(cmdDestroy, fmt.Sprintf("-var='%s'", val))
		}
	}

	if configuration.IsSet("target") {
		cmdDestroy = append(cmdDestroy, []string{"-target", configuration.GetString("target")}...)
	}

	logger.Infof("Terraform Cmd: %s", strings.Join(cmdDestroy, " "))
	return utils.CmdExec(cmdDestroy[0], cmdDestroy[1:], "", nil, "--> ")
}
