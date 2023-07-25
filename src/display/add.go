package display

import (
	"fmt"
	"strings"

	"github.com/atarte/cli-cd/data"
	"github.com/google/uuid"
)

// addScreen
func addScreen(appdata *data.AppData) {
	titleScreen()

	fmt.Println("Add a application to be monitored.")

	// App path
	fmt.Println("Application absolute path:")
	var appPath string
	fmt.Scanln(&appPath)

	splitAppPath := strings.Split(appPath, "/")
	appNameDefault := splitAppPath[len(splitAppPath)-1]

	// App name
	fmt.Printf("Application name (default: '%s'):\n", appNameDefault)
	var appName string
	fmt.Scanln(&appPath)

	if appName == "" {
		appName = appNameDefault
	}

	// fmt.Println("Do you want to activate automatic deploy (y/n)")
	// var automaticDeploy string
	// fmt.Scanln(&automaticDeploy)
	// AutomaticDeployBool := true
	// if automaticDeploy != "y" {
	// 	AutomaticDeployBool = false
	// }

	// fmt.Println("Refresh frenquence is set at 10s by default")
	// // var refreshFrequence string
	// // if automaticDeploy == "y" {
	// // 	fmt.Println("")
	// // } else {
	// //
	// // }

	// fmt.Println("Dockerfile relative path (default: \"./Dockerfile\")")
	// var dockerfilePath string
	// fmt.Scanln(&dockerfilePath)
	// if dockerfilePath == "" {
	// 	dockerfilePath = "./Dockerfile"
	// }

	// fmt.Println("Do you want to support Kubernetes (y/n) (for now no by default)")

	var id string = uuid.NewString()

	newMonitoredApp := data.MonitoredApp{
		Uuid: id,
		Name: appName,
		Path: appPath,
		// AutomaticDeploy:  AutomaticDeployBool,
		// RefreshFrequence: 1000,
		// DockerfilePath:   dockerfilePath,
		// Kubernetes:       false,
		// Checksum:         "",
	}

	appdata.AddMonitoredApp(newMonitoredApp)

	returnScreen()
}
