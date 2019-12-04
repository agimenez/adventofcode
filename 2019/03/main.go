package main

import (
	"fmt"
	"log"
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
	wires         int
	shortCircuits map[point]bool
}

func newCircuit() *circuit {
	return &circuit{
		wirings:       []map[point]bool{},
		wires:         0,
		shortCircuits: make(map[point]bool, 100),
	}
}

func (c *circuit) runWiring(wire []motion) {
	dbg("*** Running new wiring %d ***", c.wires)
	cableTip := point{0, 0}

	//c.wiring[c.wires] = make(map[point]bool, len(wire))
	c.wirings = append(c.wirings, map[point]bool{})

	for _, m := range wire {
		cableTip = c.doMotion(m, cableTip)
	}

	c.wires++
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
			dbg(" -> Point %v", point)
			if wiring[point] {
				dbg("   *** MATCH! ***")
				c.shortCircuits[point] = true
			} else {
				dbg("   (nope)")
			}
		}
	}

}

func (c *circuit) doMotion(m motion, cableTip point) point {
	dbg(" Motion Wire %d: %c -> %d", c.wires, m.dir, m.length)

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

		c.wirings[c.wires][cableTip] = true
		//dbg("  -> New cableTip: %v", c.cableTip)
	}

	return cableTip
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
