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
	seats map[Seat]rune
	rows  int
	cols  int
}

func newLayout(lines []string) *Layout {
	l := &Layout{
		seats: make(map[Seat]rune),
		rows:  0,
		cols:  0,
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
	c := &Layout{
		seats: make(map[Seat]rune, len(l.seats)),
		cols:  l.cols,
		rows:  l.rows,
	}

	for k, v := range l.seats {
		c.seats[k] = v
	}

	return c
}

func (l *Layout) occupiedNeighbours(p Seat) int {
	count := 0
	for r := -1; r < 2; r++ {
		for c := -1; c < 2; c++ {
			if c == 0 && r == 0 {
				continue
			}

			p2 := Seat{p.r + r, p.c + c}
			if n, ok := l.seats[p2]; ok && n == '#' {
				count++
			}

		}
	}

	return count
}

func (l *Layout) round() (*Layout, bool) {
	cur := l.copy()
	changed := false
	for p, c := range l.seats {
		if c == '.' {
			continue
		}

		n := l.occupiedNeighbours(p)
		if c == 'L' && n == 0 {
			cur.seats[p] = '#'
			changed = true
		} else if c == '#' && n >= 4 {
			cur.seats[p] = 'L'
			changed = true
		}
	}

	return cur, changed
}

func (l *Layout) board() *Layout {
	for {
		changed := false
		l, changed = l.round()
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

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]

	l := newLayout(lines)
	l.print()
	l = l.board()
	part1 = l.occupiedSeats()

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
