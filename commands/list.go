package commands

import (
	"fmt"

	"github.com/IdanLoo/goconnect/util"
	"github.com/teris-io/cli"
)

// List is a command
var List cli.Command

func listAction(args []string, options map[string]string) int {
	for name := range util.Nodes {
		fmt.Println(name)
	}
	return 0
}

func init() {
	List = cli.NewCommand("list", "list all nodes").
		WithShortcut("l").
		WithOption(cli.NewOption("password", "display password").
			WithChar('p').
			WithType(cli.TypeBool)).
		WithAction(listAction)
}
