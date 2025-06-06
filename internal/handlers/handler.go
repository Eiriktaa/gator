package handlers

import (
	"context"
	"fmt"

	"example.com/eiriktaa/gator/internal/commands"
	"example.com/eiriktaa/gator/internal/state"
	"example.com/eiriktaa/gator/rss"
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
			"addfeed":   handlerAddFeed,
			"feeds":     handlerFeeds,
			"follow":    handlerFollow,
			"following": handlerFollowing,
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
func handleAgg(s *state.State, cmd commands.Command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	feed.DisplayData()
	return nil
}
