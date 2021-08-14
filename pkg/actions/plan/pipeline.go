package plan

import (
	"fmt"
	"github.com/analogj/go-util/utils"
	cleanAction "github.com/analogj/terraflow/pkg/actions/clean"
	initAction "github.com/analogj/terraflow/pkg/actions/init"
	"github.com/analogj/terraflow/pkg/config"
	"github.com/sirupsen/logrus"
	"os"
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

	logger.Infof("plan '%s' component in '%s' environment", configuration.GetString("component"), configuration.GetString("environment"))

	terraformPath, err := configuration.GetComponentFolder()
	if err != nil {
		return err
	}

	tfPlanPath := fmt.Sprintf(".tfplan/%s-%s.tfplan", configuration.GetString("environment"), configuration.GetString("component"))
	os.MkdirAll(".tfplan", os.ModePerm)

	cmdPlan := []string{
		"terraform",
		"plan",
		"-input=false",
		"-refresh=true",
		"-no-color",
		fmt.Sprintf("-out=%s", tfPlanPath),
		fmt.Sprintf("-var-file=config/environments/%s.tfvars", configuration.GetString("environment")),
		fmt.Sprintf("-var-file=config/components/%s.tfvars", configuration.GetString("component")),
	}

	if configuration.IsSet("var") {
		for _, val := range configuration.GetStringSlice("var") {
			cmdPlan = append(cmdPlan, fmt.Sprintf("-var='%s'", val))
		}
	}

	if configuration.IsSet("target") {
		cmdPlan = append(cmdPlan, []string{"-target", configuration.GetString("target")}...)
	}

	cmdPlan = append(cmdPlan, terraformPath)
	logger.Debugf("Cmd: %s", strings.Join(cmdPlan, " "))
	return utils.CmdExec(cmdPlan[0], cmdPlan[1:], "", nil, "--> ")
}
