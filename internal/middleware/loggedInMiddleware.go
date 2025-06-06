package middleware

import (
	"context"

	"example.com/eiriktaa/gator/internal/commands"
	"example.com/eiriktaa/gator/internal/database"
	"example.com/eiriktaa/gator/internal/state"
)

func MiddlewareLoggedIn(handler func(
	s *state.State,
	cmd commands.Command,
	user database.User) error) func(*state.State, commands.Command) error {

	return func(s *state.State, c commands.Command) error {
		user, err := s.DB.GetUser(context.Background(), s.Config.Current_user_name)
		if err != nil {
			return err
		}
		return handler(s, c, user)
	}

}
