package main

import (
	"example.com/eiriktaa/gator/internal/commands"
	"example.com/eiriktaa/gator/internal/handlers"
	"example.com/eiriktaa/gator/internal/state"
	"fmt"
	"os"
)

func main() {
	state := state.InitalizeState()
	clicmds := handlers.NewCLIHandler()

	name, args := validateArgs()
	cmd := commands.Command{name, args}
	err := clicmds.Run(&state, cmd)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
		os.Exit(1)
	}
}

func validateArgs() (string, []string) {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("No command specified")
		os.Exit(1)
	}
	command := args[1]

	if len(args) >= 3 {
		return command, args[2:]
	}
	return command, nil
}
