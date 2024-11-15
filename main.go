package main

import (
	"bufio"
	"fmt"
	"os"
)

// Struct to handle pagination
type Config struct {
	Next     string
	Previous *string
}

// Struct to handle CLI commands
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Command function definitions
func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\nmap: Get the next page of locations\nmapb: Get the previous page of locations\n")
	return nil
}

func commandExit() error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}

func commandMap(config *Config) error {
	// Use current Next URL in config
	if config.Next == "" {
		return fmt.Errorf("no next page available")
	}

	firstResults, err := getMap(config.Next)
	if err != nil {
		return err
	}

	// Print location names
	for _, result := range firstResults.Results {
		fmt.Println(result.Name)
	}

	// Update Config with new pagination URLs
	config.Next = firstResults.Next
	config.Previous = firstResults.Previous
	return nil
}

func commandMapb(config *Config) error {
	// Use Previous URL in config to go back a page if it exists
	if config.Previous == nil || *config.Previous == "" {
		return fmt.Errorf("no previous page available")
	}

	firstResults, err := getMap(*config.Previous)
	if err != nil {
		return err
	}

	// Print location names
	for _, result := range firstResults.Results {
		fmt.Println(result.Name)
	}

	// Update Config with new pagination URLs
	config.Next = firstResults.Next
	config.Previous = firstResults.Previous
	return nil
}

func main() {
	// Initialize Config with the first page URL
	config := Config{Next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"}

	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    func() error { return commandMap(&config) },
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    func() error { return commandMapb(&config) },
		},
	}

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !reader.Scan() {
			break
		}
		input := reader.Text()

		// Execute command
		if command, found := commands[input]; found {
			err := command.callback()
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
		} else {
			fmt.Println("Invalid command. Type 'help' for a list of commands.")
		}
	}
}
