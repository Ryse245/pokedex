package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getCommands() map[string]clientCommand {
	return map[string]clientCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex >")
		scanner.Scan()
		text := scanner.Text()
		clean := cleanInput(text)

		command, exists := getCommands()[clean[0]]
		if exists {
			command.callback()
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func cleanInput(text string) []string {
	format := strings.TrimSpace(text)

	if format == "" {
		return []string{}
	}

	ret := strings.Fields(format)

	for i := range ret {
		ret[i] = strings.ToLower(ret[i])
	}

	return ret
}

type clientCommand struct {
	name        string
	description string
	callback    func() error
}
