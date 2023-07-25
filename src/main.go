package main

import (
	"github.com/atarte/cli-cd/data"
	"github.com/atarte/cli-cd/display"
)

// var (
// 	AppName string = "cli-cd"
// 	Version string = "0.0.0"
// 	// Build   string = "abcd"
// )

func main() {
	var appdata data.AppData = data.LoadCliCdConfigFile()

	for {
		display.MenuScreen(&appdata)

	}

	// err := config.CreateCliCdConfigFile()
	// fmt.Println("err", err)

}
