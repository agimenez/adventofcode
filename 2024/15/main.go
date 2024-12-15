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
		dbg("Instruction: %c\n", instr)
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

	if dir == P0.Left() || dir == P0.Right() {
		g = g.moveHorizontal(dir)
		return g
	}

	if dir == P0.Up() || dir == P0.Down() {
		g = g.moveVertical(dir)
		return g
	}

	return g
}

func (g grid) moveHorizontal(dir Point) grid {
	cur := g.r
	moves := []Point{}
	g.print()
	for {
		next := cur.Sum(dir)
		if g.m[next] == '#' {
			return g
		}

		moves = append(moves, next)
		//dbg("moves: %v", moves)
		if _, ok := g.m[next]; !ok {
			//dbg(" Found empty space!")
			for i := len(moves) - 1; i > 0; i-- {
				g.m[moves[i]] = g.m[moves[i-1]]
			}
			delete(g.m, moves[0])
			g.r = g.r.Sum(dir)

			return g
		}
		// g.print()
		cur = next
	}
}

func (g grid) moveVertical(dir Point) grid {
	cur := g.r
	moves := []Point{}
	g.print()

	// At the new Y level, leftmost and rightmost position with a box
	leftX := cur.X
	rightX := cur.X
	for {
		next := cur.Sum(dir)

		leftmost := Point{leftX, next.Y}
		rightmost := Point{rightX, next.Y}

		if g.m[leftmost] == ']' {
			leftX--
		}
		if g.m[rightmost] == '[' {
			rightX++
		}

		// Once in the next level, the number of boxes can be less than the previous level, so adjust this level's leftmost/rightmost
		// 0: []..
		// 1: [][]
		// 2: .[].
		// 3: .@..
		// After pushing level 2, in level 3 leftmost previous == 0, rightmost == 3
		// So, coming from level 3, level 1 needs to be leftmost == 0, rightmost == *1*
		for ; g.m[Point{leftX, next.Y}] == 0; leftX++ {
		}
		for ; g.m[Point{rightX, next.Y}] == 0; rightX-- {
		}

		dbg("leftmost: %v, rightmost: %v", leftX, rightX)

		// Scan all the boxes within the found limits above
		allEmpty := true
		for x := leftX; x <= rightX; x++ {
			nextBox := Point{x, next.Y}
			ch := g.m[nextBox]

			// Can't move
			if ch == '#' {
				return g
			}

			if ch == 'O' || ch == '[' || ch == ']' {
				moves = append(moves, nextBox)
				allEmpty = false
			}
		}
		dbg("Moves: %v", moves)

		if allEmpty {
			for i := len(moves) - 1; i >= 0; i-- {
				box := moves[i]
				g.m[box.Sum(dir)] = g.m[box]
				delete(g.m, box)
			}

			g.r = g.r.Sum(dir)
			return g
		}

		cur = next
		// g.print()
	}

	return g
}

func (g grid) findLastBox(dir Point) Point {
	cur := g.r
	for {
		next := cur.Sum(dir)
		v := g.m[next]
		if v != 'O' && v != '[' && v != ']' {
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
		if box == 'O' || box == '[' {
			coords = append(coords, coord)
		}
	}

	return coords
}

func solve1(s []string) int {
	res := 0
	g := parseGrid(s)
	g.print()
	g = g.run()

	coords := g.GPSCoordinates()
	for _, c := range coords {
		res += 100*c.Y + c.X
	}

	return res
}

func solve2(s []string) int {
	res := 0
	ag := augmentGrid(s)
	g := parseGrid(ag)
	// g.print()
	g.run()
	// g.print()

	coords := g.GPSCoordinates()
	for _, c := range coords {
		res += 100*c.Y + c.X
	}

	return res
}

func augmentGrid(s []string) []string {
	res := []string{}

	inGrid := true
	for _, l := range s {
		if l == "" {
			inGrid = false
			res = append(res, "")
			continue
		}

		if !inGrid {
			res = append(res, l)
			continue
		}

		var row strings.Builder
		for _, c := range l {
			switch c {
			case '#':
				row.WriteString("##")
			case 'O':
				row.WriteString("[]")
			case '@':
				row.WriteString("@.")
			case '.':
				row.WriteString("..")
			}
		}

		dbg("row: %v\n", row.String())
		res = append(res, row.String())
	}

	return res
}
