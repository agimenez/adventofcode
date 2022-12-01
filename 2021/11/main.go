package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

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
	flag.Parse()
}

type octopusMap map[Point]int

func NewOctopusMap(in []string) octopusMap {
	om := octopusMap{}

	for row := range in {
		for col := range in[row] {
			om[Point{row, col}] = int(in[row][col] - '0')
		}
	}

	return om
}

func (om octopusMap) pointNeighbours(p Point) []Point {
	n := []Point{}
	ops := []func(Point) Point{
		Point.Up,
		Point.Down,
		Point.Left,
		Point.Right,
		func(p Point) Point {
			return p.Up().Left()
		},
		func(p Point) Point {
			return p.Up().Right()
		},
		func(p Point) Point {
			return p.Down().Left()
		},
		func(p Point) Point {
			return p.Down().Right()
		},
	}

	for _, op := range ops {
		cur := op(p)
		if _, ok := om[cur]; ok {
			n = append(n, cur)
		}
	}

	return n
}

func printPoints(in []Point) {
	m := map[Point]bool{}
	for _, p := range in {
		m[p] = true
	}

	for r := 0; r < 100; r++ {
		for c := 0; c < 100; c++ {
			if v, ok := m[Point{r, c}]; ok {
				fmt.Printf("{%v %v} lowest (%v)\n", r, c, v)
			}
		}
	}
}

func (om octopusMap) countFlashes(steps int) int {
	nflashes := 0
	for i := 0; i < steps; i++ {
		toFlash := make([]Point, 0)
		for o := range *om {
			om[o]++

			// Flash
			if om[o] > 9 {
				toFlash = append(toFlash, o)
			}
		}
		flashes := make(map[Point]bool)
		for len(toFlash) > 0 {
			for _, point := range toFlash {

			}
		}
	}

	return nflashes
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]

	om := NewOctopusMap(lines)
	dbg("Octopi map: %v", om)
	//printPoints(lowPoints)
	part1 = om.countFlashes(2)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
