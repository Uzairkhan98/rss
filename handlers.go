package main

import (
	"context"
	"fmt"
	"os"
	"rss/internal/database"
	"time"

	"github.com/google/uuid"
)

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
