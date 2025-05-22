package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(state *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: follow <url>")
	}
	feedUrl := cmd.args[0]
	feed, err := state.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	followingFeed, err := state.db.FollowFeed(context.Background(), database.FollowFeedParams{
		ID:		uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}
	fmt.Printf("Successfully followed feed: %s\n", followingFeed.FeedName)
	fmt.Printf("User: %s\n", followingFeed.UserName)
	return nil
}