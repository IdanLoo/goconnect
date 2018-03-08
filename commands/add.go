package commands

import (
	"fmt"

	"github.com/IdanLoo/goconnect/util"
	"github.com/teris-io/cli"
)

var add cli.Command

func addAction(args []string, options map[string]string) int {
	var (
		name     = args[0]
		user     string
		password string
		host     string
		port     int
	)
	print("user: ")
	fmt.Scan(&user)

	print("password: ")
	fmt.Scan(&password)

	print("host: ")
	fmt.Scan(&host)

	print("port: ")
	fmt.Scan(&port)

	util.Nodes.Append(
		name,
		util.NewNode(user, password, host, port),
	)
	fmt.Printf("Add %s success!", name)
	return 0
}

func init() {
	add = cli.NewCommand("add", "add service to list").
		WithShortcut("a").
		WithArg(cli.NewArg("name", "any name easy to remember")).
		WithAction(addAction)
}

// Add get add command
func Add() cli.Command {
	return add
}
