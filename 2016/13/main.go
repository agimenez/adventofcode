package main

import (
	"bytes"
	"cmp"
	"flag"
	"fmt"
	"io"
	"iter"
	"log"
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

type path struct {
	node Point
	cost int
}

type InfiniteGrid struct {
	Grid

	maxX, maxY int
}

func NewInfiniteGrid() InfiniteGrid {
	g := InfiniteGrid{
		Grid: NewGrid(0, 0),
	}

	return g
}

func (g InfiniteGrid) String() string {
	var b bytes.Buffer

	b.WriteString("  0123456789\n")
	for y := 0; y < g.maxY; y++ {
		b.WriteString(fmt.Sprintf("%v ", y))

		for x := 0; x < g.maxX; x++ {
			r := g.GetRune(NewPoint(x, y))
			if r != '.' && r != '#' {
				r = '?'
			}

			b.WriteRune(r)
		}
		b.WriteRune('\n')
	}
	b.WriteRune('\n')

	return b.String()

}

func (g *InfiniteGrid) SetRune(p Point, r rune) {
	g.Grid.SetRune(p, r)
	g.maxX = Max(g.maxX, p.X)
	g.maxY = Max(g.maxY, p.Y)
}

func (g *InfiniteGrid) isWall(p Point, fav int) bool {
	r := g.GetRune(p)
	if r == '.' {
		return false
	}

	if r == '#' {
		return true
	}

	x := p.X
	y := p.Y

	tmp := x*x + 3*x + 2*x*y + y + y*y
	tmp += fav

	bin := fmt.Sprintf("%b", tmp)
	numOnes := strings.Count(bin, "1")
	if numOnes%2 == 0 {
		g.SetRune(p, '.')
		return false
	}

	g.SetRune(p, '#')
	return true
}

func (g *InfiniteGrid) AdjacentPoints(p Point, fav int) iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for p := range p.Adjacent(false) {
			if p.X < 0 || p.Y < 0 {
				continue
			}

			if g.isWall(p, fav) {
				continue
			}

			if !yield(p) {
				return
			}
		}
	}
}

func findShortestPath(start, end Point, fav int, maxDist int) map[Point]int {
	queue := []path{{node: start, cost: start.ManhattanDistance(end)}}
	distances := map[Point]int{
		start: 0,
	}

	g := NewInfiniteGrid()
	g.SetRune(start, '.')

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
			dbg("GRID:\n%v", g)
			return distances
		}

		for next := range g.AdjacentPoints(cur.node, fav) {
			dbg(" -> %v", next)

			nextDist := distances[cur.node] + 1
			if maxDist != -1 && nextDist > maxDist {
				continue
			}

			if _, ok := distances[next]; !ok {
				distances[next] = nextDist
				queue = append(queue, path{
					node: next,
					cost: nextDist + next.ManhattanDistance(end),
				})
			}
		}
		dbg("DIST: %v", distances)

	}

	dbg("GRID:\n%v", g)
	return distances
}

func TestRun() {
	start := NewPoint(1, 1)
	end := NewPoint(7, 4)
	fav := 10

	shortest := findShortestPath(start, end, fav, -1)
	dbg("Shortest test run: %v", shortest)
}

func solve1(s []string) int {
	res := 0

	if debug {
		TestRun()
	}

	start := NewPoint(1, 1)
	end := NewPoint(31, 39)
	dist := findShortestPath(start, end, 1358, -1)
	res = dist[end]

	return res
}

func solve2(s []string) int {
	res := 0

	start := NewPoint(1, 1)
	end := NewPoint(31, 39)
	dist := findShortestPath(start, end, 1358, 50)

	for _, v := range dist {
		if v <= 50 {
			res++
		}
	}

	return res
}
