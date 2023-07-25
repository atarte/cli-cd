package display

import (
	"fmt"
	"os"

	"github.com/atarte/cli-cd/data"
)

// MenuScreen
func MenuScreen(appdata *data.AppData) {
	titleScreen()

	fmt.Println("(1) - View app list")
	fmt.Println("(2) - Add app")
	fmt.Println("(3) - Update app")
	fmt.Println("(4) - Delete app")
	fmt.Println("(5) - Info")
	fmt.Println("(6) - Quit")

	var input string
	fmt.Scanln(&input)

	switch input {
	case "1":
		AppListScreen(appdata)
	case "2":
		addScreen(appdata)
	case "3":
		// UpdateApp()
	case "4":
		// DeleteApp()
	case "5":
		infoScreen()
	case "6":
		os.Exit(0)
	default:
		// no valid inout detected
		return
	}
}
