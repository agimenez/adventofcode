package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
	"time"

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

type gate struct {
	// inputs
	in0, in1 string

	// output to
	out string

	// Operation
	op string
}

func (c circuit) printGate(g gate) string {
	v1 := fmt.Sprintf("%d", c.wires[g.in0])
	if _, ok := c.wires[g.in0]; !ok {
		v1 = "nil"
	}

	v2 := fmt.Sprintf("%d", c.wires[g.in1])
	if _, ok := c.wires[g.in1]; !ok {
		v2 = "nil"
	}
	v3 := fmt.Sprintf("%d", c.wires[g.out])
	if _, ok := c.wires[g.out]; !ok {
		v3 = "nil"
	}
	return fmt.Sprintf("%s (%v) %s %s (%v) -> %s (%v)", g.in0, v1, g.op, g.in1, v2, g.out, v3)
}

type circuit struct {
	wires map[string]int

	gates map[string]gate
}

func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)
	circuit := parseInput(lines)
	dbg("Circuit: %+v", circuit)

	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 = solve1(circuit)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 = solve2(lines)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}
func (c circuit) simulate() circuit {
	i := 0
	pending := map[string]bool{}
	for {
		for _, g := range c.gates {
			pending[g.out] = true
			dbg("Wires: %v", c.wires)
			dbg("Gate %v", c.printGate(g))
			dbg("Pending %v", pending)

			// Inputs are ready?
			if _, ok := c.wires[g.in0]; !ok {
				dbg(" -> wire %v not ready!", g.in0)
				continue
			}

			if _, ok := c.wires[g.in1]; !ok {
				dbg(" -> wire %v not ready!", g.in1)
				continue
			}

			switch g.op {
			case "AND":
				c.wires[g.out] = c.wires[g.in0] & c.wires[g.in1]
			case "OR":
				c.wires[g.out] = c.wires[g.in0] | c.wires[g.in1]
			case "XOR":
				c.wires[g.out] = c.wires[g.in0] ^ c.wires[g.in1]
			}
			delete(pending, g.out)
			dbg(" -> WIRES READY! %v -> %v", g.out, c.wires[g.out])
			dbg("===============")
		}
		dbg("")
		dbg("== Pending %v", pending)
		dbg("")

		i++
		if len(pending) == 0 {
			break
		}
	}

	return c
}

func parseInput(lines []string) circuit {
	c := circuit{
		wires: map[string]int{},
		gates: map[string]gate{},
	}

	inGates := false
	for _, l := range lines {
		if l == "" {
			inGates = true
			continue
		}

		if !inGates {
			parts := strings.Split(l, ": ")
			c.wires[parts[0]] = ToInt(parts[1])
		} else {
			// x00 AND y00 -> z00
			parts := strings.Split(l, " ")
			gate := gate{
				in0: parts[0],
				op:  parts[1],
				in1: parts[2],
				out: parts[4],
			}

			c.gates[parts[4]] = gate
		}
	}

	return c
}

func solve1(c circuit) int {
	res := 0
	c = c.simulate()
	zWires := []string{}
	for w := range c.wires {
		if w[0] == 'z' {
			zWires = append(zWires, w)
		}
	}
	slices.Sort(zWires)
	dbg("Sorted zWires: %v", zWires)
	for w, v := range zWires {
		res |= c.wires[v] << w
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
