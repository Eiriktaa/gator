package handlers

import (
	"fmt"

	"example.com/eiriktaa/gator/internal/commands"
	"example.com/eiriktaa/gator/internal/middleware"
	"example.com/eiriktaa/gator/internal/state"
)

type CommandHandlers struct {
	Handlers map[string]func(*state.State, commands.Command) error
}

func NewCLIHandler() CommandHandlers {
	return CommandHandlers{
		Handlers: map[string]func(*state.State, commands.Command) error{
			"login":     handlerLogin,
			"register":  handlerRegister,
			"reset":     handlerReset,
			"users":     handlerUsers,
			"agg":       handleAgg,
			"addfeed":   middleware.MiddlewareLoggedIn(handlerAddFeed),
			"feeds":     handlerFeeds,
			"follow":    middleware.MiddlewareLoggedIn(handlerFollow),
			"following": middleware.MiddlewareLoggedIn(handlerFollowing),
			"unfollow":  middleware.MiddlewareLoggedIn(handlerUnfollow),
			"browse":    middleware.MiddlewareLoggedIn(handlerBrowse),
		},
	}
}

func (c *CommandHandlers) Run(s *state.State, cmd commands.Command) error {
	handler, ok := c.Handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("%s is not a valid cmd", cmd.Name)
	}
	return handler(s, cmd)
}

func SaveConfig(s *state.State) error {
	cfg := s.Config
	return cfg.WriteToFile()
}
