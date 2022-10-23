package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

//Clean local state (backend) files
func CleanState() {

}

//List available components
func ListComponents() {

}

func ListEnvironments() {

}

// Check if a folder exists
func FolderExists(pathFolder string) bool {
	if pathAbs, err := filepath.Abs(pathFolder); err != nil {
		return false
	} else if fileInfo, err := os.Stat(pathAbs); os.IsNotExist(err) || !fileInfo.IsDir() {
		return false
	}

	return true
}

func TFVarFiles(envName string, compName string) ([]string, error) {
	envTFvars, err := filepath.Abs(fmt.Sprintf("config/environments/%s.tfvars", envName))
	if err != nil {
		return nil, err
	}
	compTfvars, err := filepath.Abs(fmt.Sprintf("config/components/%s.tfvars", compName))
	if err != nil {
		return nil, err
	}

	return []string{
		fmt.Sprintf("-var-file=%s", envTFvars),
		fmt.Sprintf("-var-file=%s", compTfvars),
	}, nil
}
