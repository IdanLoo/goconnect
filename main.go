package main

import (
	"log"
	"os"

	"github.com/IdanLoo/goconnect/commands"
	"github.com/IdanLoo/goconnect/util"
	"github.com/teris-io/cli"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func mainAction(args []string, options map[string]string) int {
	name := args[0]
	target := util.Nodes[name]
	if target == nil {
		log.Fatal("Cannot find service named " + name)
	}

	session, err := util.Connect(target)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(fd, oldState)

	// excute command
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	termWidth, termHeight, err := terminal.GetSize(fd)
	if err != nil {
		panic(err)
	}

	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal
	if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
		log.Fatal(err)
	}

	if err := session.Shell(); err != nil {
		log.Fatal(err)
	}

	if err := session.Wait(); err != nil {
		log.Fatal(err)
	}

	return 0
}

func main() {
	app := cli.New("ssh tool").
		WithCommand(commands.List).
		WithCommand(commands.Add).
		WithCommand(commands.Rm).
		WithCommand(commands.Scp).
		WithArg(cli.NewArg("name", "name of service")).
		WithAction(mainAction)

	os.Exit(app.Run(os.Args, os.Stdout))
}
