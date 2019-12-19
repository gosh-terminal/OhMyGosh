package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func initInteractive() {
	prompt := promptui.Select{
		Label: "What are you configuring",
		Items: []string{"Plugin", "Theme", "Never Mind"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		println("Error, exiting...")
		os.Exit(1)
	}
	if result == "Never Mind" {
		fmt.Println("Ok, bye!")
		os.Exit(0)
	} else if result == "Plugin" {
		fmt.Println("Plugin, Cool!")
	} else if result == "Theme" {
		fmt.Println("Theme, got it!")
	} else {
		return
	}
}

func main() {
	initInteractive()
}
