package main

import (
	"fmt"
)

func commandPokedex(conf *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range conf.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
