package main

import (
	"fmt"

	"github.com/Norrun/gator/internal/config"
	"github.com/Norrun/gator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	name string
	args []string
}
type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, exist := c.cmds[cmd.name]
	if !exist {
		return fmt.Errorf("unknown command")
	}
	err := handler(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}
