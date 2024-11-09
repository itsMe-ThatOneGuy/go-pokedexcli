package main

import (
	"os"
)

func commandExit(conf *config, args ...string) error {
	os.Exit(0)
	return nil
}
