package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(state *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: addfeed <feed_name> <feed_url>")
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
	return nil
}