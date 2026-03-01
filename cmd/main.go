// Package main is the entrypoint for the meal planning service.
package main

import (
	"os"

	"github.com/b-sea/meal-planner/cmd/cli"
)

var version = ""

func main() {
	if err := cli.New(version).Execute(); err != nil {
		os.Exit(1)
	}
}
