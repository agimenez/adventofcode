package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

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

type Operation func(a, b int) int

var Sum Operation = func(a, b int) int {
	return a + b
}
var Mul Operation = func(a, b int) int {
	return a * b
}

// lazy version
var Concat Operation = func(a, b int) int {
	return ToInt(fmt.Sprintf("%d%d", a, b))
}

func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	for _, l := range lines {
		dbg("========= New Op: %q", l)
		if res, ok := eqIsCalibrated(l, []Operation{Sum, Mul}); ok {
			dbg("===== PART 1: Calibrated!!!")
			part1 += res
		}

		if res, ok := eqIsCalibrated(l, []Operation{Sum, Mul, Concat}); ok {
			dbg("===== PART 2: Calibrated!!!")
			part2 += res
		}
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}

func eqIsCalibrated(s string, ops []Operation) (int, bool) {
	parts := strings.Split(s, ": ")
	result := ToInt(parts[0])
	operands := CSVToIntSlice(parts[1], " ")
	partial := operands[0]
	rest := operands[1:]

	if isCalibrated(result, partial, rest, ops) {
		return result, true
	}

	return result, false

}

func isCalibrated(res int, partial int, rest []int, ops []Operation) bool {
	dbg("isCalibrated: result = %d, partial = %d, rest = %v", res, partial, rest)
	if len(rest) == 0 {
		return false
	}

	for _, op := range ops {
		opresult := op(partial, rest[0])
		dbg(" -> partial OP rest[0] -> %d OP %d = %d", partial, rest[0], opresult)
		if len(rest) == 1 && res == opresult {
			return true
		}

		if isCalibrated(res, opresult, rest[1:], ops) {
			return true
		}
	}

	return false
}
