package config

import (
	"bytes"
	"fmt"
	"github.com/analogj/terraflow/pkg"
	"github.com/spf13/viper"
	"path"
	"strings"
	"text/template"
)

// When initializing this class the following methods must be called:
// Config.New
// Config.Init
// This is done automatically when created via the Factory.

func New() Interface {
	newConfig := new(configuration)
	newConfig.Init()
	return newConfig
}

type configuration struct {
	*viper.Viper
}

func (c *configuration) Init() {
	c.Viper = viper.New()
}

func (c *configuration) BackendConfig() (map[string]string, error) {
	backendConfig := map[string]string{}

	for key, val := range c.AllSettings() {
		if strings.HasPrefix(strings.ToLower(key), "backend_config") {
			backendConfigKey := strings.TrimPrefix(strings.ToLower(key), "backend_config")
			tmpl, err := template.New(fmt.Sprintf("templ_%s", key)).Parse(val.(string))
			if err != nil {
				return nil, err
			}

			var populatedString bytes.Buffer
			err = tmpl.Execute(&populatedString, c.AllSettings())
			if err != nil {
				return nil, err
			}

			backendConfig[backendConfigKey] = populatedString.String()
		}
	}
	return backendConfig, nil
}

func (c *configuration) GetComponentFolder() (string, error) {
	if !c.IsSet("component") {
		return "", fmt.Errorf("no component set")
	}

	terraformPath := path.Join("components", c.GetString("component"))

	if !pkg.FolderExists(terraformPath) {
		return "", fmt.Errorf("component directory is missing: %s", terraformPath)
	}
	return terraformPath, nil
}
