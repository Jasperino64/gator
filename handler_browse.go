package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"strconv"
)

func handlerBrowse(state *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		var err error
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}
	posts, err := state.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error fetching posts: %w", err)
	}
	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		if post.Description.Valid {
			fmt.Printf("Description: %s\n", post.Description.String)
		} else {
			fmt.Println("Description: Not available")
		}
		fmt.Printf("URL: %s\n", post.Url)
		if post.PublishedAt.Valid {
			fmt.Printf("Published: %s\n", post.PublishedAt.Time)
		} else {
			fmt.Println("Published: Not available")
		}
		fmt.Printf("Published: %v\n", post.PublishedAt.Time)
		fmt.Printf("Created At: %s\n", post.CreatedAt)
		fmt.Printf("Updated At: %s\n", post.UpdatedAt)
		fmt.Println()
	}
	return nil
}