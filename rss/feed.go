package rss

import (
	"fmt"
	"html"
	"time"

	"example.com/eiriktaa/gator/internal/database"
	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (feed *RSSFeed) unescapeText() {
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].unescapeText()
	}
}
func (item *RSSItem) unescapeText() {
	fmt.Printf("BEFORE Title: %q\n", item.Title)
	fmt.Printf("BEFORE Description: %q\n", item.Description)

	item.Description = html.UnescapeString(item.Description)
	item.Title = html.UnescapeString(item.Title)

	fmt.Printf("AFTER Title: %q\n", item.Title)
	fmt.Printf("AFTER Description: %q\n", item.Description)
	fmt.Println("---")
}

func (item RSSItem) DisplayItem() {
	fmt.Println("-", item.Title)
	fmt.Println("-", item.Link)
	fmt.Println("-", item.Description)
	fmt.Println("-", item.PubDate)

}

func (feed *RSSFeed) DisplayData() {
	fmt.Println(feed.Channel.Title)
	fmt.Println(feed.Channel.Link)
	fmt.Println(feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		item.DisplayItem()
	}
}
func (feed *RSSFeed) GenerateInsertPostRecords(feedId uuid.UUID) []database.CreatePostsParams {
	var records []database.CreatePostsParams
	for _, item := range feed.Channel.Item {
		records = append(records, database.CreatePostsParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			PublishedAt: item.PubDate,
			Description: item.Description,
			Url:         item.Link,
			Title:       item.Title,
			FeedID:      feedId,
		})
	}
	return records

}
