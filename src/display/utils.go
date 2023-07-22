package display

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Title() {
	ClearScreen()

	fmt.Println("  ___ _    ___     ___ ___  ")
	fmt.Println(" / __| |  |_ _|__ / __|   \\ ")
	fmt.Println("| (__| |__ | |___| (__| |) |")
	fmt.Println(" \\___|____|___|   \\___|___/ ")
	fmt.Println("")
}
