package main

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)


func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 || len(cmd.args) > 2 {
		return fmt.Errorf("usage: %v <time between fetches>", cmd.name)
	}
	timeBetweenFetches, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error parsing time: %w", err)
	}

	fmt.Printf("Collecting feeds every %v...\n", timeBetweenFetches)

	ticker := time.NewTicker(timeBetweenFetches)

	for ; ; <- ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("error fetching feed: %v", err)
	}
	scrapeFeed(s, feed)
}

func scrapeFeed(s *state, feed database.Feed) {
	log.Printf("Scraping feed: %s\n", feed.Url)
	feedUrl := feed.Url
	feedData, err := fetchFeed(context.Background(), feedUrl)
	if err != nil {
		log.Printf("error fetching feed: %v", err)
	}
	fmt.Printf("Feed: %v\n", feedData.Channel.Title)
	for _, item := range feedData.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt.Time = t
			publishedAt.Valid = true
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
		ID:         uuid.New(),
		FeedID:     feed.ID,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		Title:      item.Title,
		Description: sql.NullString{String: item.Description, Valid: true},
		Url:		item.Link,
		PublishedAt: publishedAt,
		})

		if err != nil {
			if strings.Contains(err.Error(), "unique constraint") {
				continue
			}
			log.Printf("error creating post: %v", err)
		}
	}
	fmt.Printf("Feed %s scraped, %v posts found\n", feed.Name, len(feedData.Channel.Item))

	
	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("error marking feed as fetched: %v", err)
	}
	fmt.Printf("Feed %s marked as fetched\n", feed.Name)
}