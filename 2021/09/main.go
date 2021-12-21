package main

import (
	"flag"
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

type heightMap map[Point]int

func NewHeightMap(in []string) heightMap {
	hm := heightMap{}

	for row := range in {
		for col := range in[row] {
			hm[Point{row, col}] = int(in[row][col] - '0')
		}
	}

	return hm
}

func (hm heightMap) pointNeighbours(p Point) []Point {
	n := []Point{}
	ops := []func(Point) Point{
		Point.Up,
		Point.Down,
		Point.Left,
		Point.Right,
	}

	for _, op := range ops {
		cur := op(p)
		if _, ok := hm[cur]; ok {
			n = append(n, cur)
		}
	}

	return n
}

func (hm heightMap) findLowPoints() []Point {
	lowPoints := []Point{}
	for p, v := range hm {
		pointIsLower := false
		for _, n := range hm.pointNeighbours(p) {
			if hm[n] < v {
				break
			}
		}

		if pointIsLower {
			lowPoints = append(lowPoints, p)
		}

	}

	return lowPoints
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]

	hm := NewHeightMap(lines)
	lowPoints := hm.findLowPoints()
	dbg("Low points: %v", lowPoints)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
