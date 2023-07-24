package display

import (
	"fmt"
	"os"
	"os/exec"
)

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
	fmt.Println("")
}

// returnScreen
func returnScreen() {
	fmt.Println("")
	fmt.Println("(Press a key to return...)")
	fmt.Scanln()
}

// isInputValid
// func isInputValid(input string, constrainInputList []string) bool {
// 	for _, constrain := range constrainInputList {
// 		if input == constrain {
// 			return true
// 		}
// 	}

// 	return false
// }

// func isInputValid[T comparable](input T, constrainInputList []T) bool {

// 	for _, constrain := range constrainInputList {
// 		if input == constrain {
// 			return true
// 		}
// 	}

// 	return false
// }
