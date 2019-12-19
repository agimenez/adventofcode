package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

var (
	debug int
)

func dbg(level int, fmt string, v ...interface{}) {
	if debug >= level {
		log.Printf(fmt, v...)
	}
}

type Asteroid struct {
	x, y int
}

type Map map[Asteroid]map[float64]Asteroid

func parseInput(in io.Reader) Map {
	m := Map{}
	scanner := bufio.NewScanner(in)
	line := 0
	for scanner.Scan() {
		row := scanner.Text()
		for i, c := range row {
			dbg(2, "Pos (%d, %d): %c\n", i, line, c)
			if c != '#' {
				continue
			}
			dbg(1, "Asteroid @ (%d, %d): %c\n", i, line, c)

			a := Asteroid{
				y: line,
				x: i,
			}

			m[a] = make(map[float64]Asteroid)

		}
		line++
	}

	return m
}

func (a *Asteroid) distance(b Asteroid) float64 {
	return math.Sqrt(float64((b.x-a.x)*(b.x-a.x) + (b.y-a.y)*(b.y-a.y)))
}

// Get all seen asteroids from one of them
func (m Map) calculateSights(ast Asteroid) {
	for candidate := range m {
		if candidate == ast {
			continue
		}
		dbg(3, "Candidate: %v", candidate)

		angle := math.Atan2(float64(candidate.y-ast.y), float64(candidate.x-ast.x))
		dbg(3, "  -> angle %f", angle)

		distance := ast.distance(candidate)
		dbg(3, "  -> distance %f", distance)

		if prev, ok := m[ast][angle]; ok {
			prevdist := ast.distance(prev)
			dbg(3, "   -> (got %v at same angle, distance %f)", prev, prevdist)
			if prevdist <= distance {
				continue
			}

			dbg(3, "   -> (new distance is less. Updating)")
		}
		dbg(3, "   -> New best is %v", candidate)
		m[ast][angle] = candidate

	}
}

func (m *Map) calculateAllSights() {
	for ast := range *m {
		dbg(3, "Calculating sight for Asteroid at %v", ast)
		m.calculateSights(ast)
	}
}

func (m *Map) getBestLocation() Asteroid {
	bestLocation := Asteroid{}
	bestSights := 0

	for a, s := range *m {
		dbg(3, "Asteroid %v, sights %d", a, len(s))
		if len(s) > bestSights {
			dbg(3, " -> New best! (%d > %d)", len(s), bestSights)
			bestLocation = a
			bestSights = len(s)
		}
	}

	return bestLocation
}

func init() {
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.Parse()
}

func main() {

	m := parseInput(os.Stdin)
	m.calculateAllSights()
	a := m.getBestLocation()

	fmt.Printf("Best Asteroid %v, max sights %d\n", a, len(m[a]))

}
