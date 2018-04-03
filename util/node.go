package util

import (
	"log"
	"strconv"
	"strings"
)

// Node is a remote server
type Node struct {
	User     string
	Password string
	Host     string
	Port     int
}

// NodeMap is alias of map[string]*Node
type NodeMap map[string]*Node

// Nodes is all node in configuration file
var Nodes = make(NodeMap)

// NewNode is constructor of Node
func NewNode(user, password, host string, port int) *Node {
	return &Node{user, password, host, port}
}

func init() {
	for _, line := range Lines() {
		items := strings.Split(line, " ")
		port, err := strconv.Atoi(items[4])
		if err != nil {
			log.Fatal(err)
		}
		Nodes[items[0]] = NewNode(items[1], items[2], items[3], port)
	}
}

// Append node to nodelist
func (n NodeMap) Append(name string, node *Node) {
	n[name] = node
	Save(n)
}

// Remove node from nodelist
func (n NodeMap) Remove(name string) {
	delete(n, name)
	Save(n)
}
