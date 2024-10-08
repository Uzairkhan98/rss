package main

import (
	"fmt"
	"rss/internal/database"

	"github.com/uzairkhan98/rss/config"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

type command struct {
	name string
	args []string
}

type commands struct {
	mapper map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.mapper[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	err := c.mapper[cmd.name](s, cmd)

	if err != nil {
		return fmt.Errorf("error running command '%s': %w", cmd.name, err)
	}
	return nil
}
