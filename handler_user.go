package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.name)
	}
	username := cmd.args[0]

	// Check if user exists
	existingUser, err := s.db.GetUserByName(context.Background(), username)
	if err != nil {
		return fmt.Errorf("user %s does not exist: %v", username, err)
	}
	if existingUser.ID == uuid.Nil {
		return fmt.Errorf("user %s does not exist", username)
	}

	err = s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
		
	}
	fmt.Printf("User set to %s\n", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.name)
	}
	username := cmd.args[0]

	// Check if user already exists
	// existingUser, err := s.db.GetUser(context.Background(), username)
	// if err != nil {
	// 	return fmt.Errorf("error checking user existence: %v", err)
	// }
	// if existingUser.ID != uuid.Nil {
	// 	return fmt.Errorf("user %s already exists with ID %s", username, existingUser.ID)
	// }
	
	// Create new user
	newUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		return fmt.Errorf("error creating user: %s", username)
	}
	// fmt.Printf("User %s created with ID %s\n", newUser.Name, newUser.ID)

	err = s.config.SetUser(newUser.Name)
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
		
	}
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %v", err)
	}
	
	for _, user := range users {
		if s.config.CurrentUserName == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}