package main

import (
	"flag"
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

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
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

	g := parseGraph(lines)

	now = time.Now()
	part1 := solve1(g)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 := solve2(g)
	dur[1] = time.Since(now)

	return part1, part2, dur[0], dur[1]

}

type graph map[string][]string

func parseGraph(s []string) graph {
	g := graph{}

	for _, line := range s {
		parts := strings.Fields(line)

		from := parts[0][:len(parts[0])-1]
		to := parts[1:]

		g[from] = to

		dbg("%v", g)
	}

	return g
}

// DFS
func (g graph) findAllPaths(start, end string) [][]string {
	allPaths := [][]string{}

	queue := [][]string{[]string{start}}

	for len(queue) > 0 {
		curPath := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		curNode := curPath[len(curPath)-1]
		dbg("Current Path: %v", curPath)
		dbg("Queue: %v", queue)
		dbg("ALL PATHS: %v", allPaths)
		dbg("Checking node %v", curNode)

		if curNode == end {
			allPaths = append(allPaths, curPath)
		}

		for _, child := range g[curNode] {
			dbg(" >> CHILD %v", child)
			newPath := append([]string{}, curPath...)
			newPath = append(newPath, child)
			queue = append(queue, newPath)
		}

		dbg("Queue AFTER: %v", queue)
		dbg("")

	}

	return allPaths
}

func solve1(g graph) int {
	res := 0

	paths := g.findAllPaths("you", "out")
	dbg("FINAL PATHS: %v", paths)
	res = len(paths)

	return res
}

func solve2(g graph) int {
	res := 0

	return res
}
