package display

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/atarte/cli-cd/config"
	"github.com/atarte/cli-cd/data"
)

// MenuScreen
func MenuScreen(appdata *data.AppData) {
	titleScreen()

	appListScreen(appdata)

	fmt.Println("")
	// fmt.Println("(1) - View app list")
	// fmt.Println("(%index) - App details")
	fmt.Println("(a) - Add app")
	fmt.Println("(u) - Update app")
	fmt.Println("(s) - Delete app")
	fmt.Println("(i) - App info (need to put a bit more, like config file path and other)")
	fmt.Println("(q) - Quit")

	var input string
	fmt.Scanln(&input)

	input = strings.ToLower(input)

	switch input {
	// case "1":
	// 	AppListScreen(appdata)
	case "a":
		addScreen(appdata)
	case "u":
		// UpdateApp()
	case "d":
		// DeleteApp()
	case "i":
		infoScreen()
	case "q":
		os.Exit(0)
	default:
		// app detail
		for i := range appdata.MonitoredAppList {
			if input == strconv.Itoa(i) {
				detailScreen(appdata, i)
			}
		}
	}
}

// clearScreen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// titleScreen
func titleScreen() {
	clearScreen()

	fmt.Println("  ___ _    ___     ___ ___  ")
	fmt.Println(" / __| |  |_ _|__ / __|   \\ ")
	fmt.Println("| (__| |__ | |___| (__| |) |")
	fmt.Println(" \\___|____|___|   \\___|___/ ")
}

// returnScreen
func returnScreen() {
	fmt.Println("")
	fmt.Println("(Press a key to return...)")
	fmt.Scanln()
}

// infoScreen
func infoScreen() {
	titleScreen()

	fmt.Println()
	fmt.Printf("App name:%10s\n", config.AppName)
	fmt.Printf("Version:%10s\n", config.Version)
	fmt.Printf("Copyright: atarte ©2023-%s\n", config.ReleaseYear)

	returnScreen()
}

// appListScreen
func appListScreen(appdata *data.AppData) {
	fmt.Println("")
	// do a array style list similar to a docker ps
	fmt.Printf("      %-10v %-20v (status and other param)\n", "Name", "Path")
	for i, monitoredApp := range appdata.MonitoredAppList {
		fmt.Printf("(%v) - %-10v %-20v\n", i, monitoredApp.Name, monitoredApp.Path)
	}
}
