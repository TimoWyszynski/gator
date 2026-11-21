package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	cmdMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if s == nil {
		return errors.New("no state provided")
	}

	cmdFunc := c.cmdMap[cmd.name]
	err := cmdFunc(s, cmd)

	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdMap[name] = f
}
