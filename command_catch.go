package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(conf *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("You need to provide a Pokemon name")
	}

	name := args[0]
	pokemon, err := conf.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("You threw a Pokeball at %s\n", pokemon.Name)
	fmt.Println("*shake shake*")
	time.Sleep(1 * time.Second)
	fmt.Println("*shake shake*")
	time.Sleep(2 * time.Second)

	catchChance := rand.Intn(pokemon.BaseExperience)
	if catchChance > pokemon.BaseExperience/2 {
		fmt.Printf("%s escaped..\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	conf.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
