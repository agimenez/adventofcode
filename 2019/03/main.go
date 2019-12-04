package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

const (
	OpSum = 1
	OpMul = 2

	debug = true
)

type point struct {
	x, y int
}

type motion struct {
	dir    rune
	length int
}

func (m motion) String() string {
	return fmt.Sprintf("%c:%d", m.dir, m.length)
}

type circuit struct {
	wirings       []map[point]bool
	shortCircuits map[point]bool
}

func newCircuit() *circuit {
	return &circuit{
		wirings:       []map[point]bool{},
		shortCircuits: make(map[point]bool, 100),
	}
}

func (c *circuit) runWiring(wire []motion) {
	dbg("*** Running new wiring %d ***", len(c.wirings))
	cableTip := point{0, 0}

	c.wirings = append(c.wirings, map[point]bool{})

	for _, m := range wire {
		cableTip = c.doMotion(m, cableTip)
	}

	c.checkShortCircuits()
}

// checkShortCircuits matches the last wiring against the previous ones to find short
// circuits
func (c *circuit) checkShortCircuits() {
	if len(c.wirings) < 2 {
		return
	}

	topWire := c.wirings[len(c.wirings)-1]

	for i, wiring := range c.wirings[:len(c.wirings)-1] {
		dbg("Checking shortcircuits %d/%d", i, len(c.wirings)-1)
		for point := range topWire {
			dbg("\t-> Point %v", point)
			if wiring[point] {
				dbg("\t   *** MATCH! ***")
				c.shortCircuits[point] = true
			}
		}
	}

}

func (c *circuit) doMotion(m motion, cableTip point) point {
	wire := len(c.wirings) - 1
	dbg(" Motion Wire %d: %c -> %d", wire, m.dir, m.length)

	for i := 0; i < m.length; i++ {
		switch m.dir {
		case 'U':
			cableTip.y++
		case 'D':
			cableTip.y--
		case 'L':
			cableTip.x--
		case 'R':
			cableTip.x++
		default:
			panic(fmt.Sprintf("Bad input motion direction %c", m.dir))
		}

		c.wirings[wire][cableTip] = true
		//dbg("  -> New cableTip: %v", c.cableTip)
	}

	return cableTip
}

func (c *circuit) minManhattanShortCircuit() int {
	minDistance := math.MaxInt32

	for point := range c.shortCircuits {
		d := abs(point.x) + abs(point.y)
		if d < minDistance {
			dbg("Got new min distance %d", d)
			minDistance = d
		}
	}

	return minDistance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

func main() {
	var wire1 string
	var wire2 string

	fmt.Scan(&wire1)
	fmt.Scan(&wire2)

	dbg("wire1: %s", wire1)
	dbg("wire2: %s", wire2)

	w1 := parseWiring(wire1)
	w2 := parseWiring(wire2)

	c := newCircuit()
	c.runWiring(w1)
	c.runWiring(w2)

	log.Printf("Min Manhattan: %d", c.minManhattanShortCircuit())

}

func run(p []int) {
}

func parseWiring(p string) []motion {
	var wiring []motion

	pSlice := strings.Split(p, ",")
	for _, b := range pSlice {
		var m motion

		fmt.Sscanf(b, "%c%d", &m.dir, &m.length)
		dbg("Got motion: %v", m)

		wiring = append(wiring, m)
	}

	return wiring
}
