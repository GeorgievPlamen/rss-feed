package commands

import (
	"fmt"

	"github.com/GeorgievPlamen/rss-feed/internal/config"
)

type State struct {
	Config *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	All map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	command, ok := c.All[cmd.Name]
	if !ok {
		return fmt.Errorf("Command not found: %v", cmd.Name)
	}

	fmt.Println(cmd)

	err := command(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.All[name] = f
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Login expects a user as an argument")
	}

	err := s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("\nUser has been set to: %s", cmd.Args[0])
	return nil
}
