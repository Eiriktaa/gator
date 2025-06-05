package handlers

import (
	"context"
	"example.com/eiriktaa/gator/internal/commands"
	"example.com/eiriktaa/gator/internal/state"
)

func handlerReset(s *state.State, cmd commands.Command) error {
	ctx := context.Background()
	return s.DB.TruncateUsers(ctx)
}
