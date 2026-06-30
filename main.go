package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/GeorgievPlamen/rss-feed/internal/commands"
	"github.com/GeorgievPlamen/rss-feed/internal/config"
	"github.com/GeorgievPlamen/rss-feed/internal/database"
)

func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Printf("\nCould not read config. Err: %v", err)
		os.Exit(-1)
	}

	db, err := sql.Open("postgres", config.DbUrl)
	if err != nil {
		fmt.Printf("Connecion to PostgreSQL failed: %v", err)
		os.Exit(-1)
	}

	dbQueries := database.New(db)

	state := commands.State{
		Config: &config,
		Db:     dbQueries,
	}

	availableCommands := commands.Commands{
		All: map[string]func(*commands.State, commands.Command) error{},
	}

	availableCommands.Register("login", commands.HandlerLogin)
	availableCommands.Register("register", commands.HandlerRegister)
	availableCommands.Register("reset", commands.HandlerReset)
	availableCommands.Register("users", commands.HandlerUsers)
	availableCommands.Register("agg", commands.HandlerAgg)
	availableCommands.Register("feeds", commands.HandlerFeeds)
	availableCommands.Register("addfeed", commands.MiddlewareLoggedIn(commands.HandlerAddFeed))
	availableCommands.Register("follow", commands.MiddlewareLoggedIn(commands.HandlerFollow))
	availableCommands.Register("following", commands.MiddlewareLoggedIn(commands.HandlerFollowing))
	availableCommands.Register("unfollow", commands.MiddlewareLoggedIn(commands.HandlerUnfollow))

	if len(os.Args) < 2 {
		fmt.Printf("\n You need to provide atleast one argument.")
		os.Exit(-1)
	}

	commandArgs := os.Args[1:]

	command := commands.Command{
		Name: commandArgs[0],
		Args: commandArgs[1:],
	}

	err = availableCommands.Run(&state, command)
	if err != nil {
		fmt.Printf("\nCould not run command. Error: %v", err)
		os.Exit(-1)
	}
}
