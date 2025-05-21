package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(state *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: addfeed <feed_name> <feed_url>")
	}
	user, err := state.db.GetUserByName(context.Background(), state.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error fetching user: %w", err)
	}
	feedName := cmd.args[0]
	feedUrl := cmd.args[1]
	feed, err := state.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:   uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: feedName,
		Url:  feedUrl,
		UserID: user.ID,
	})
	fmt.Printf("Feed added: %s\n", feed.Name)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	return nil
}