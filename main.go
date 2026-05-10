package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
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
		"map": {
			name:        "map",
			description: "Displays next Pokemon map data",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous Pokemon map data",
			callback:    commandMapB,
		},
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	saveCommand := commandInfo{}

	for {
		fmt.Printf("Pokedex >")
		scanner.Scan()
		text := scanner.Text()
		clean := cleanInput(text)

		command, exists := getCommands()[clean[0]]
		if exists {
			command.callback(&saveCommand)
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}
func commandExit(config *commandInfo) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *commandInfo) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(config *commandInfo) error {
	url := config.next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data := pokeAPIResponse{}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return err
	}
	for _, item := range data.Results {
		fmt.Println(item.Name)
	}
	config.next = data.Next
	config.previous = data.Previous

	return nil
}

func commandMapB(config *commandInfo) error {
	config.next = config.previous
	commandMap(config)
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

type commandInfo struct {
	next     string
	previous string
}

type clientCommand struct {
	name        string
	description string
	callback    func(*commandInfo) error
}
