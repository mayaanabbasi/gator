package main

import (
	"context"
	"fmt"
)

func handlerFetchFeed(s *state, cmd command) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Print(rssFeed)
	return nil
}
