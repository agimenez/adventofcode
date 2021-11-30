package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	flag.Parse()
}

type Seat struct {
	r, c int
}

type Layout struct {
	seats     map[Seat]rune
	rows      int
	cols      int
	tolerance int
}

func newLayout(lines []string, tolerance int) *Layout {
	l := &Layout{
		seats:     make(map[Seat]rune),
		rows:      0,
		cols:      0,
		tolerance: tolerance,
	}

	for row, line := range lines {
		l.rows = row
		l.cols = len(line)
		for col, c := range line {
			p := Seat{row, col}
			l.seats[p] = c
		}
	}

	return l
}

func (l *Layout) copy() *Layout {
	c := *l
	c.seats = make(map[Seat]rune, len(l.seats))

	for k, v := range l.seats {
		c.seats[k] = v
	}

	return &c
}

func (l *Layout) occupiedNeighbours(p Seat, maxDist int) int {
	count := 0
	dbg("Point: %v", p)
	for r := -1; r < 2; r++ {
		for c := -1; c < 2; c++ {
			if c == 0 && r == 0 {
				continue
			}

			for i := 1; i <= maxDist; i++ {
				p2 := Seat{p.r + r*i, p.c + c*i}
				if n, ok := l.seats[p2]; ok && n == '#' {
					count++
					break
				} else if n == 'L' {
					break
				}
			}

		}
	}
	dbg("Neighbours: %v", count)

	return count
}

func (l *Layout) round(maxDist int) (*Layout, bool) {
	cur := l.copy()
	changed := false
	for p, c := range l.seats {
		if c == '.' {
			continue
		}

		n := l.occupiedNeighbours(p, maxDist)
		if c == 'L' && n == 0 {
			cur.seats[p] = '#'
			changed = true
		} else if c == '#' && n >= l.tolerance {
			cur.seats[p] = 'L'
			changed = true
		}
	}

	return cur, changed
}

func (l *Layout) board(maxDist int) *Layout {
	for {
		changed := false
		l, changed = l.round(maxDist)
		l.print()

		if !changed {
			break
		}
	}

	return l
}

func (l *Layout) print() {
	if !debug {
		return
	}
	for r := 0; r < l.rows; r++ {
		for c := 0; c < l.cols; c++ {
			fmt.Printf("%c", l.seats[Seat{r, c}])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (l Layout) occupiedSeats() int {
	count := 0
	for _, c := range l.seats {
		if c == '#' {
			count++
		}
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]

	l := newLayout(lines, 4)
	l.print()
	l = l.board(1)
	part1 = l.occupiedSeats()
	log.Printf("Part 1: %v\n", part1)

	l = newLayout(lines, 5)
	l.print()
	l = l.board(max(l.rows, l.cols))
	part2 = l.occupiedSeats()

	log.Printf("Part 2: %v\n", part2)

}
