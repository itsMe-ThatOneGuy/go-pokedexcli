package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/itsMe-ThatOneGuy/go-pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocation  *string
	prevLocation  *string
}

func repl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")

		scanner.Scan()
		input := parseInput(scanner.Text())

		command, ok := commands()[input]
		if ok {
			err := command.callback(conf)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func parseInput(input string) string {
	lower := strings.ToLower(input)
	trimmed := strings.TrimSpace(lower)
	return trimmed
}

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Ge the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}
}
