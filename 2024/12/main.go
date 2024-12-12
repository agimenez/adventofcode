package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	. "github.com/agimenez/adventofcode/utils"
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

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	m := lines2map(lines)
	part1 = solve1(m)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 = solve2(m)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}

func lines2map(s []string) map[Point]string {
	m := map[Point]string{}

	for y := range s {
		for x, c := range s[y] {
			m[Point{x, y}] = string(c)
		}
	}

	return m
}

type region struct {
	area  int
	perim int
}

func solve1(m map[Point]string) int {
	res := 0

	getAllAdjacentFilter := func(p Point) bool {
		if _, ok := m[p]; ok {
			return true
		}

		return false
	}

	queue := []Point{{0, 0}}
	visited := map[Point]bool{}
	for len(queue) > 0 {
		//dbg("Queue: %v", queue)
		cur := queue[0]
		queue = queue[1:]

		if _, ok := visited[cur]; ok {
			continue
		}
		visited[cur] = true

		plant := m[cur]

		// Get the whole plot
		plot := getPlot(m, cur, plant)
		area := len(plot)
		perimeter := 0
		// Calculate the perimeter by checking neighbours (including off map) that
		// are a different type of plant
		for _, p := range plot {
			visited[p] = true
			adj := getAdjacent(m, p, getAllAdjacentFilter)
			queue = append(queue, adj...)

			// Now get the non-plot neighbours for the perimeter
			neigh := getAdjacent(m, p, func(dir Point) bool {
				if neighbour, ok := m[dir]; !ok || neighbour != plant {
					return true
				}

				return false

			})
			dbg("Plot %v (plant %v) neigh: %v", plant, p, neigh)
			perimeter += len(neigh)
		}

		dbg("Plant %v (%v) plot: %v", plant, cur, plot)
		dbg("Plant %v (%v) neig: %v", plant, cur, perimeter)
		res += area * perimeter

		adj := getAdjacent(m, cur, getAllAdjacentFilter)
		queue = append(queue, adj...)
	}

	return res
}

func getPlot(m map[Point]string, start Point, plant string) []Point {
	queue := []Point{start}
	visited := map[Point]bool{}
	for len(queue) > 0 {
		//dbg("getPlot (%v) queue: %v, visited: %v", plant, queue, visited)
		cur := queue[0]
		queue = queue[1:]

		//dbg("getPlot (%v) -> checking %v", plant, cur)

		if _, ok := visited[cur]; ok {
			continue
		}

		visited[cur] = true
		sameType := getAdjacent(m, cur, func(dir Point) bool {
			if neighbour, ok := m[dir]; ok && neighbour == plant {
				return true
			}

			return false
		})
		//dbg("getPlot (%v) ---> sametype: %v", plant, sameType)

		queue = append(queue, sameType...)
	}
	plot := []Point{}
	for p := range visited {
		plot = append(plot, p)
	}

	return plot
}

func getAdjacent(m map[Point]string, p Point, filter func(dir Point) bool) []Point {
	adj := []Point{}

	candidates := []Point{
		p.Up(),
		p.Right(),
		p.Down(),
		p.Left(),
	}

	for _, c := range candidates {
		if filter(c) {
			adj = append(adj, c)
		}
	}

	return adj
}

func solve2(m map[Point]string) int {
	res := 0

	return res
}
