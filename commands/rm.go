package commands

import (
	"github.com/IdanLoo/goconnect/util"
	"github.com/teris-io/cli"
)

// Rm is a command
var Rm cli.Command

func rmAction(args []string, options map[string]string) int {
	util.Nodes.Remove(args[0])
	return 0
}

func init() {
	Rm = cli.NewCommand("rm", "remove service from list").
		WithArg(cli.NewArg("name", "name of service")).
		WithAction(rmAction)
}
