package handlers

import (
	"context"
	"fmt"

	"example.com/eiriktaa/gator/internal/commands"
	"example.com/eiriktaa/gator/internal/state"
)

func handlerUsers(s *state.State, cmd commands.Command) error {
	ctx := context.Background()
	users, err := s.DB.GetUsers(ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		prefix := "* "
		suffix := ""
		if user.Name == s.Config.Current_user_name {
			suffix = "(current)"
		}
		fmt.Println(prefix, user.Name, suffix)
	}
	return nil
}
