package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
	// . "github.com/agimenez/adventofcode/utils"
)

var (
	debug bool
)

func dbg(f string, v ...interface{}) {
	if debug {
		fmt.Printf(f+"\n", v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
}

func main() {
	flag.Parse()

	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	part1, part2, dur1, dur2 := solve(lines)
	log.Printf("Part 1 (%v): %v\n", dur1, part1)
	log.Printf("Part 2 (%v): %v\n", dur2, part2)

}

func solve(lines []string) (int, int, time.Duration, time.Duration) {
	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 := solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 := solve2(lines)
	dur[1] = time.Since(now)

	return part1, part2, dur[0], dur[1]

}

type Node struct {
	name   string
	weight int

	parent   *Node
	children []*Node
}

func (n *Node) Root() *Node {
	for ; n.parent != nil; n = n.parent {
	}

	return n
}

func parseTree(s []string) *Node {
	edges := map[string]*Node{}
	vertices := map[string][]string{}

	for _, line := range s {
		dbg("Line: %v", line)
		dbg("Edges: %v", edges)
		dbg("vertices: %v", vertices)
		parts := strings.Split(line, " -> ")
		var name string
		var value int
		fmt.Sscanf(parts[0], "%s (%d)", &name, &value)

		edges[name] = &Node{
			name:   name,
			weight: value,
		}

		if len(parts) == 1 {
			continue
		}

		children := strings.Split(parts[1], ", ")
		for _, c := range children {
			vertices[name] = append(vertices[name], c)
		}

	}

	var t *Node
	for name, node := range edges {
		for _, child := range vertices[name] {
			ch := edges[child]
			ch.parent = node
			node.children = append(node.children, ch)
		}

		if node.parent == nil {
			t = node
		}
	}

	t = t.Root()

	return t
}

func solve1(s []string) int {
	res := 0

	t := parseTree(s)
	dbg("%v", t)
	fmt.Println(t.name)

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")

	return res
}
