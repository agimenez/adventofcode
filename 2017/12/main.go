package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"maps"
	"os"
	"strings"
	"time"
	. "github.com/agimenez/adventofcode/utils"
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

type Graph map[string]Set[string]

func (g Graph) AddEdge(n1, n2 string) {
	if g[n1] == nil {
		g[n1] = NewSet[string]()
	}

	g[n1].Add(n2)
}

// SpanTree returns a set of the visitable nodes from `start`
// BFS
func SpanTree(g Graph, start string) Set[string] {
	queue := []string{start}
	visited := Set[string]{}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		dbg("CUR: %v", cur)
		for next := range maps.Keys(g[cur]) {
			dbg("NEXT: %v", next)

			if !visited.Contains(next) {
				queue = append(queue, next)
				visited.Add(next)
			}
		}
	}

	return visited
}

func solve1(s []string) int {
	res := 0

	g := Graph{}
	for _, line := range s {
		parts := strings.Split(line, " <-> ")
		for _, dst := range strings.Split(parts[1], ", ") {
			g.AddEdge(parts[0], dst)
		}

	}
	dbg("%v", g)
	visited := SpanTree(g, "0")
	dbg("%v", visited)
	res = visited.Len()

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")

	return res
}
