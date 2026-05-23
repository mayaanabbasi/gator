package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	activeUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedsFollowed, err := s.db.GetFeedFollowsForUser(context.Background(), activeUser.ID)
	if err != nil {
		return err
	}

	for _, feed := range feedsFollowed {
		fmt.Println(feed.FeedName)
	}

	return nil
}
