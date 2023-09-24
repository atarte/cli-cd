package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

const CLI_CD_DIR string = ".cli-cd"

// getAppDataDirPath
func getCliCdDirPath() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Join(currentUser.HomeDir, CLI_CD_DIR), nil
}

// isCliCdDirExist
func isCliCdDirExist() bool {
	dirPath, err := getCliCdDirPath()
	if err != nil {
		return false
	}

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateCliCdDir
func CreateCliCdDir() error {
	if !isCliCdDirExist() {
		dirPath, err := getCliCdDirPath()
		if err != nil {
			return fmt.Errorf("Cannot found the main directory path, `%s`, cause by : `%s`", dirPath, err)
		}

		if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
			return fmt.Errorf("Cannot create the main directory: %s", err)
		}
	}

	return nil
}
