package commands

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/GeorgievPlamen/rss-feed/internal/config"
	"github.com/GeorgievPlamen/rss-feed/internal/database"
	"github.com/GeorgievPlamen/rss-feed/internal/rss"
	"github.com/google/uuid"
)

type State struct {
	Config *config.Config
	Db     *database.Queries
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

	err := command(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.All[name] = f
}

const SqlNotFoundError string = "sql: no rows in result set"

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Login expects a user as an argument")
	}

	existingUser, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		if err.Error() == SqlNotFoundError {
			fmt.Println("User not  found!")
			os.Exit(1)
		}
		return err
	}

	err = s.Config.SetUser(existingUser.Name)
	if err != nil {
		return err
	}

	fmt.Printf("\nUser has been set to: %s", cmd.Args[0])
	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Register expects a name as an argument")
	}

	_, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		if err.Error() != SqlNotFoundError {
			return err
		}
	}

	user, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		Name: cmd.Args[0],
	})

	if err != nil {
		return err
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User created succefully.")
	fmt.Println(user)

	return nil
}

func HandlerReset(s *State, cmd Command) error {
	err := s.Db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Reset db")
	return nil
}

func HandlerUsers(s *State, cmd Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, v := range users {
		if s.Config.CurrentUserName == v.Name {
			fmt.Println(v.Name + " (current)")
		} else {
			fmt.Println(v.Name)
		}
	}

	return nil
}

func HandlerAgg(s *State, cmd Command) error {
	rss, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(rss)
	return nil
}

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("Add feed takes 2 args, name and url of the feed.")
	}

	curUser, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.Db.AddFeed(context.Background(), database.AddFeedParams{
		ID: uuid.New(),
		CreatedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		Name: sql.NullString{
			String: cmd.Args[0],
			Valid:  true,
		},
		Url: sql.NullString{
			String: cmd.Args[1],
			Valid:  true,
		},
		UserID: curUser.ID,
	})

	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
