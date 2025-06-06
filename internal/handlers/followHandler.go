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

func handlerFollow(s *state.State, cmd commands.Command) error {
	args := cmd.Args
	if len(args) != 1 {
		return fmt.Errorf("Expected follow <url>")
	}
	url := args[0]
	//empty context
	ctx := context.Background()
	//TODO optimalize duplicate requests
	currentUser, err := s.DB.GetUser(ctx, s.Config.Current_user_name)
	if err != nil {
		return fmt.Errorf("Unable to find current user")
	}

	feed, err := s.DB.GetFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("Unable to find requested feed, try registering")
	}
	_, err = s.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	return nil
}
func handlerFollowing(s *state.State, cmd commands.Command) error {

	ctx := context.Background()
	//TODO optimalize duplicate requests
	currentUser, err := s.DB.GetUser(ctx, s.Config.Current_user_name)
	if err != nil {
		return fmt.Errorf("Unable to find current user")
	}

	feeds, err := s.DB.GetFollowings(ctx, currentUser.ID)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed)
	}
	return nil

}
