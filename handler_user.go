package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/TimoWyszynski/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if cmd.args == nil {
		return errors.New("missing username")
	}

	if len(cmd.args) != 1 {
		return errors.New("to many usernames")
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])

	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)

	if err != nil {
		return err
	}

	fmt.Printf("Logged in as %v\n", cmd.args[0])

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if cmd.args == nil {
		return errors.New("missing username")
	}

	if len(cmd.args) != 1 {
		return errors.New("to many usernames")
	}

	db := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.args[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := s.db.CreateUser(context.Background(), db)

	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)

	if err != nil {
		return err
	}

	return nil
}
