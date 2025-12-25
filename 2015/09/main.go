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

	if g.edges[node2] == nil {
		g.edges[node2] = map[string]int{}
	}
	g.edges[node2][node1] = cost

	return g
}

func (g Graph) GetVertices() []string {
	s := slices.AppendSeq([]string{}, maps.Keys(g.edges))

	return s
}

func (g Graph) FloydWarshall() map[string]map[string]int {
	vertices := g.GetVertices()

	dist := g.edges
	for _, k := range vertices {
		for _, i := range vertices {
			for _, j := range vertices {
				var dist_ik int
				var dist_kj int
				var ok bool

				if dist_ik, ok = dist[i][k]; !ok {
					continue
				}

				if dist_kj, ok = dist[k][j]; !ok {
					continue
				}

				if dist[i][j] > dist_ik+dist_kj {
					dist[i][j] = dist_ik + dist_kj
				}

			}
		}
	}

	return dist
}

func (g Graph) PathCost(cities []string) int {

	if len(cities) < 2 {
		return 0
	}

	res := g.edges[cities[0]][cities[1]]
	dbg("  >> %v -> %v (%v)", cities[0], cities[1], res)
	res += g.PathCost(cities[1:])

	return res

}

func solve1(s []string) int {
	res := 0

	g := NewGraph()
	for _, str := range s {
		parts := strings.Split(str, " = ")
		cost := ToInt(parts[1])

		cities := strings.Split(parts[0], " to ")

		g = g.AddEdge(cities[0], cities[1], cost)

		dbg("%v", g)
	}
	cities := g.GetVertices()
	dbg("CITIES: %v", cities)
	res = math.MaxInt
	for p := range Permutations(cities) {
		cost := g.PathCost(p)
		dbg("%v -> %v (%v)", p, cost, res)
		if cost < res {
			res = cost
		}
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
