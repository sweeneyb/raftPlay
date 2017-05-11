package main

import (
	"github.com/mitchellh/cli"
)

type ConfCommand struct {
	Ui cli.Ui
}

func (c *ConfCommand) Run(_ []string) int {
	c.Ui.Output("Conf run")
	return 0
}

func (c *ConfCommand) Help() string {
	return "Conf [view]"
}

func (c *ConfCommand) Synopsis() string {
	return "Interact with the Conf"
}
