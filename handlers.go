package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Norrun/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username.")
	}
	username := cmd.args[0]
	_, err := s.db.GetUserByName(context.Background(), username)
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

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("need at least one time argument")
	}
	durations := strings.Join(cmd.args, "")
	duration, err := time.ParseDuration(durations)
	if err != nil {
		return err
	}
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {

		err := scrapeFeeds(s)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		user, err := s.db.GetUserByName(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return err

		}
		return handler(s, c, user)
	}
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("the addfeed command expects 2 arguments name and url")
	}
	name := cmd.args[0]
	url := cmd.args[1]
	feed, err := s.db.AddFeed(context.Background(), database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	return nil
}

func handlerFeeds(s *state, _ command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, v := range feeds {
		fmt.Println(v.Name)
		fmt.Println(v.Url)
		fmt.Println(v.UserName)
	}
	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the follow command expects a single argument, feed url.")
	}
	url := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	feedFollowRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println(feedFollowRow.FeedName)
	fmt.Println(feedFollowRow.UserName)

	return nil
}

func handlerFollowing(s *state, _ command, _ database.User) error {
	following, err := s.db.GetFeedFollowsForUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}
	for _, v := range following {
		fmt.Println(v.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the unfollow command expects a single argument, feed url.")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	return s.db.RemoveUserFollow(context.Background(), database.RemoveUserFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

}
