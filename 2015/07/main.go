package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"maps"
	"os"
	"strings"
	"time"

	. "github.com/agimenez/adventofcode/utils"
)

var (
	debug bool
)

func dbg(f string, v ...interface{}) {
	if debug {
		fmt.Printf(f+"\n", v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
}
func main() {
	flag.Parse()

	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	part1, part2, dur1, dur2 := solve(lines)
	log.Printf("Part 1 (%v): %v\n", dur1, part1)
	log.Printf("Part 2 (%v): %v\n", dur2, part2)

}

func solve(lines []string) (int, int, time.Duration, time.Duration) {
	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 := solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 := solve2(lines)
	dur[1] = time.Since(now)

	return part1, part2, dur[0], dur[1]

}

type GateFunc func(inputs []*wire) int

type wire struct {
	inputs []string
	value  uint16
}

type circuit struct {
	wires   map[string]wire
	signals map[string]uint16
}

func NewCircuit() circuit {
	return circuit{
		wires:   map[string]wire{},
		signals: map[string]uint16{},
	}
}

func (c circuit) AddLine(line string) circuit {
	parts := strings.Split(line, " -> ")
	dest := parts[1]

	src := strings.Fields(parts[0])
	// Simple case: 123 -> x
	// Can also be lx -> x
	if len(src) == 1 && '0' <= src[0][0] && src[0][0] <= '9' {
		c.signals[dest] = uint16(ToInt(src[0]))

		return c
	}

	// Rest of the operations:
	// lx > x
	// NOT y -> i
	// x AND y -> d
	// etc...
	c.wires[dest] = wire{
		inputs: src,
	}

	return c
}

func (c circuit) resolveWireRecursive(wire string) (circuit, uint16) {
	var res uint16 = 0

	dbg("RESOLVE %q", wire)
	// Immediate value. I don't like doing this 100%, but it's an easy hack
	if wire[0] >= '0' && wire[0] <= '9' {
		v := uint16(ToInt(wire))
		c.signals[wire] = v
		dbg(" >> RETURN immediate %v", v)
		return c, v
	}

	// immediate signal, return value
	if v, ok := c.signals[wire]; ok {
		dbg(" >> RETURN processed %v", v)
		return c, v
	}

	w := c.wires[wire]

	dbg("WIRE: %v", w)
	// lx -> x
	if len(w.inputs) == 1 {
		c, res = c.resolveWireRecursive(w.inputs[0])
		dbg("  >> Resolved input %q -> %016b", w.inputs[0], res)
		c.signals[wire] = res
		dbg("  >> RESULT:  input %q -> %016b", wire, res)

		return c, res
	}

	// NOT x -> h
	if len(w.inputs) == 2 {
		c, res = c.resolveWireRecursive(w.inputs[1])
		dbg("  >> Resolved input %q -> %016b", w.inputs[1], res)
		res = ^res
		c.signals[wire] = res
		dbg("  >> RESULT:  input %q -> %016b", wire, res)

		return c, res
	}

	// Other ops:
	// x AND y -> d
	// x LSHIFT 2 -> f
	var in1, in2 uint16
	c, in1 = c.resolveWireRecursive(w.inputs[0])
	c, in2 = c.resolveWireRecursive(w.inputs[2])
	dbg("  >> Resolved input %q -> %016b", w.inputs[0], in1)
	dbg("  >> Resolved input %q -> %016b", w.inputs[2], in2)

	switch w.inputs[1] {
	case "AND":
		res = in1 & in2
	case "OR":
		res = in1 | in2
	case "LSHIFT":
		res = in1 << ToInt(w.inputs[2])
	case "RSHIFT":
		res = in1 >> ToInt(w.inputs[2])
	default:
		panic("Unknown op: " + w.inputs[1])
	}
	dbg("  >> RESULT:  input %q -> %016b", wire, res)
	dbg("")

	c.signals[wire] = res

	return c, res
}

func solve1(s []string) int {
	var res uint16 = 0

	c := NewCircuit()
	for _, line := range s {
		c = c.AddLine(line)

	}

	c, res = c.resolveWireRecursive("a")

	return int(res)
}

func solve2(s []string) int {
	var res uint16 = 0

	c := NewCircuit()
	for _, line := range s {
		c = c.AddLine(line)

	}
	// Keep the initial map of direct signals
	origSignals := maps.Clone(c.signals)

	c, res = c.resolveWireRecursive("a")
	origSignals["b"] = res
	c.signals = origSignals
	c, res = c.resolveWireRecursive("a")

	return int(res)
}
