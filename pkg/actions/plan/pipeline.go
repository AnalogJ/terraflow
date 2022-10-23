package plan

import (
	"fmt"
	"github.com/analogj/go-util/utils"
	"github.com/analogj/terraflow/pkg"
	cleanAction "github.com/analogj/terraflow/pkg/actions/clean"
	initAction "github.com/analogj/terraflow/pkg/actions/init"
	"github.com/analogj/terraflow/pkg/config"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
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
	tfPlanPath, err := filepath.Abs(fmt.Sprintf(".tfplan/%s-%s.tfplan", configuration.GetString("environment"), configuration.GetString("component")))
	if err != nil {
		return err
	}
	err = os.MkdirAll(".tfplan", os.ModePerm)
	if err != nil {
		return err
	}

	cmdPlan := []string{
		"terraform",
		fmt.Sprintf("-chdir=%s", terraformPath),
		"plan",
		"-input=false",
		"-refresh=true",
		"-no-color",
		fmt.Sprintf("-out=%s", tfPlanPath),
	}

	tfVars, err := pkg.TFVarFiles(configuration.GetString("environment"), configuration.GetString("component"))
	if err != nil {
		return err
	}
	cmdPlan = append(cmdPlan, tfVars...)

	if configuration.IsSet("var") {
		for _, val := range configuration.GetStringSlice("var") {
			cmdPlan = append(cmdPlan, fmt.Sprintf("-var='%s'", val))
		}
	}

	if configuration.IsSet("target") {
		cmdPlan = append(cmdPlan, []string{"-target", configuration.GetString("target")}...)
	}

	logger.Infof("Terraform Cmd: %s", strings.Join(cmdPlan, " "))
	return utils.CmdExec(cmdPlan[0], cmdPlan[1:], "", nil, "--> ")
}
