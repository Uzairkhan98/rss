package main

import (
	"fmt"
	"os"

	"github.com/uzairkhan98/rss/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		panic(err)
	}
	statePoint := &state{config: config}
	commands := &commands{
		mapper: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)

	args := os.Args

	if len(args) <= 2 {
		fmt.Println("you need to provide at least two cli arguments")
		os.Exit(1)
	}
	command := &command{name: args[1], args: args[2:]}
	err = commands.run(statePoint, *command)
	if err != nil {
		fmt.Println(err)
	}
}
