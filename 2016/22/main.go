package main

import (
	"cmp"
	"flag"
	"fmt"
	"io"
	"log"
	"maps"
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

type Node struct {
	size, used, avail int
}

type Cluster struct {
	nodes     map[Point]Node
	maxPoint  Point
	emptyNode Point
}

func ReadCluster(df []string) Cluster {
	cluster := Cluster{
		nodes:    map[Point]Node{},
		maxPoint: P0,
	}

	for _, line := range df {
		// Filesystem              Size  Used  Avail  Use%
		// /dev/grid/node-x0-y0     85T   67T    18T   78%
		parts := strings.Fields(line)
		coords := strings.Split(parts[0], "-")
		p := NewPoint(ToInt(coords[1][1:]), ToInt(coords[2][1:]))

		cluster.maxPoint = p.Max(cluster.maxPoint)

		n := Node{
			size:  ToInt(parts[1][:len(parts[1])-1]),
			used:  ToInt(parts[2][:len(parts[2])-1]),
			avail: ToInt(parts[3][:len(parts[3])-1]),
		}
		cluster.nodes[p] = n

		if n.used == 0 {
			cluster.emptyNode = p
		}

	}

	return cluster
}

func (n Node) FitsInto(n2 Node) bool {
	return n.used != 0 && n.used <= n2.avail
}

func solve1(s []string) int {
	res := 0

	cluster := ReadCluster(s[2:])
	nodes := slices.Collect(maps.Values(cluster.nodes))
	for pair := range Combinations(nodes, 2) {
		if pair[0].FitsInto(pair[1]) || pair[1].FitsInto(pair[0]) {
			res++
		}
	}

	dbg("")
	return res
}

func cluster2Grid(c Cluster) Grid {
	g := NewGrid(c.maxPoint.X+1, c.maxPoint.Y+1)
	for p, n := range c.nodes {
		if n.used == 0 {
			g.SetRune(p, '_')
		} else if n.size > 200 {
			g.SetRune(p, '#')

		} else if p == NewPoint(c.maxPoint.X, 0) {
			g.SetRune(p, 'G')
		} else if p == P0 {
			g.SetRune(p, '*')
		} else {
			g.SetRune(p, '.')
		}
	}

	return g
}

type path struct {
	node Point
	cost int
	path []Point
}

func findShortestPath(g Grid, start, end Point) []Point {
	queue := []path{{
		node: start,
		cost: start.ManhattanDistance(end),
		path: []Point{start},
	}}

	distances := map[Point]int{
		start: 0,
	}

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
			// Do not include the goal in the path
			return cur.path[1:]

		}

		for next := range g.AdjacentPoints(cur.node, false) {
			dbg(" -> %v", next)
			if g.GetRune(next) == '#' {
				continue
			}

			if _, ok := distances[next]; !ok {
				nextDist := distances[cur.node] + 1
				distances[next] = nextDist

				nextPath := slices.Clone(cur.path)
				nextPath = append(nextPath, next)

				queue = append(queue, path{
					node: next,
					cost: nextDist + next.ManhattanDistance(end),
					path: nextPath,
				})
			}
		}
		dbg("DIST: %v", distances)

	}

	dbg("GRID:\n%v", g)
	return []Point{}
}

func printGrid(g Grid, special map[Point]rune, path []Point) {
	g2 := g.Clone()

	for _, p := range path {
		g2.SetRune(p, '*')
	}

	for p, r := range special {
		g2.SetRune(p, r)
	}
	fmt.Printf("%v\n", g2)
}

// I got desperate with this, and looks like most people on Reddit, including Eric
// mentioned to just print the grid, and solve it manually. FML
// Update: found a somewhat "simple" BFS-based algo in Reddit that can solve this specific
// puzzle
func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")

	cluster := ReadCluster(s[2:])
	grid := cluster2Grid(cluster)
	fmt.Println(grid)

	// First find the path from the empty node to the goal
	g := grid.Clone()
	empty := cluster.emptyNode
	goal := NewPoint(cluster.maxPoint.X, 0)
	start := NewPoint(0, 0)
	pathGoalStart := findShortestPath(g, goal, start)

	fmt.Println(pathGoalStart)
	for goal != start {
		nextG := pathGoalStart[0]
		pathGoalStart = pathGoalStart[1:]
		fmt.Printf("Empty: %v\n", empty)
		fmt.Printf("P0: %v\n", nextG)
		fmt.Printf("Goal(obst): %v\n", goal)

		g.SetRune(goal, '#')
		path := findShortestPath(g, empty, nextG)
		fmt.Printf("Path: %v\n", path)
		g.SetRune(goal, '.')
		printGrid(g, map[Point]rune{
			empty: '_',
			goal:  '#',
		}, path)
		empty = goal
		g.SetRune(goal, '.')
		goal = path[len(path)-1]
		printGrid(g, map[Point]rune{
			empty: '_',
			goal:  '#',
		}, path)

		res += len(path) + 1
	}
	fmt.Printf("%v\n", g)

	return res
}
