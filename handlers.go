package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}
	err := s.config.SetUser(cmd.args[0])

	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}
	fmt.Println("set user to ", cmd.args[0])
	return nil
}
