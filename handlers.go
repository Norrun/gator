package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Norrun/gator/internal/database"
	"github.com/Norrun/gator/internal/rss"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username.")
	}
	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("User %s was not found", username)
	}
	err = s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("User is set to %s\n", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the register handler expects a single argument, the username.")
	}
	username := cmd.args[0]
	usr, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username})
	if err != nil {
		return err
	}
	err = s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("user %s was created\n", username)
	fmt.Println(usr)
	return nil
}

func handlerReset(s *state, _ command) error {
	return s.db.Reset(context.Background())
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Name == s.config.CurrentUserName {
			fmt.Printf("* %s (current) \n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}

func handlerAgg(_ *state, _ command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(*feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("the addfeed command expects 2 arguments name and url")
	}
	name := cmd.args[0]
	url := cmd.args[1]
	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}
	_, err = s.db.AddFeed(context.Background(), database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	return nil
}
