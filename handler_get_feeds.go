package main

import (
	"context"
	"fmt"
)

func handlerGetFeeds(state *state, cmd command) error {
	feeds, err := state.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		user, err := state.db.GetUser(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error fetching user: %w", err)
		}
		fmt.Printf("User: %s\n\n", user.Name)
	}
	return nil
}