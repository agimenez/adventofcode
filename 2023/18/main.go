package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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
}

func toInt(s string) int {
	v, _ := strconv.Atoi(s)

	return v
}

type polygon []Point

func (p polygon) AddDigInstruction(s string) polygon {
	parts := strings.Fields(s)
	dir := parts[0]
	l := parts[1]
	p = p.AddVertex(dir, toInt(l))

	return p

}

func (p polygon) AddVertex(dir string, length int) polygon {
	last := p[len(p)-1]
	var sum Point

	switch dir {
	case "R":
		sum = Point{length, 0}
	case "D":
		sum = Point{0, length}
	case "L":
		sum = Point{-length, 0}
	case "U":
		sum = Point{0, -length}
	}

	p = append(p, last.Sum(sum))

	return p
}

func (p polygon) Area() int {
	// Shoelace formula
	sum := 0
	p0 := p[0]
	for _, p1 := range p {
		sum += p0.X*p1.Y - p1.X*p0.Y
		dbg("%s -> %s (%v). Dist: %v", p0, p1, sum, p0.ManhattanDistance(p1))
		sum += p0.ManhattanDistance(p1)
		p0 = p1
	}

	return Abs(sum/2) + 1
}

func (p polygon) print() {
	if !debug {
		return
	}

	exist := map[Point]bool{}
	for _, p := range p {
		exist[p] = true
	}

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			pt := Point{x, y}
			if _, ok := exist[pt]; ok {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}

}
func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)
	polygon := polygon{{0, 0}}
	for _, l := range lines {
		polygon = polygon.AddDigInstruction(l)
		dbg("Add %s", l)
		polygon.print()
	}

	part1 = polygon.Area()

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
