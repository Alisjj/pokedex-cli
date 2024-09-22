package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
	if c.Next != nil {
		url = *c.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&c)
	if err != nil {
		return err
	}

	for _, r := range c.Results {
		fmt.Println(r.Name)
	}

	return nil

}

func commandMapB(c *config) error {

	var url string
	if c.Previous == nil {
		return fmt.Errorf("there is not previous page\n")

	}
	url = *c.Previous
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&c)
	if err != nil {
		return err
	}

	for _, r := range c.Results {
		fmt.Println(r.Name)
	}

	return nil

}
