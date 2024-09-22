package main

import (
	"fmt"
	"os"
	"os/exec"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCliCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"clear": {
			name:        "clear",
			description: "Cleans the screen",
			callback:    commandClear,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandClear() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}

func commandHelp() error {
	fmt.Println("\nWelcome to Pokedex!")
	fmt.Println("Usage:")

	for _, i := range getCliCommands() {
		fmt.Printf("\n%v: %v", i.name, i.description)

	}
	fmt.Print("\n\n")
	return nil
}
