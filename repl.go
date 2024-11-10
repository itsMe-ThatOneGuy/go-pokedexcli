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
	caughtPokemon map[string]pokeapi.Pokemon
}

func repl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")

		scanner.Scan()
		input := parseInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		cmdName := input[0]

		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		command, ok := commands()[cmdName]
		if ok {
			err := command.callback(conf, args...)
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

func parseInput(input string) []string {
	lower := strings.ToLower(input)
	feilds := strings.Fields(lower)
	return feilds
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch <pokemon-name",
			description: "Attempt to catch",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore <location-name>",
			description: "Explore a location",
			callback:    commandExplore,
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
