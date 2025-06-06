package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

func getRequest(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")
	client := http.Client{
		Timeout: time.Second * 5,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf(fmt.Sprintf("HTTP request failed with code %d", res.StatusCode))
	}
	return res, nil
}
func unmashalFeedResponse(res *http.Response) (*RSSFeed, error) {
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var feed RSSFeed
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}
	return &feed, nil
}
func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	res, err := getRequest(ctx, feedURL)
	if err != nil {
		return nil, err
	}
	feed, err := unmashalFeedResponse(res)
	if err != nil {
		return nil, err
	}
	feed.unescapeText()
	return feed, nil
}
