package main

import (
	"context"
	"gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User)  error) func (s *state, cmd command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUserByName(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}