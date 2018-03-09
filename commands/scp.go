package commands

import (
	"github.com/IdanLoo/goconnect/util"
	"github.com/teris-io/cli"
	"github.com/tmc/scp"
)

// Scp is a command
var Scp cli.Command

func scpAction(args []string, options map[string]string) int {
	target := util.Nodes[args[0]]

	session, err := util.Connect(target)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	if err := scp.CopyPath(args[1], args[2], session); err != nil {
		panic(err)
	}

	println("success!")
	return 0
}

func init() {
	Scp = cli.NewCommand("upload", "upload file to service").
		WithArg(cli.NewArg("name", "name of service")).
		WithArg(cli.NewArg("file-path", "file path")).
		WithArg(cli.NewArg("destination-path", "destination path")).
		WithAction(scpAction)
}
