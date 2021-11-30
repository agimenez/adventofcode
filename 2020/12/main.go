package main

import (
	"bufio"
	"flag"
	"log"
	"math"
	"os"
	"strconv"
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

var Facing []Point = []Point{
	North: Point{0, 1},
	East:  Point{1, 0},
	South: Point{0, -1},
	West:  Point{-1, 0},
}

const (
	North = iota
	East
	South
	West
)

type Point struct {
	x, y int
}

func (p Point) rotate90(degrees int) Point {
	p2 := []Point{
		Point{p.x, p.y},
		Point{p.y, -p.x},
		Point{-p.x, -p.y},
		Point{-p.y, p.x},
	}

	jumps := degrees / 90
	return p2[mod(jumps, len(p2))]
}

type Ship struct {
	wp   Point
	loc  Point
	path []Point
}

func newShip(facing Point) *Ship {
	s := &Ship{
		wp:   facing,
		loc:  Point{0, 0},
		path: []Point{{0, 0}},
	}

	return s
}

func (s *Ship) manhattanDistance() int {
	return int(math.Abs(float64(s.loc.x)) + math.Abs(float64(s.loc.y)))
}

func (s *Ship) forward(units int) {
	s.loc.x += s.wp.x * units
	s.loc.y += s.wp.y * units

	s.path = append(s.path, s.loc)
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func (s *Ship) rotate(dir rune, degrees int) {
	if dir == 'L' {
		degrees = -degrees
	}

	dbg("Rotate %d cur: wp = %v", degrees, s.wp)
	s.wp = s.wp.rotate90(degrees)
	dbg("Rotate %d aft: wp = %v", degrees, s.wp)

}

func (s *Ship) shift(dir rune, value int) {
	dirMap := map[rune]int{
		'N': North,
		'E': East,
		'S': South,
		'W': West,
	}

	s.wp.x += value * Facing[dirMap[dir]].x
	s.wp.y += value * Facing[dirMap[dir]].y

	s.path = append(s.path, s.loc)
}

func (s *Ship) move(inst string) {
	action := rune(inst[0])
	value, err := strconv.Atoi(inst[1:])
	if err != nil {
		panic("can't parse value")
	}

	switch action {

	case 'R', 'L':
		s.rotate(action, value)
	case 'F':
		s.forward(value)
	default:
		s.shift(action, value)
	}
}

func main() {

	s := bufio.NewScanner(os.Stdin)
	part1, part2 := 0, 0
	f := newShip(Point{10, 1})

	for s.Scan() {
		l := s.Text()
		f.move(l)
	}
	part2 = f.manhattanDistance()

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
