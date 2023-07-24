package display

import (
	"fmt"

	"github.com/atarte/cli-cd/config"
)

func infoScreen() {
	titleScreen()

	fmt.Printf("App name: %s\n", config.AppName)
	fmt.Printf("Version: %s\n", config.Version)
	fmt.Printf("Copyright: ©2023-%s\n", config.ReleaseYear)

	returnScreen()
}
