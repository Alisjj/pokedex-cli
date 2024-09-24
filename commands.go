package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alisjj/pokedex/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	cache pokecache.Cache
	l     Location
}

func getCliCommands(*config) map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func(*config) error { return commandHelp() },
		},
		"clear": {
			name:        "clear",
			description: "Cleans the screen",
			callback:    func(*config) error { return commandClear() },
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous names of 20 location areas",
			callback:    commandMapB,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func(*config) error { return commandExit() },
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

	for _, i := range getCliCommands(nil) {
		fmt.Printf("\n%v: %v", i.name, i.description)

	}
	fmt.Print("\n\n")
	return nil
}

func commandMap(c *config) error {

	var url string
	if c.l.Next != nil {
		url = *c.l.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	err := c.getLocations(url)
	if err != nil {
		return err
	}

	for _, r := range c.l.Results {
		fmt.Println(r.Name)
	}

	return nil

}

func commandMapB(c *config) error {

	var url string
	if c.l.Previous == nil {
		return fmt.Errorf("there is no previous page")

	}
	url = *c.l.Previous

	err := c.getLocations(url)
	if err != nil {
		return err
	}

	for _, r := range c.l.Results {
		fmt.Println(r.Name)
	}

	return nil

}
