package main

import (
	"flag"
	"io"
	"log"
	"os"
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

type machine struct {
	lights  int
	wiring  []int
	joltage []int
}

func (m machine) bfs() int {

	dist := map[int]int{}
	// alues of the current path's lights
	queue := []int{0}
	for len(queue) > 0 {
		dbg("[target: %08b] state: %v, queue: %v", m.lights, dist, queue)
		cur := queue[0]
		queue = queue[1:]
		cost := dist[cur]
		dbg("  >> cur: %08b", cur)

		if cur == m.lights {
			return cost
		}

		for _, button := range m.wiring {
			next := button ^ cur
			dbg("  >>>> next = button ^ cur => %08b = %08b ^ %08b", next, button, cur)
			if _, visited := dist[next]; !visited {
				dbg("  >>>>>>> next (%08b) not visited!", next)
				dist[next] = cost + 1
				queue = append(queue, next)
			}
			dbg("  >>>> Queue: %v", queue)
		}
		dbg("  >> Queue: %v", queue)
	}

	return -1
}

func parseMachine(s string) machine {
	m := machine{}

	parts := strings.Fields(s)
	// First, parse the lights [..##..]
	for i, c := range parts[0][1 : len(parts[0])-1] {
		if c == '#' {
			// This will be "reversed", but who cares?
			m.lights |= (1 << i)
		}
	}

	dbg("Parts: %v", parts)
	// then parse the wiring diagrams
	for _, dia := range parts[1 : len(parts)-1] {
		dbg("  >> Diagram: %v", dia)
		button := 0
		for _, pos := range strings.Split(dia[1:len(dia)-1], ",") {
			v := ToInt(pos)
			button |= (1 << v)
			dbg("  >>>> pos: %v, v: %8b, b: %08b", pos, (1 << v), button)
		}
		m.wiring = append(m.wiring, button)
	}

	return m
}

func minButtons(s string) int {
	m := parseMachine(s)
	dbg("Machine: %v", m)

	return m.bfs()
}

func solve1(s []string) int {
	res := 0
	for _, line := range s {
		res += minButtons(line)
	}

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
