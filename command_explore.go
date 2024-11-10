package main

import (
	"errors"
	"fmt"
)

func commandExplore(conf *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("You need to provide a location name")
	}

	name := args[0]
	location, err := conf.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Pokemon Found:")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
