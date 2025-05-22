package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func handlerUnfollow(state *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: unfollow <url>")
	}
	feedUrl := cmd.args[0]
	feed, err := state.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	err = state.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,})
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %w", err)
	}
	fmt.Printf("Successfully unfollowed feed: %s\n", feedUrl)
	return nil
}