package handlers

import (
	"fmt"

	"example.com/eiriktaa/gator/internal/commands"
	"example.com/eiriktaa/gator/internal/state"
)

type CommandHandlers struct {
	Handlers map[string]func(*state.State, commands.Command) error
}

func NewCLIHandler() CommandHandlers {
	return CommandHandlers{
		Handlers: map[string]func(*state.State, commands.Command) error{
			"login": handlerLogin,
		},
	}
}

func (c CommandHandlers) Run(s *state.State, cmd commands.Command) error {
	fnName := cmd.Name
	handler, ok := c.Handlers[fnName]
	if !ok {
		return fmt.Errorf("%s is not a valid cmd", fnName)
	}
	return handler(s, cmd)
}

func SaveConfig(s *state.State) error {
	cfg := s.Config
	return cfg.WriteToFile()
}
