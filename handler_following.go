package main

import (
	"context"
	"fmt"

	"github.com/mayaanabbasi/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedsFollowed, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range feedsFollowed {
		fmt.Println(feed.FeedName)
	}

	return nil
}
