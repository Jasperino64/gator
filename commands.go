package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type Commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *Commands) register(name string, handler func(*state, command) error) {
	c.registeredCommands[name] = handler
}
func (c *Commands) run(s *state, cmd command) error {
	if handler, ok := c.registeredCommands[cmd.name]; ok {
		return handler(s, cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.name)
}
