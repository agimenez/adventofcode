package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
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
	m[ast] = make(map[float64]Asteroid)
	for candidate := range m {
		if candidate == ast {
			continue
		}
		dbg(3, "Candidate: %v", candidate)

		// Y axis is inverted; 0 is top, and it "grows" downwards, so we reverse the
		// first part of the Y coordinate calculation of the vector to be passed to
		// math.Atan2. Also, we invert x and y to shift pi/2, or start angles at
		// North
		tmp := Asteroid{
			x: candidate.x - ast.x,
			y: ast.y - candidate.y,
		}
		dbg(3, " Got point for angle: %v", tmp)
		// Swap x and y to get clockwise angles starting from pi/4 (north)
		angle := math.Atan2(float64(tmp.x), float64(tmp.y))
		dbg(3, "  -> angle %f", angle)
		if angle < 0 {
			angle += 2 * math.Pi
			dbg(3, "  |-> recalculated to %f", angle)
		}

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

func (m Map) getSortedListSights(a Asteroid) []Asteroid {
	l := []float64{}
	s := []Asteroid{}
	for angle := range m[a] {
		l = append(l, angle)
	}

	sort.Float64s(l)
	for _, angle := range l {
		s = append(s, m[a][angle])
	}

	return s
}

func (m Map) vaporizeFrom(station Asteroid) []Asteroid {
	list := []Asteroid{}

	for len(m) > 1 {
		m.calculateSights(station)
		s := m.getSortedListSights(station)
		dbg(3, "List of sorted sights: %v", s)

		list = append(list, s...)
		dbg(2, "map before deletion: len %d, v = %v", len(m), m)
		for _, a := range s {
			delete(m, a)
		}
		//m.deleteAsteroids(s)
		dbg(2, "New len: %d, v = %v", len(m), m)
	}

	return list
}

func (m *Map) deleteAsteroids(ast []Asteroid) {
	for _, a := range ast {
		delete(*m, a)
	}

	dbg(3, "map after deletion: %v", *m)
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

	// part two: vaporize all the map
	vaporized := m.vaporizeFrom(a)

	fmt.Printf("Vaporized[200]: %v, result = %d\n", vaporized[199], vaporized[199].x*100+vaporized[199].y)

}
