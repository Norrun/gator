package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"strings"

	"github.com/Norrun/gator/internal/helpers"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	reader := strings.Reader{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, &reader)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	req.Header.Set("User-Agent", "gator")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	feed := RSSFeed{}
	err = xml.Unmarshal(data, &feed)
	fixFeedText(&feed)
	return &feed, nil
}

func fixFeedText(feed *RSSFeed) {
	unescaper := helpers.NewModifyer(html.UnescapeString)
	unescaper(&feed.Channel.Title)
	unescaper(&feed.Channel.Description)
	for i := 0; i < len(feed.Channel.Items); i++ {
		unescaper(&feed.Channel.Items[i].Title)
		unescaper(&feed.Channel.Items[i].Description)
	}
}
