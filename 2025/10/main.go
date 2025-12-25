package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
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
	numRows := len(m.joltage)
	numCols := len(m.buttons)

	maxJoltage := 0

	// Build augmented matrix [A | b]
	aug := make([][]float64, numRows)
	for row := range m.joltage {
		aug[row] = make([]float64, numCols+1)
		for col, v := range m.buttons {
			if v&(1<<row) > 0 {
				aug[row][col] = 1
			}
		}
		aug[row][numCols] = float64(m.joltage[row]) // b column

		maxJoltage = max(maxJoltage, m.joltage[row])
	}

	// Gaussian elimination with partial pivoting → RREF
	pivotCols := []int{} // tracks which columns have pivots
	pivotRow := 0

	const tolerance = 1e-6

	for col := 0; col < numCols && pivotRow < numRows; col++ {
		// Find pivot (largest absolute value in this column)
		maxRow := pivotRow
		for row := pivotRow + 1; row < numRows; row++ {
			if math.Abs(aug[row][col]) > math.Abs(aug[maxRow][col]) {
				maxRow = row
			}
		}

		if math.Abs(aug[maxRow][col]) < tolerance {
			continue // no pivot in this column, it's a free variable
		}

		// Swap rows
		aug[pivotRow], aug[maxRow] = aug[maxRow], aug[pivotRow]

		// Scale pivot row to make pivot = 1
		scale := aug[pivotRow][col]
		for j := col; j <= numCols; j++ {
			aug[pivotRow][j] /= scale
		}

		// Eliminate all other rows (both above and below for RREF)
		for row := 0; row < numRows; row++ {
			if row != pivotRow && math.Abs(aug[row][col]) > tolerance {
				factor := aug[row][col]
				for j := col; j <= numCols; j++ {
					aug[row][j] -= factor * aug[pivotRow][j]
				}
			}
		}

		pivotCols = append(pivotCols, col)
		pivotRow++
	}

	// Identify free columns
	freeCols := []int{}
	pivotSet := make(map[int]bool)
	for _, p := range pivotCols {
		pivotSet[p] = true
	}
	for col := 0; col < numCols; col++ {
		if !pivotSet[col] {
			freeCols = append(freeCols, col)
		}
	}

	dbg("Pivot columns: %v", pivotCols)
	dbg("Free columns: %v", freeCols)

	// Brute force over free variables
	return bruteForce(aug, pivotCols, freeCols, numCols, maxJoltage)
}

func bruteForce(aug [][]float64, pivotCols, freeCols []int, numCols int, maxJoltage int) int {
	numFree := len(freeCols)
	maxVal := maxJoltage // upper bound: no button needs more presses than max joltage

	minSum := math.MaxInt

	// Generate all combinations of free variable values
	freeVals := make([]int, numFree)

	var search func(idx int)
	search = func(idx int) {
		if idx == numFree {
			// Evaluate this combination
			x := make([]float64, numCols)

			// Set free variables
			for i, col := range freeCols {
				x[col] = float64(freeVals[i])
			}

			// Back-substitute for pivot variables
			// RREF: each pivot row i gives us: x[pivotCols[i]] = aug[i][numCols] - sum(aug[i][j]*x[j]) for free j
			valid := true
			for i := len(pivotCols) - 1; i >= 0; i-- {
				pivotCol := pivotCols[i]
				val := aug[i][numCols]
				for j := pivotCol + 1; j < numCols; j++ {
					val -= aug[i][j] * x[j]
				}
				x[pivotCol] = val

				// Check non-negative integer
				rounded := int(math.Round(val))
				if math.Abs(val-float64(rounded)) > 1e-6 || rounded < 0 {
					valid = false
					break
				}
			}

			if valid {
				sum := 0
				for _, v := range x {
					sum += int(math.Round(v))
				}
				if sum < minSum {
					minSum = sum
				}
			}
			return
		}

		for v := 0; v <= maxVal; v++ {
			freeVals[idx] = v
			search(idx + 1)
		}
	}

	search(0)

	dbg("MinSum: %v", minSum)
	if minSum == math.MaxInt {
		return 0
	}
	return minSum
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
