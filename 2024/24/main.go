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
	part2 = solve2(circuit)
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

func (g gate) inputsXY() bool {
	return (g.in0[0] == 'x' && g.in1[0] == 'y' || g.in0[0] == 'y' && g.in1[0] == 'x')
}

func (g gate) inputsXY00() bool {
	return g.in0 == "x00" && g.in1 == "y00" || g.in0 == "y00" && g.in1 == "x00"
}

// Oh boi, this was too much for me.
// Just try to implement https://www.reddit.com/r/adventofcode/comments/1hla5ql/2024_day_24_part_2_a_guide_on_the_idea_behind_the/?utm_name=web3xcss
func solve2(c circuit) int {
	res := 0
	swap := []string{}
	debug = true
	for name, gate := range c.gates {
		// If the output of a gate is z, then the operation has to be XOR unless it is the last bit.
		if gate.out[0] == 'z' && gate.out != "z45" && gate.op != "XOR" {
			dbg(" -> %v: %+v BAD Z NON-XOR!", name, gate)
			swap = append(swap, name)
			continue
		}

		// If the output of a gate is not z and the inputs are not x, y then it has to be AND / OR, but not XOR.
		if gate.out[0] != 'z' && gate.op == "XOR" && !gate.inputsXY() {
			if gate.op == "XOR" {
				dbg(" -> %v: %+v BAD NON-Z XOR!", name, gate)
				swap = append(swap, name)
				continue
			}
		}

		// If you have a XOR gate with inputs x, y, there must be
		// another XOR gate with this gate as an input. Search through
		// all gates for an XOR-gate with this gate as an input; if it
		// does not exist, your (original) XOR gate is faulty.
		// These don't apply for the gates with input x00, y00
		if gate.op == "XOR" && gate.inputsXY() && !gate.inputsXY00() {
			found := false
			for _, g := range c.gates {
				if g.op == "XOR" && (g.in0 == name || g.in1 == name) {
					found = true
					break
				}

			}
			if !found {
				dbg(" -> %v: %+v BAD XOR", name, gate)
				swap = append(swap, name)
			}

			continue
		}

		// If you have an AND-gate, there must be an OR-gate with this
		// gate as an input. If that gate doesn't exist, the original
		// AND gate is faulty.
		// These don't apply for the gates with input x00, y00
		if gate.op == "AND" && gate.inputsXY() && !gate.inputsXY00() {
			found := false
			for _, g := range c.gates {
				if g.op == "OR" && (g.in0 == name || g.in1 == name) {
					found = true
					break
				}

			}
			if !found {
				dbg(" -> %v: %+v BAD AND", name, gate)
				swap = append(swap, name)
			}

			continue
		}

	}
	slices.Sort(swap)
	dbg("Swaps: %v", strings.Join(swap, ","))

	return res
}
