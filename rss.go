package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Norrun/gator/internal/database"
	"github.com/Norrun/gator/internal/rss"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	_, err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        feed.ID,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}
	fmt.Println("debug::fetching feed")
	rssFeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	for _, item := range rssFeed.Channel.Items {
		fmt.Println(item.Title)
	}
	return nil
}
