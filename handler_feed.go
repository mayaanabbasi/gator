package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mayaanabbasi/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      name,
		Url:       url,
	})
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("================================")

	return nil
}

func handlerPrintFeeds(s *state, cmd command) error {
	feeds, err := s.db.FetchFeeds(context.Background())
	if err != nil {
		return err
	}

	feedUsers, err := s.db.FetchFeedUsers(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		var userName string
		for _, user := range feedUsers {
			if user.ID == feed.UserID {
				userName = user.Name
				break
			}
		}
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(userName)
		fmt.Println("=============================")
	}

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
