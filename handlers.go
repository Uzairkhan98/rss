package main

import (
	"context"
	"fmt"
	"os"
	"rss/internal/database"
	"time"

	"github.com/google/uuid"
)

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
		os.Exit(1)
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
		os.Exit(1)
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
		os.Exit(1)
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
		os.Exit(1)
	}

	err = s.config.SetUser(user.Name)

	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}

	fmt.Println("created user: ", user)

	return nil
}
