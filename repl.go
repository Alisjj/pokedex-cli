package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alisjj/pokedex/pokecache"
)

func repl() {
	reader := bufio.NewScanner(os.Stdin)
	cfg := &config{
		cache: pokecache.NewCache(5 * time.Minute),
		l:     Location{},
		p:     NewPokedex(),
	}
	for {
		fmt.Print("Pokedex> ")
		reader.Scan()
		text := reader.Text()
		command := cleanInput(text)

		if _, ok := getCliCommands()[command[0]]; !ok {
			fmt.Println("command not found")
			continue
		}

		if len(command) > 1 {
			err := getCliCommands()[command[0]].callback(cfg, command[1])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			continue
		}
		err := getCliCommands()[command[0]].callback(cfg, "")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
