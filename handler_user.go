package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mayaanabbasi/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User created successfully: ")
	printUser(user)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser((name))
	if err != nil {
		return err
	}

	fmt.Printf("User set")
	return nil
}

func printUser(user database.User) {
	fmt.Printf("* ID: 		%v\n", user.ID)
	fmt.Printf("* Name: 	%v\n", user.Name)
}
