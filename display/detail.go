package display

import (
	"fmt"
	"strings"

	"github.com/atarte/cli-cd/data"
)

// detailScreen
func detailScreen(appdata *data.AppData, index int) {
	monitoredApp := appdata.MonitoredAppList[index]

	for {
		titleScreen()

		fmt.Println()
		fmt.Println("Some detail / stat/ info on the app")
		fmt.Println("info error if on ocure")
		fmt.Printf("%#v\n", monitoredApp)

		fmt.Println()
		fmt.Println("(u) - up")
		fmt.Println("(d) - down")
		fmt.Println("(r) - restart")
		fmt.Println("(l) - log")
		fmt.Println("(x) - back")

		var input string
		fmt.Scanln(&input)

		input = strings.ToLower(input)

		switch input {
		case "u":

		case "d":

		case "r":

		case "l":

		case "x":
			return
		}
	}
}
