package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func handlerFollowing(state *state, cmd command, user database.User) error {
	user, err := state.db.GetUserByName(context.Background(), state.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error fetching user: %w", err)
	}
	follows, err := state.db.GetFeedFollows(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error fetching following: %w", err)
	}
	for _, follow := range follows {
		fmt.Printf("Following: %s\n", follow.FeedName)
		fmt.Printf("User: %s\n\n", follow.UserName)
	}
	return nil
}