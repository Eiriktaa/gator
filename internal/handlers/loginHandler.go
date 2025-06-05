package handlers

import (
	"fmt"

	"example.com/eiriktaa/gator/internal/commands"
	"example.com/eiriktaa/gator/internal/state"
)

func handlerLogin(s *state.State, cmd commands.Command) error {
	args := cmd.Args
	if len(args) != 1 {
		return fmt.Errorf("Expected only username as argument")
	}
	user := args[0]
	s.SetCurrentUser(user)
	return SaveConfig(s)
}
