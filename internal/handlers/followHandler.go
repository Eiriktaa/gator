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

func handlerFollow(s *state.State, cmd commands.Command, currentUser database.User) error {
	args := cmd.Args
	if len(args) != 1 {
		return fmt.Errorf("Expected follow <url>")
	}
	url := args[0]
	//empty context
	ctx := context.Background()
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
func handlerFollowing(s *state.State, cmd commands.Command, currentUser database.User) error {
	ctx := context.Background()
	feeds, err := s.DB.GetFollowings(ctx, currentUser.ID)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed)
	}
	return nil
}
func handlerUnfollow(s *state.State, cmd commands.Command, currentUser database.User) error {
	args := cmd.Args
	if len(args) != 1 {
		return fmt.Errorf("Expected unfollow <url>")
	}
	url := args[0]
	//empty context
	ctx := context.Background()
	_, err := s.DB.DeleteFollowing(ctx,
		database.DeleteFollowingParams{
			UserID: currentUser.ID,
			Url:    url})
	if err != nil {
		return err
	}
	return nil
}
