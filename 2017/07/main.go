package main

import (
	"flag"
	"fmt"
	// . "github.com/agimenez/adventofcode/utils"
	"io"
	"log"
	"os"
	"strings"
	"time"
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

func (n *Node) ChildSums() int {
	// dbg(">> %s (%d)", n.name, n.weight)
	sum := n.weight
	for _, child := range n.children {
		sum += child.ChildSums()
		// dbg("   >> %s -> %d", child.name, child.ChildSums())
	}

	return sum
}

func findUnbalance(n *Node) (int, int) {
	res := 0

	sums := map[int][]*Node{}

	dbg(">> %s (%d)", n.name, n.weight)
	for _, child := range n.children {
		sum := child.ChildSums()
		sums[sum] = append(sums[sum], child)

		dbg("   >> %s (%d)", child.name, sum)
	}

	var childSums [2]int // good and bad sum
	var unbalancedChild *Node
	for sum, children := range sums {
		if len(children) == 1 {
			dbg("Unbalance detected for sum %d (%v)", sum, children[0])
			// First one to detect an offset with its children is the one to rebalance
			unbalancedChild = children[0]
			childSums[1] = sum
		} else {
			childSums[0] = sum
		}
	}

	// If it's unbalanced, go deep to the next unbalanced child
	if unbalancedChild != nil {
		offset := childSums[0] - childSums[1]
		dbg("UNBalanced (%s): going deeper -> %s, offset: %d", n.name, unbalancedChild.name, offset)
		orig, _ := findUnbalance(unbalancedChild)
		return orig, offset
	}

	res = n.weight
	dbg("Balanced (%s): returning %d", n.name, res)

	return res, 0
}

func solve1(s []string) int {
	res := 0

	t := parseTree(s)
	// dbg("%v", t)
	fmt.Println(t.name)

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")
	t := parseTree(s)
	res, offset := findUnbalance(t)
	res += offset

	return res
}
