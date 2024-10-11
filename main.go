package main

import (
	"database/sql"
	"fmt"
	"os"
	"rss/internal/database"

	_ "github.com/lib/pq"

	"github.com/uzairkhan98/rss/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("postgres", config.DbURL)
	if err != nil {
		panic(err)
	}
	dbQueries := database.New(db)

	statePoint := &state{config: config, db: dbQueries}
	commands := &commands{
		mapper: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegistration)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUserList)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerGetFeedList)
	commands.register("follow", middlewareLoggedIn(handlerAddFollow))
	commands.register("following", middlewareLoggedIn(handlerGetUserFollowedFeeds))

	args := os.Args

	if len(args) <= 1 {
		fmt.Println("you need to provide at least two cli arguments")
		os.Exit(1)
	}
	command := &command{name: args[1], args: args[2:]}
	err = commands.run(statePoint, *command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
