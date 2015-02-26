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
	fmt.Printf("\x1b[31;1mSwarmsearch version %s\x1b[0m \n", version.Version)
	return
}
