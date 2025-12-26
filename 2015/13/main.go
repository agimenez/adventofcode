package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"maps"
	"math"
	"os"
	"slices"
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

// Undirected weighted graph
type Graph struct {
	edges map[string]map[string]int
}

func NewGraph() Graph {
	return Graph{
		edges: map[string]map[string]int{},
	}
}

func (g Graph) AddEdge(node1, node2 string, cost int) Graph {
	if g.edges[node1] == nil {
		g.edges[node1] = map[string]int{}
	}
	g.edges[node1][node2] = cost

	return g
}

func (g Graph) GetVertices() []string {
	s := slices.AppendSeq([]string{}, maps.Keys(g.edges))

	return s
}

func (g Graph) PathCost(seats []string) int {

	if len(seats) < 2 {
		return 0
	}

	res := g.edges[seats[0]][seats[1]]
	res += g.edges[seats[1]][seats[0]]
	dbg("  >> %v -> %v (%v)", seats[0], seats[1], res)
	dbg("  >> %v -> %v (%v)", seats[1], seats[0], res)

	res += g.PathCost(seats[1:])

	return res

}

func solve1(s []string) int {
	res := 0

	g := NewGraph()
	for _, str := range s {
		parts := strings.Fields(str)
		from := parts[0]
		to := parts[len(parts)-1]
		to = to[:len(to)-1] // Final dot
		cost := ToInt(parts[3])
		if parts[2] == "lose" {
			cost = -cost
		}

		g = g.AddEdge(from, to, cost)

		dbg("%v", g)
	}
	seats := g.GetVertices()
	dbg("SEATS: %v", seats)
	res = math.MinInt
	for p := range Permutations(seats) {
		// Close the cycle
		p = append(p, p[0])
		happiness := g.PathCost(p)
		dbg("%v -> %v (%v)", p, happiness, res)
		if happiness > res {
			res = happiness
		}
	}

	return res
}

func solve2(s []string) int {
	res := 0

	g := NewGraph()

	for _, str := range s {
		parts := strings.Fields(str)
		from := parts[0]
		to := parts[len(parts)-1]
		to = to[:len(to)-1] // Final dot
		cost := ToInt(parts[3])
		if parts[2] == "lose" {
			cost = -cost
		}

		g = g.AddEdge(from, to, cost)

		// Add myself
		g = g.AddEdge(from, "myself", 0)
		g = g.AddEdge("myself", from, 0)
	}

	seats := g.GetVertices()

	// Add myself
	res = math.MinInt
	for p := range Permutations(seats) {
		// Close the cycle
		p = append(p, p[0])
		happiness := g.PathCost(p)
		if happiness > res {
			res = happiness
		}
	}

	return res
}
