package display

import (
	"fmt"

	"github.com/atarte/cli-cd/data"
)

func AppListScreen(appdata *data.AppData) {
	titleScreen()

	fmt.Println("App list:")
	for i, monitoredApp := range appdata.MonitoredAppList {
		fmt.Printf("(%d) - %s | path: %s\n", i, monitoredApp.Name, monitoredApp.Path)
	}

	returnScreen()
}
