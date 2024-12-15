package main

import (
	"flag"
	"fmt"
	"io"
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

type grid struct {
	m     map[Point]rune
	r     Point
	moves []rune

	w int
	h int
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
	part1 = solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 = solve2(lines)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}

func parseGrid(s []string) grid {
	g := grid{
		m:     map[Point]rune{},
		r:     P0,
		moves: []rune{},

		w: len(s[0]),
	}

	inInstructions := false
	for y, l := range s {
		if l == "" {
			inInstructions = true
			g.h = y
			continue
		}
		for x, c := range l {
			if !inInstructions {
				if c == '.' {
					continue
				}

				if c == '@' {
					g.r = Point{x, y}
					continue
				}

				g.m[Point{x, y}] = c
			} else {
				g.moves = append(g.moves, c)
			}
		}
	}

	return g
}

func (g grid) run() grid {
	for _, instr := range g.moves {
		switch instr {
		case '<':
			g = g.move(P0.Left())
		case '^':
			g = g.move(P0.Up())
		case '>':
			g = g.move(P0.Right())
		case 'v':
			g = g.move(P0.Down())
		}
		dbg("Instruction: %c\n", instr)
		g.print()
	}

	return g
}

func (g grid) move(dir Point) grid {
	next := g.r.Sum(dir)

	// Hit a wall: do nothing
	if g.m[next] == '#' {
		return g
	}

	// Empty space, just move
	if _, ok := g.m[next]; !ok {
		g.r = next
		return g
	}

	// There is at least a box
	lastBox := g.findLastBox(dir)
	nextAfterLast := lastBox.Sum(dir)
	// All against a wall, do nothing
	if g.m[nextAfterLast] == '#' {
		return g
	}

	// There is empty space after the stack of boxes
	g.m[nextAfterLast] = 'O'
	delete(g.m, next)
	g.r = next

	return g
}

func (g grid) findLastBox(dir Point) Point {
	cur := g.r
	for {
		next := cur.Sum(dir)
		if g.m[next] != 'O' {
			return cur
		}

		cur = next
	}
}

func (g grid) print() {
	if !debug {
		return
	}
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			p := Point{x, y}
			if p == g.r {
				fmt.Print("@")
				continue
			}
			ch, ok := g.m[p]

			if !ok {
				fmt.Print(".")
			} else {
				fmt.Printf("%c", ch)
			}
		}
		fmt.Println()
	}

}

func (g grid) GPSCoordinates() []Point {
	coords := []Point{}
	for coord, box := range g.m {
		if box == 'O' {
			coords = append(coords, coord)
		}
	}

	return coords
}

func solve1(s []string) int {
	res := 0
	g := parseGrid(s)
	g = g.run()

	coords := g.GPSCoordinates()
	for _, c := range coords {
		res += 100*c.Y + c.X
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
