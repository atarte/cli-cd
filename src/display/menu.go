package display

import (
	"fmt"
	"os"
)

// MenuScreen
func MenuScreen() {
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
		// ViewAppList(data)
	case "2":
		// AddApp(data)
	case "3":
		// UpdateApp()
	case "4":
		// DeleteApp()
	case "5":
		infoScreen()
	case "6":
		os.Exit(0)
	default:
		fmt.Println("pas cool")
	}
}
