package config

import (
	"fmt"
	"github.com/analogj/terraflow/pkg"
	"github.com/spf13/viper"
	"os"
	"path"
	"strings"
)

// When initializing this class the following methods must be called:
// Config.New
// Config.Init
// This is done automatically when created via the Factory.

type configuration struct {
	*viper.Viper
}

func New() Interface {
	newConfig := new(configuration)
	newConfig.Init()
	return newConfig
}

func (c *configuration) Init() {
	c.Viper = viper.New()
	c.SetEnvPrefix("TF")
	c.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	c.AutomaticEnv()
}

func (c *configuration) BackendConfig() (map[string]string, error) {
	backendConfig := map[string]string{}

	for _, element := range os.Environ() {
		variable := strings.Split(element, "=")
		fmt.Println(variable[0], "=>", variable[1])
	}

	for _, envVarPair := range os.Environ() {
		envVarParts := strings.SplitN(envVarPair, "=", 2)
		if strings.HasPrefix(strings.ToLower(envVarParts[0]), "backend_config") && len(envVarParts) >= 2 {
			backendConfigKey := strings.TrimPrefix(strings.ToLower(envVarParts[0]), "backend_config_")
			backendConfig[backendConfigKey] = envVarParts[1]
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
