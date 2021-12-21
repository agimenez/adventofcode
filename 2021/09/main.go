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
		dbg("Checking %v (%v)", p, v)
		pointIsLower := true
		for _, n := range hm.pointNeighbours(p) {
			dbg(" -> Neighbour %v: %v", n, hm[n])
			if hm[n] <= v {
				pointIsLower = false
				break
			}
		}

		if pointIsLower {
			lowPoints = append(lowPoints, p)
		}

	}

	return lowPoints
}

func (hm heightMap) riskLevel(p Point) int {
	return hm[p] + 1
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
func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]

	hm := NewHeightMap(lines)
	//dbg("Height map: %v", hm)
	lowPoints := hm.findLowPoints()
	dbg("Low points: %v", lowPoints)
	//printPoints(lowPoints)
	for _, p := range lowPoints {
		part1 += hm.riskLevel(p)
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
