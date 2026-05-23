package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mayaanabbasi/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.FetchFeedByUrl(context.Background(), url)
	if err != nil {
		return nil
	}

	activeUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    activeUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed: ", feedFollow.FeedName)
	fmt.Println("User: ", feedFollow.UserName)

	return nil
}
