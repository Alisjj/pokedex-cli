package main

import (
	"fmt"
	"os"
	"os/exec"
)

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
		"explore": {
			name:        "explore",
			description: "Displays details about a particular relationship",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catches Pokemon and adds them to the user's Pokedex.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "See details about a Pokemon.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "inspect",
			description: "See details about a Pokemon.",
			callback:    commandInspect,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandExit(cfg *config, url string) error {
	os.Exit(0)
	return nil
}

func commandClear(cfg *config, url string) error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}

func commandHelp(cfg *config, url string) error {
	fmt.Println("\nWelcome to Pokedex!")
	fmt.Println("Usage:")

	for _, i := range getCliCommands() {
		fmt.Printf("\n%v: %v", i.name, i.description)

	}
	fmt.Print("\n\n")
	return nil
}

func commandMap(c *config, v string) error {

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

func commandMapB(c *config, v string) error {

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

func commandExplore(c *config, area string) error {
	err := c.exploreArea(area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s....\n", area)
	fmt.Println("Found Pokemon:")

	for _, r := range c.c.PokemonEncounters {
		fmt.Printf(" - %s \n", r.Pokemon.Name)
	}

	return nil
}

func commandCatch(c *config, name string) error {
	pok, err := c.getPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pok.Name)

	catch := catch(pok.BaseExperience)

	if catch {
		c.p.pokemon[pok.Name] = pok
		fmt.Printf("%v was caught!\n", pok.Name)
		return nil
	}

	fmt.Printf("%v escaped!\n", pok.Name)
	return nil

}

func commandInspect(c *config, name string) error {
	pokemon, ok := c.p.pokemon[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s:\n", pokemon.Name)
	fmt.Printf("Height: %v:\n", pokemon.Height)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf(" -%v: %v:\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" -%v\n", t.Type.Name)
	}

	return nil

}
