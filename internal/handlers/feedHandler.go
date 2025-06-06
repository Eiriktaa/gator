package handlers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"example.com/eiriktaa/gator/internal/commands"
	"example.com/eiriktaa/gator/internal/database"
	"example.com/eiriktaa/gator/internal/state"
	"example.com/eiriktaa/gator/rss"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state.State, cmd commands.Command, currentUser database.User) error {
	args := cmd.Args
	if len(args) != 2 {
		return fmt.Errorf("Expected addfeed <name> <url>")
	}
	name := args[0]
	url := args[1]
	//empty context
	ctx := context.Background()
	feed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	feedRecord, err := s.DB.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	})
	if err != nil {
		return err
	}

	_, err = s.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feedRecord.ID,
	})
	if err != nil {
		return err
	}

	feed.DisplayData()
	return nil
}

func handlerFeeds(s *state.State, cmd commands.Command) error {
	ctx := context.Background()

	data, err := s.DB.GetFeeds(ctx)
	if err != nil {
		return err
	}

	for _, item := range data {
		fmt.Println(fmt.Sprintf("%s - %s <%s>", item.Name, item.Username, item.Url))
	}
	return nil
}
func handlerBrowse(s *state.State, cmd commands.Command, currentUser database.User) error {
	args := cmd.Args
	limit := 2
	if len(args) >= 1 {
		val, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		limit = val
	}
	//empty context
	ctx := context.Background()
	posts, err := s.DB.GetPosts(ctx, currentUser.ID)
	if err != nil {
		return err
	}

	for i, post := range posts {
		if i >= limit {
			return nil
		}
		fmt.Println("*******************")
		fmt.Println("-", post.Title)
		fmt.Println("-", post.Url)
		fmt.Println("-", post.Description)
		fmt.Println("-", post.PublishedAt)

		fmt.Println("*******************")
	}
	return nil

}
