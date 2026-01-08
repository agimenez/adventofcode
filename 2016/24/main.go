package main

import (
	"cmp"
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
	"unicode"

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

type DuctMap struct {
	grid  Grid
	poi   map[Point]rune
	graph Graph

	// meh.
	poi0 Point
}

func (dm *DuctMap) genGraph() {
	d := Graph{}
	for pairs := range Combinations(dm.POIs(), 2) {
		cost := findShortestPath(*dm, pairs[0], pairs[1])
		dbg("DISTANCE %v: %d", pairs, cost)
		d.AddEdge(pairs[0], pairs[1], cost)
		dbg("")

	}
	dbg("%v", d)
	dm.graph = d
}

func ReadDuctMap(s []string) DuctMap {
	poi := map[Point]rune{}
	var p0 Point
	grid := NewGridFromStrFunc(s, func(p Point, r rune) rune {
		if unicode.IsDigit(r) {
			poi[p] = r
			if r == '0' {
				p0 = p
			}
		}
		return r
	})

	dm := DuctMap{
		grid: grid,
		poi:  poi,
		poi0: p0,
	}
	dm.genGraph()

	return dm
}

func (dm DuctMap) String() string {
	var b strings.Builder
	b.WriteString(dm.grid.String())
	b.WriteString("POIs:\n")
	for p, n := range dm.poi {
		b.WriteString(fmt.Sprintf("%c: %v\n", n, p))
	}
	b.WriteRune('\n')

	for p1, m := range dm.graph {
		for p2, cost := range m {
			b.WriteString(fmt.Sprintf("[%c - %c] %v - %v -> %d\n", dm.poi[p1], dm.poi[p2], p1, p2, cost))
		}
		b.WriteRune('\n')
	}

	return b.String()
}

type Graph map[Point]map[Point]int

func (g Graph) String() string {
	var b strings.Builder
	for p1, m := range g {
		for p2, cost := range m {
			b.WriteString(fmt.Sprintf("%v - %v -> %d\n", p1, p2, cost))
		}
		b.WriteRune('\n')
	}

	return b.String()
}

func (g Graph) AddEdge(node1, node2 Point, cost int) Graph {
	if g[node1] == nil {
		g[node1] = map[Point]int{}
	}
	g[node1][node2] = cost

	if g[node2] == nil {
		g[node2] = map[Point]int{}
	}

	g[node2][node1] = cost

	return g
}

type path struct {
	node Point
	cost int
}

func findShortestPath(dm DuctMap, start, end Point) int {
	queue := []path{{node: start, cost: start.ManhattanDistance(end)}}
	distances := map[Point]int{
		start: 0,
	}

	g := dm.grid

	for len(queue) > 0 {
		// Poor man's priority queue
		slices.SortFunc(queue, func(i, j path) int {
			return cmp.Compare(i.cost, j.cost)
		})

		dbg("Q: %v", queue)

		cur := queue[0]
		queue = queue[1:]

		dbg("CUR: %v", cur)
		if cur.node == end {
			return distances[end]
		}

		for next := range g.AdjacentPoints(cur.node, false) {
			dbg(" -> %v", next)
			if g.GetRune(next) == '#' {
				continue
			}

			if _, ok := distances[next]; !ok {
				nextDist := distances[cur.node] + 1
				distances[next] = nextDist
				queue = append(queue, path{
					node: next,
					cost: nextDist + next.ManhattanDistance(end),
				})
			}
		}
		dbg("DIST: %v", distances)

	}

	return -1
}

func (dm DuctMap) POIs() []Point {
	return slices.Collect(maps.Keys(dm.poi))
}

func (dm DuctMap) PathCost(pois []Point) int {

	if len(pois) < 2 {
		return 0
	}

	g := dm.graph
	from, to := pois[0], pois[1]

	res := g[from][to]
	dbg("  >> [%c to %c] %v -> %v (%v)", dm.poi[from], dm.poi[to], from, to, res)
	res += dm.PathCost(pois[1:])

	return res

}

func shortestRoute(dm DuctMap, circular bool) int {
	res := 0

	res = math.MaxInt
	for p := range Permutations(dm.POIs()) {
		if p[0] != dm.poi0 {
			continue
		}
		if circular {
			p = append(p, dm.poi0)
		}

		dbg("%v", p)
		cost := dm.PathCost(p)
		dbg("COST -> %v", cost)
		if cost < res {
			dbg("NEW MIN!!")
			res = cost
		}
	}

	return res
}

func solve1(s []string) int {
	res := 0

	dm := ReadDuctMap(s)
	dbg("%v", dm)

	res = shortestRoute(dm, false)

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")

	dm := ReadDuctMap(s)
	res = shortestRoute(dm, true)

	return res
}
