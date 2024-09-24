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
	}
	for {
		fmt.Print("Pokedex> ")
		reader.Scan()
		text := reader.Text()
		command := cleanInput(text)

		if _, ok := getCliCommands(cfg)[command]; !ok {
			fmt.Println("command not found")
			continue
		}

		err := getCliCommands(cfg)[command].callback(cfg)
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

	}
}

func cleanInput(text string) string {
	return strings.ToLower(text)
}
