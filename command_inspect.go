package main

import (
	"errors"
	"fmt"
)

func commandInspect(conf *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("You need to provide a Pokemon name")
	}

	pokemon, ok := conf.caughtPokemon[args[0]]
	if !ok {
		fmt.Println("You have not caught this Pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf(" - %s\n", types.Type.Name)
	}

	return nil
}
