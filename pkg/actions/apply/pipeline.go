package apply

import (
	"fmt"
	"github.com/analogj/go-util/utils"
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

	logger.Infof("apply '%s' component in '%s' environment", configuration.GetString("component"), configuration.GetString("environment"))

	terraformPath, err := configuration.GetComponentFolder()
	if err != nil {
		return err
	}

	cmdApply := []string{
		"terraform", "apply",
		"-input=false",
		"-no-color",
		fmt.Sprintf("-var-file=config/environments/%s.tfvars", configuration.GetString("environment")),
		fmt.Sprintf("-var-file=config/components/%s.tfvars", configuration.GetString("component")),
	}

	if configuration.IsSet("var") {
		for _, val := range configuration.GetStringSlice("var") {
			cmdApply = append(cmdApply, fmt.Sprintf("-var='%s'", val))
		}
	}

	if configuration.IsSet("target") {
		cmdApply = append(cmdApply, []string{"-target", configuration.GetString("target")}...)
	}

	cmdApply = append(cmdApply, terraformPath)

	logger.Infof("Terraform Cmd: %s", strings.Join(cmdApply, " "))
	return utils.CmdExec(cmdApply[0], cmdApply[1:], "", nil, "--> ")
}
