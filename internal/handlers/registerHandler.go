package handlers

import (
	"context"
	"fmt"
	"time"

	"example.com/eiriktaa/gator/internal/commands"
	"example.com/eiriktaa/gator/internal/database"
	"example.com/eiriktaa/gator/internal/state"
	"github.com/google/uuid"
)

func handlerRegister(s *state.State, cmd commands.Command) error {
	args := cmd.Args
	if len(args) != 1 {
		return fmt.Errorf("Expected only username as argument")
	}
	userName := args[0]

	//empty context
	ctx := context.Background()
	//check if user exixts
	_, err := s.DB.GetUser(ctx, userName)
	if err == nil {
		return fmt.Errorf("User already exists in system")
	}
	s.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	})
	s.Config.SetUser(userName)
	return SaveConfig(s)
}
