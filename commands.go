package main

import (
	"errors"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (cmds *commands) run(s *state, cmd command) error {
	c, ok := cmds.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}

	return c(s, cmd)
}

func (cmds *commands) register(name string, f func(*state, command) error) {
	cmds.registeredCommands[name] = f
}
