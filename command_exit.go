package main

import (
	"os"
)

func commandExit(conf *config) error {
	os.Exit(0)
	return nil
}
