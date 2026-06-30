package commands

import (
	"context"

	"github.com/GeorgievPlamen/rss-feed/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, c Command) error {
		curUser, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, c, curUser)
	}
}
