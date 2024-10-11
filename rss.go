package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}
	rss, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	for _, item := range rss.Channel.Item {
		fmt.Println(item.Title)
	}
	return nil
}

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

// func fetchFeedURLAndUsers(feed_url string, s *state) (*database.User, *database.GetFeedByURLRow, error) {
// 	var (
// 		feed        database.GetFeedByURLRow // assuming Feed is your data structure for feed
// 		currentUser database.User            // assuming User is your data structure for user
// 		errFeed     error
// 		errUser     error
// 	)

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	// Run the first function to get the feed concurrently
// 	go func() {
// 		defer wg.Done()
// 		feed, errFeed = s.db.GetFeedByURL(context.Background(), feed_url)
// 		if errFeed != nil {
// 			log.Println("Error fetching feed:", errFeed)
// 			os.Exit(1)
// 		}
// 	}()

// 	// Run the second function to get the user concurrently
// 	go func() {
// 		defer wg.Done()
// 		currentUser, errUser = s.db.GetUser(context.Background(), s.config.CurrentUserName)
// 		if errUser != nil {
// 			log.Println("Error fetching user:", errUser)
// 		}
// 	}()

// 	// Wait for both goroutines to complete
// 	wg.Wait()

// 	// Check for errors
// 	if errFeed != nil {
// 		return &database.User{}, &database.GetFeedByURLRow{}, errFeed
// 	}

// 	if errUser != nil {
// 		return &database.User{}, &database.GetFeedByURLRow{}, errUser
// 	}

// 	return &currentUser, &feed, nil
// }
