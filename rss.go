package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Gator")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bod, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var feed RSSFeed

	err = xml.Unmarshal(bod, &feed)

	if err != nil {
		return nil, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feed.Channel.Item[i] = item
	}

	return &feed, nil
}
