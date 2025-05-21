package main

import (
	"context"
	"fmt"
)


func handlerAgg(s *state, cmd command) error {
	feedUrl := "https://www.wagslane.dev/index.xml"

	ctx := context.Background()
	feed, err := fetchFeed(ctx, feedUrl)
	if err != nil {
		return fmt.Errorf("error fetching feed: %v", err)
	}

	fmt.Printf("Feed: %v\n", feed)

	return nil
}