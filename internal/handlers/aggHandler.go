package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"example.com/eiriktaa/gator/internal/commands"
	"example.com/eiriktaa/gator/internal/database"
	"example.com/eiriktaa/gator/internal/state"
	"example.com/eiriktaa/gator/rss"
	"github.com/lib/pq"
)

func handleAgg(s *state.State, cmd commands.Command) error {
	if len(cmd.Args) != 1 {
		fmt.Errorf("Expected agg <duration>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
func scrapeFeeds(s *state.State) error {
	ctx := context.Background()
	feed, err := s.DB.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	feedData, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	postRecords := feedData.GenerateInsertPostRecords(feed.ID)
	for _, post := range postRecords {
		if _, err := s.DB.CreatePosts(ctx, post); err != nil {
			if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
				continue
				// Handle the unique constraint violation - maybe print a message?
			} else if err != nil {
				// Handle other errors - return them
				return err
			}

			fmt.Printf("Error type: %T, value: %v\n", err, err)
			return err
		}
	}
	_, err = s.DB.MarkFeedFetch(ctx, database.MarkFeedFetchParams{
		LastFetchedAt: sql.NullTime{time.Now(), true},
		ID:            feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println("Fetching: ", feedData.Channel.Title)
	return nil
}
