package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	debug bool
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
}

type node struct {
	name     string
	size     int
	children []*node
	parent   *node
}

func printTree(n *node, level int) {
	padding := strings.Repeat(" ", level)
	fmt.Printf("%s - %s (%d)\n", padding, n.name, n.size)
	for _, child := range n.children {
		printTree(child, level+3)
	}

}

func updateParentsSizes(n node) {
	parent := n.parent

	for parent != nil {
		parent.size += n.size
		parent = parent.parent
	}
}

func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	var root node
	cur := &root
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			dbg("cmd: %v", line)
			// command
			if parts[1] == "cd" {
				if parts[2] == ".." {
					cur = cur.parent
				} else {
					// New directory, create node
					n := node{name: parts[2]}
					n.parent = cur
					cur.children = append(cur.children, &n)
					cur = &n
				}
			}

		} else { // output of "ls"
			// If listing a dir, do nothing. Will handle it when cd'ing into it
			if parts[0] == "dir" {
				continue
			}

			size, _ := strconv.Atoi(parts[0])
			n := node{
				name:   parts[1],
				size:   size,
				parent: cur,
			}
			cur.children = append(cur.children, &n)
			updateParentsSizes(n)

		}
		//printTree(&root, 0)

	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
