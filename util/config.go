package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
)

var (
	home, _    = homedir.Dir()
	configFile = home + "/.goconnect"
)

func isConfigExist() bool {
	_, err := os.Stat(configFile)
	return err == nil || os.IsExist(err)
}

func createConfigFile() {
	os.Create(configFile)
}

func init() {
	if !isConfigExist() {
		createConfigFile()
	}
}

func linesOf(nodes NodeMap) []string {
	lines := make([]string, 0)
	for name, value := range nodes {
		lines = append(
			lines,
			fmt.Sprintf("%s %s %s %s %d", name, value.User, value.Password, value.Host, value.Port),
		)
	}
	return lines
}

// Lines is config line by line
func Lines() []string {
	fi, err := os.Open(configFile)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	defer fi.Close()

	lines := make([]string, 0)
	br := bufio.NewReader(fi)

	for {
		line, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		lines = append(lines, string(line))
	}

	return lines
}

// Save config
func Save(nodes NodeMap) {
	fi, err := os.OpenFile(configFile, os.O_WRONLY|os.O_TRUNC, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	fi.WriteString(strings.Join(linesOf(nodes), "\n") + "\n")
	fi.Sync()
}
