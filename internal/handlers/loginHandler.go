package handlers

import (
	"context"
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
	ctx := context.Background()
	//check if user exixts
	_, err := s.DB.GetUser(ctx, user)
	if err != nil {
		return fmt.Errorf("User already exists in system")
	}
	s.SetCurrentUser(user)
	return SaveConfig(s)
}
