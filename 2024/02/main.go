package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
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
		if isSafe(lines[i]) {
			part1++
		}
	}
	//dbg("lines: %#v", lines)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

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

func isSafe(s string) bool {
	// -1
	direction := 0
	digits := strings.Fields(s)
	for i := 0; i < len(digits)-1; i++ {
		n, _ := strconv.Atoi(digits[i])
		n1, _ := strconv.Atoi(digits[i+1])
		dir := getDirection(n, n1)
		if direction == 0 {
			direction = dir
		}

		if direction != dir {
			return false
		}

		diff := utils.Abs(n - n1)
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}
