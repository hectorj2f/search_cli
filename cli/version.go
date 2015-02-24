package main

import (
	"fmt"

	"github.com/hectorj2f/search_cli/version"
)

var cmdVersion = &Command{
	Name:        "version",
	Description: "Print the version and exit",
	Summary:     "Print the version and exit",
	Run:         runVersion,
}

func init() {
	commands = append(commands, cmdVersion)
}

func runVersion(args []string) (exit int) {
	fmt.Printf("swarmsearch version %s\n", version.Version)
	return
}
