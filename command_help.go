package main

import (
	"fmt"
)

func commandHelp(conf *config, args ...string) error {
	fmt.Println()
	fmt.Println("*Beep* Booting Pokedex *Boop*")
	fmt.Println("Usage:")
	for _, cmd := range commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()

	return nil
}
