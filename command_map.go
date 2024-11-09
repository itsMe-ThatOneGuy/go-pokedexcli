package main

import (
	"fmt"
)

func commandMapf(conf *config) error {
	locations, err := conf.pokeapiClient.ListLocations(conf.nextLocation)
	if err != nil {
		return err
	}

	conf.nextLocation = locations.Next
	conf.prevLocation = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(conf *config) error {
	if conf.prevLocation == nil {
		return fmt.Errorf("No previouse page of locations: On the first page")
	}

	locations, err := conf.pokeapiClient.ListLocations(conf.prevLocation)
	if err != nil {
		return err
	}

	conf.nextLocation = locations.Next
	conf.prevLocation = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
