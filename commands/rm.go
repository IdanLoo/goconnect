package commands

import (
	"github.com/IdanLoo/goconnect/util"
	"github.com/teris-io/cli"
)

var rm cli.Command

func rmAction(args []string, options map[string]string) int {
	util.Nodes.Remove(args[0])
	return 0
}

func init() {
	rm = cli.NewCommand("rm", "remove service from list").
		WithArg(cli.NewArg("name", "name of service")).
		WithAction(rmAction)
}

// Rm get rm command
func Rm() cli.Command {
	return rm
}
