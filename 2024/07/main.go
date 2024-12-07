package main

import (
	"flag"
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
		if res, ok := eqIsCalibrated(l); ok {
			dbg("===== Calibrated!!!")
			part1 += res
		}
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}

func eqIsCalibrated(s string) (int, bool) {
	parts := strings.Split(s, ": ")
	result := ToInt(parts[0])
	operands := CSVToIntSlice(parts[1], " ")
	partial := operands[0]
	rest := operands[1:]

	if isCalibrated(result, partial, rest) {
		return result, true
	}

	return result, false

}

func isCalibrated(res int, partial int, rest []int) bool {
	dbg("isCalibrated: result = %d, partial = %d, rest = %v", res, partial, rest)
	sum := partial + rest[0]
	prod := partial * rest[0]
	dbg(" -> sum = %d", sum)
	dbg(" -> prod = %d", prod)
	if len(rest) == 1 {
		return res == sum || res == prod
	}

	return isCalibrated(res, sum, rest[1:]) || isCalibrated(res, prod, rest[1:])
}
