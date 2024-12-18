package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
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
	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1, part2 = solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	//part2 = solve2(lines)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], lines[part2])

}

func solve1(s []string) (int, int) {
	res := math.MaxInt
	res2 := 0
	start, end := Point{0, 0}, Point{70, 70}

	grid := newGrid(70, 70)

	for step, coords := range s {
		printGrid(grid, nil)
		coords := strings.Split(coords, ",")
		x := ToInt(coords[0])
		y := ToInt(coords[1])
		grid[y][x] = '#'
		if step < 1024 {
			continue
		}
		score := findLowestScore(grid, start, end)
		debug = true
		debug = false
		if step == 1024 {
			dbg("%v: Score: %v", step, score)
			res = score
		}

		if score == -1 {
			dbg("%v: Score: %v", step, score)
			res2 = step
			break
		}

	}
	// res = findLowestScore(grid, start, end)

	//res, paths := findLowestScore(s, start, end)
	//res2 = len(paths[res])

	return res, res2
}

func printGrid(s [][]rune, d map[Point]int) {
	if !debug {
		return
	}
	for y, l := range s {
		for x, c := range l {
			p := Point{x, y}
			if _, ok := d[p]; ok {
				fmt.Print("O")
			} else {
				fmt.Printf("%c", c)
			}
		}
		println()
	}
	println()
}

func newGrid(w, h int) [][]rune {
	h++
	w++
	s := make([][]rune, h)
	for y := 0; y < h; y++ {
		s[y] = append(s[y], []rune(strings.Repeat(".", w))...)
	}
	dbg("Grid: %v", s)

	return s
}

type path struct {
	node Point
	cost int
}

func findLowestScore(g [][]rune, start, end Point) int {
	queue := []path{{node: start, cost: start.ManhattanDistance(end)}}
	distances := map[Point]int{
		start: 0,
	}

	for len(queue) > 0 {
		// Poor man's priority queue
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].cost < queue[j].cost
		})
		dbg("Q: %v", queue)

		cur := queue[0]
		queue = queue[1:]

		dbg("CUR: %v", cur)
		if cur.node == end {
			return distances[cur.node]
		}

		directions := []Point{
			P0.Right(),
			P0.Down(),
			P0.Left(),
			P0.Up(),
		}

		for _, dir := range directions {
			next := cur.node.Sum(dir)
			dbg(" -> %v", next)
			if next.X < 0 || next.X >= len(g) || next.Y < 0 || next.Y >= len(g[0]) {
				continue
			}

			if ch := g[next.X][next.Y]; ch == '#' {
				continue
			}

			nextDist := distances[cur.node] + 1
			if _, ok := distances[next]; !ok {
				distances[next] = nextDist
				queue = append(queue, path{
					node: next,
					cost: nextDist + next.ManhattanDistance(end),
				})
			}
		}
		dbg("DIST: %v", distances)
		// debug = true
		printGrid(g, distances)
		// debug = false

	}

	return -1
}

func findNodes(s []string, start, end rune) (Point, Point) {
	startPoint := Point{}
	endPoint := Point{}

	for y, l := range s {
		for x, c := range l {
			switch c {
			case start:
				startPoint = Point{x, y}
			case end:
				endPoint = Point{x, y}

			}
		}
	}

	return startPoint, endPoint
}

func solve2(s []string) int {
	res := 0

	return res
}
