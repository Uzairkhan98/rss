package main

import (
	"context"
	"fmt"
	"os"
	"rss/internal/database"
	"time"

	"github.com/google/uuid"
)

func handleGetUserFollowedFeeds(s *state, _ command) error {
	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}
	for _, feed := range followedFeeds {
		fmt.Printf("Feed Name: %s\n, User Name: %s\n", feed.FeedName.String, feed.UserName.String)
	}
	return nil
}

func handleAddFollow(s *state, c command) error {
	if len(c.args) < 1 {
		return fmt.Errorf("please provide a feed URL")
	}
	currentUser, feed, err := fetchFeedURLAndUsers(c.args[0], s)
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFollow(context.Background(), database.CreateFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return err
	}

	if feedFollow.FeedName.Valid && feedFollow.UserName.Valid {
		fmt.Printf("Feed Name: %s\n, User Name: %s\n", feedFollow.FeedName.String, feedFollow.UserName.String)
	} else {
		fmt.Println("Feed Name or User Name is not valid")
	}
	return nil
}

func handleGetFeedList(s *state, _ command) error {
	feeds, err := s.db.GetFeedList(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Println(feed)
	}
	return nil
}

func handleAddFeed(s *state, c command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}

	if len(c.args) < 2 {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:      c.args[0],
		Url:       c.args[1],
		UserID:    currentUser.ID,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	_, err = s.db.CreateFollow(context.Background(), database.CreateFollowParams{
		UserID:    currentUser.ID,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

func handleAgg(_ *state, _ command) error {
	res, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("Feed Aggregator:\n")
	fmt.Println(res)
	return nil
}

func handlerUserList(s *state, _ command) error {
	users, err := s.db.GetUserList(context.Background())

	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, user := range users {
		temp := ""
		if user.Name == s.config.CurrentUserName {
			temp = "(current)"
		}
		fmt.Printf("* %s %s\n", user.Name, temp)
	}
	return nil
}

func handlerReset(s *state, _ command) error {
	err := s.db.ResetUsers(context.Background())

	if err != nil {
		fmt.Println(err)
		return err
	}
	os.Exit(0)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])

	if err != nil {
		return err
	}

	err = s.config.SetUser(user.Name)

	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}
	fmt.Println("set user to ", user.Name)
	return nil
}

func handlerRegistration(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the registration handler expects a single argument, the username")
	}

	temp := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), temp)

	if err != nil {
		return err
	}

	err = s.config.SetUser(user.Name)

	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}

	fmt.Println("created user: ", user)

	return nil
}
