package main

import (
	"bufio"
	"log"
	"os"
)

const (
	debug = false
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

type Point struct {
	x, y int
}

type slope struct {
	width  int
	height int
	grid   map[Point]struct{}
}

func newSlope(s *bufio.Scanner) *slope {
	g := &slope{
		width:  0,
		height: 0,
		grid:   make(map[Point]struct{}),
	}

	for s.Scan() {
		l := s.Text()
		dbg("Line: '%v'\n", l)
		g.width = len(l)

		for x, square := range l {
			if square == '#' {
				g.grid[Point{x, g.height}] = struct{}{}

			}
		}
		g.height++
		dbg("After parsing: W: %v, H: %v\n", g.width, g.height)

	}

	return g
}

func (s *slope) countTrees(v Point) int {
	cur := Point{0, 0}
	count := 0
	for cur.y < s.height {
		if _, ok := s.grid[cur]; ok {
			dbg("Tree at %v!\n", cur)
			count++
		}

		cur.x = (cur.x + v.x) % s.width
		cur.y = (cur.y + v.y)
		dbg("New cur: %v\n", cur)
	}

	return count
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	g := newSlope(s)
	part1 := g.countTrees(Point{3, 1})

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: \n")

}
