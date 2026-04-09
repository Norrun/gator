package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"strings"
	"time"

	"github.com/Norrun/gator/internal/bt"
	"github.com/Norrun/gator/internal/database"
	"github.com/Norrun/gator/internal/rss"
	"github.com/google/uuid"
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
	fmt.Printf("FETCHING: %s", feed.Name)
	rssFeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	var errsum error
	for _, item := range rssFeed.Channel.Items {

		_, ierr := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       makeNullString(item.Title),
			Url:         item.Link,
			Description: makeNullString(item.Description),
			PublishedAt: makeNullTime(item.PubDate),
			FeedID:      feed.ID,
		})
		if ierr != nil && strings.Contains(ierr.Error(), "duplicate key value violates unique constraint \"posts_url_key\"") == false {
			errsum = errors.Join(errsum, ierr)
		}
	}

	if errsum != nil {
		return errsum
	}

	//old

	return nil
}

func makeNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}
func makeNullTime(s string) sql.NullTime {
	tim, err := bt.FallbackArg(
		func(st string) (time.Time, error) { return time.Parse(st, s) },
		time.RFC1123Z,
	)
	if err != nil {
		log.Println(err.Error(), s)
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: tim, Valid: true}

}
