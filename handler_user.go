package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if cmd.args == nil {
		return errors.New("missing username")
	}

	if len(cmd.args) != 1 {
		return errors.New("to many usernames")
	}

	err := s.cfg.SetUser(cmd.args[0])

	if err != nil {
		return err
	}

	fmt.Printf("Logged in as %v\n", cmd.args[0])

	return nil
}
