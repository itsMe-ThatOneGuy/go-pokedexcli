package main

import "github.com/itsMe-ThatOneGuy/go-pokedexcli/internal/pokeapi"

func main() {
	pokeClient := pokeapi.NewClient()
	conf := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	repl(conf)
}
