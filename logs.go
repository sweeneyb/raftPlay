package main

import (
	"github.com/mitchellh/cli"
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

func (c *ViewLog) Run(args []string) int {
	c.Ui.Output("run view")
	viewLogs()
	return 0
}

func (c *ViewLog) Help() string     { return "Dumps the logs" }
func (c *ViewLog) Synopsis() string { return c.Help() }
