package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/agimenez/adventofcode/utils"
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

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	for i := range lines {
		ints := line2digits(lines[i])
		if isSafe(ints) {
			part1++
			part2++
		} else if isSafeTolerant(ints) {
			part2++
		}
	}
	//dbg("lines: %#v", lines)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}

func line2digits(s string) []int {
	ints := make([]int, 0, len(s))

	for _, v := range strings.Fields(s) {
		n, _ := strconv.Atoi(v)
		ints = append(ints, n)
	}

	return ints

}

func getDirection(n, n1 int) int {

	diff := n - n1
	switch {
	case diff < 0:
		return -1
	case diff > 0:
		return 1
	}

	return 0
}

func isSafe(report []int) bool {
	dbg("   * Checking safety of %v", report)
	// -1
	direction := 0
	for i := 0; i < len(report)-1; i++ {
		n := report[i]
		n1 := report[i+1]
		dir := getDirection(n, n1)
		if direction == 0 {
			direction = dir
		}

		diff := utils.Abs(n - n1)
		if direction != dir || diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isSafeTolerant(report []int) bool {
	dbg("Trying report: %v", report)
	dbg(" - %v ", report[1:])
	if isSafe(report[1:]) {
		return true
	}

	for i := 1; i < len(report)-1; i++ {
		candidate := slices.Concat(report[:i], report[i+1:])
		dbg(" - %v ", candidate)
		if isSafe(candidate) {
			return true
		}
	}

	dbg(" - %v ", report[:len(report)-1])
	if isSafe(report[:len(report)-1]) {
		return true
	}

	return false
}
