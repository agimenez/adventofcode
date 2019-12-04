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
	wired    map[point]bool
	cableTip point
}

func (c *circuit) runWiring(wire []motion) {
	c.cableTip = point{0, 0}

	for _, m := range wire {
		c.doMotion(m)
	}
}

func (c *circuit) doMotion(m motion) {
	dbg("Motion: %c -> %d", m.dir, m.length)
	switch m.dir {
	case 'U':
		c.cableTip.y += m.length
	case 'D':
		c.cableTip.y -= m.length
	case 'L':
		c.cableTip.x -= m.length
	case 'R':
		c.cableTip.x += m.length
	default:
		panic(fmt.Sprintf("Bad input motion direction %c", m.dir))
	}

	dbg("  -> New cableTip: %v", c.cableTip)
	c.wired[c.cableTip] = true
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
	w2 := parseWiring(wire1)

	c := circuit{
		wired:    make(map[point]bool, len(w1)),
		cableTip: point{0, 0},
	}
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
