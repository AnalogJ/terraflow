package validate

import (
	"fmt"
	"github.com/analogj/go-util/utils"
	"github.com/analogj/terraflow/pkg/config"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func Start(logger logrus.FieldLogger, configuration config.Interface) error {
	baseDir := "components"
	componentDirectories, err := ioutil.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range componentDirectories {
		if !file.IsDir() {
			continue
		}

		log.Printf("validating component %s:", file.Name())
		cmdValidate := []string{
			"terraform",
			fmt.Sprintf("-chdir=%s", filepath.Join(baseDir, file.Name())),
			"validate",
			"-no-color",
		}

		logger.Infof("Terraform Cmd: %s", strings.Join(cmdValidate, " "))
		err := utils.CmdExec(cmdValidate[0], cmdValidate[1:], "", nil, "--> ")
		if err != nil {
			return fmt.Errorf("an error occurred validating %s: %w", file.Name(), err)
		}
	}
	return nil
}
