package main

import (
	"fmt"

	"github.com/atarte/cli-cd/config"
	"github.com/atarte/cli-cd/display"
)

var (
	AppName string = "cli-cd"
	Version string = "0.0.0"
	// Build   string = "abcd"
)

func main() {
	display.Title()

	err := config.CreateCliCdConfigFile()
	fmt.Println("err", err)

}
