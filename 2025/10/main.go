package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	. "github.com/agimenez/adventofcode/utils"
	"gonum.org/v1/gonum/mat"
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

type machine struct {
	lights  int
	buttons []int
	joltage []int
}

func (m machine) bfs() int {

	dist := map[int]int{}
	// Values of the current path's lights
	queue := []int{0}
	for len(queue) > 0 {
		// dbg("[target: %08b] state: %v, queue: %v", m.lights, dist, queue)
		cur := queue[0]
		queue = queue[1:]
		cost := dist[cur]
		// dbg("  >> cur: %08b", cur)

		if cur == m.lights {
			return cost
		}

		for _, button := range m.buttons {
			next := button ^ cur
			// dbg("  >>>> next = button ^ cur => %08b = %08b ^ %08b", next, button, cur)
			if _, visited := dist[next]; !visited {
				// dbg("  >>>>>>> next (%08b) not visited!", next)
				dist[next] = cost + 1
				queue = append(queue, next)
			}
			// dbg("  >>>> Queue: %v", queue)
		}
		// dbg("  >> Queue: %v", queue)
	}

	return -1
}

func (m machine) solveJoltage() int {
	res := 0

	// Solve Ax = b, where A is the wired buttons matrix, and b the joltages
	dbg("== Machine: %v", m)
	A := mat.NewDense(len(m.joltage), len(m.buttons), nil)
	for joltagePos, joltage := range m.joltage {
		dbg("  >> Joltage: %v, pos: %v", joltage, joltagePos)
		for buttonPos, v := range m.buttons {
			// check if the `joltagePos`th bit is set in this wire
			bitmask := 1 << joltagePos
			if v&bitmask > 0 {
				A.Set(joltagePos, buttonPos, 1)
			}
		}
	}
	dbg("Matrix:\n%v", mat.Formatted(A))
	jFloat := []float64{}
	for i := range m.joltage {
		jFloat = append(jFloat, float64(m.joltage[i]))
	}
	b := mat.NewVecDense(len(m.joltage), jFloat)
	dbg("Joltage vector: %v (%v)", b, jFloat)

	var qr mat.QR
	// m<n, so need to transpose
	AT := A.T()
	qr.Factorize(AT)

	// Need to multiply by the transposed
	var ATb mat.VecDense
	ATb.MulVec(AT, b)

	var x mat.VecDense
	err := qr.SolveTo(&x, false, &ATb)
	if err != nil {
		fmt.Printf("ERROR SOLVING: %v\n", err)
		return 0
	}

	dbg("SOLUTION:\n%v", x)

	return res
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

	// then parse the wiring diagrams
	for _, dia := range parts[1 : len(parts)-1] {
		button := 0
		for _, pos := range strings.Split(dia[1:len(dia)-1], ",") {
			v := ToInt(pos)
			button |= (1 << v)
		}
		m.buttons = append(m.buttons, button)
	}

	// parse joltages
	joltages := parts[len(parts)-1]
	for _, joltage := range strings.Split(joltages[1:len(joltages)-1], ",") {
		m.joltage = append(m.joltage, ToInt(joltage))
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
	for _, line := range s {
		m := parseMachine(line)
		dbg("Machine parse: %v", m)
		res += m.solveJoltage()
	}

	return res
}
