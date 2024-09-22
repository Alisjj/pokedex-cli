package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		reader.Scan()
		text := reader.Text()
		command := cleanInput(text)

		if _, ok := getCliCommands()[command]; !ok {
			fmt.Println("command not found")
			continue
		}

		err := getCliCommands()[command].callback()
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

	}
}

func cleanInput(text string) string {
	return strings.ToLower(text)
}
