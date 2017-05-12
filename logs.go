package main

import (
	"github.com/mitchellh/cli"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"log"
)

type LogCommand struct {
	Ui cli.Ui
}

func (c *LogCommand) Run(args []string) int {
	//	c.Ui.Output("Log run")

	logC := cli.NewCLI("log subcommand", "")
	logC.Args = args

	logC.Commands = map[string]cli.CommandFactory{
		"view": func() (cli.Command, error) {
			return &ViewLog{Ui: c.Ui}, nil
		},
		"add": func() (cli.Command, error) {
			return &AddPeer{Ui: c.Ui}, nil
		},
		
		
	}
	if exitStatus, err := logC.Run(); err != nil {
		c.Ui.Error(err.Error())
		return exitStatus
	} else {
		return exitStatus
	}

}

func (c *LogCommand) Help() string {
	return "Log [view|appendAddPeer|appendRemovePeer]"
}

func (c *LogCommand) Synopsis() string {
	return "Interact with the Logs"
}

type ViewLog struct {
	Ui cli.Ui
}

func doNothing(store *raftboltdb.BoltStore, log *raft.Log) error {
     return nil
}

func (c *ViewLog) Run(args []string) int {
	c.Ui.Output("run view")
	viewLogs(doNothing)
	return 0
}

func (c *ViewLog) Help() string     { return "Dumps the logs" }
func (c *ViewLog) Synopsis() string { return c.Help() }

type AddPeer struct { Ui cli.Ui }
func (c *AddPeer) Help() string {return "append an AddPeer record to the end of the log"}
func (c *AddPeer) Synopsis() string {return c.Help()}
func (c *AddPeer) Run(args []string) int {
     viewLogs(addPeer)
     return 0
}

func addPeer(store *raftboltdb.BoltStore, raftLog *raft.Log) error {
     raftLog.Type = raft.LogAddPeer
     raftLog.Data = encodePeers([]string{"127.0.0.1"})
     log.Printf("to be appended: %s\n", raftLog)
     return nil
}