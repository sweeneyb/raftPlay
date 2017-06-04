package main

import (
	"fmt"
	"github.com/hashicorp/raft-boltdb"
	"github.com/mitchellh/cli"
	"log"
)

type ConfCommand struct {
	Ui cli.Ui
}

func (c *ConfCommand) Run(args []string) int {
	c.Ui.Output("Conf run")
	store, err := raftboltdb.NewBoltStore("raft.db")
	if err != nil {
		log.Fatal(err)
	}
	if store != nil {
		log.Print("opened the bolt store")
	}
	var key string = args[0]
	c.Ui.Output(fmt.Sprintf("searching for %s", key))
	message, err := store.Get([]byte(key))
	if err != nil {
		c.Ui.Output(fmt.Sprintf("%s", err))
		return 1
	}
	c.Ui.Output(fmt.Sprintf("%s", message))
	return 0
}

func (c *ConfCommand) Help() string {
	return "Conf [view]"
}

func (c *ConfCommand) Synopsis() string {
	return "Interact with the Conf"
}
