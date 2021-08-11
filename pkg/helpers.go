package pkg

import (
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
