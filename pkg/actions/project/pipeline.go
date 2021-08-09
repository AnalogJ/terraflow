package project

import (
	"fmt"
	"os"
)

func Start(componentName string, environmentName string) error {
	err := os.MkdirAll("config/environments", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll("config/components", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll("components", os.ModePerm)
	if err != nil {
		return err
	}

	if len(environmentName) > 0 {
		_, err = os.OpenFile(fmt.Sprintf("config/environments/%s.tfvars", environmentName), os.O_RDONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if len(componentName) > 0 {
		_, err = os.OpenFile(fmt.Sprintf("config/components/%s.tfvars", componentName), os.O_RDONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		err = os.MkdirAll(fmt.Sprintf("components/%s", componentName), os.ModePerm)
		if err != nil {
			return err
		}
		_, err = os.OpenFile(fmt.Sprintf("components/%s/main.tf", componentName), os.O_RDONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		_, err = os.OpenFile(fmt.Sprintf("components/%s/output.tf", componentName), os.O_RDONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		_, err = os.OpenFile(fmt.Sprintf("components/%s/secrets.tf", componentName), os.O_RDONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		_, err = os.OpenFile(fmt.Sprintf("components/%s/security.tf", componentName), os.O_RDONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		_, err = os.OpenFile(fmt.Sprintf("components/%s/variables.tf", componentName), os.O_RDONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		_, err = os.OpenFile(fmt.Sprintf("components/%s/provider.tf", componentName), os.O_RDONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		_, err = os.OpenFile(fmt.Sprintf("components/%s/backend.tf", componentName), os.O_RDONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
