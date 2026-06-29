package main

import (
	"fmt"
	"os"

	"github.com/GeorgievPlamen/rss-feed/internal/commands"
	"github.com/GeorgievPlamen/rss-feed/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Printf("\nCould not read config. Err: %v", err)
		os.Exit(-1)
	}

	state := commands.State{
		Config: &config,
	}

	availableCommands := commands.Commands{
		All: map[string]func(*commands.State, commands.Command) error{},
	}

	availableCommands.Register("login", commands.HandlerLogin)

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
