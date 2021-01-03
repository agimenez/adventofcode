package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"testing"
)

var (
	debug bool
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

// This is disappointing from Go :(
// https://stackoverflow.com/a/58192326/4735682
var _ = func() bool {
	testing.Init()
	return true
}()

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
	flag.Parse()
}

func MemoryGame(start []int, until int) int {
	spoken := make(map[int]int)
	lastSpoken := 0
	turn := 1
	for _, n := range start {
		spoken[n] = turn
		lastSpoken = n
		turn++
	}

	for ; turn <= until; turn++ {
		var speak int
		dbg("Turn %d, last = %d", turn, lastSpoken)
		if t, ok := spoken[lastSpoken]; ok {
			speak = turn - 1 - t
		} else {
			speak = 0
		}
		dbg("To speak: %d", speak)
		spoken[lastSpoken] = turn - 1
		lastSpoken = speak

	}

	return lastSpoken
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]

	numbers := strings.Split(lines[0], ",")
	var nums []int
	for i := range numbers {
		n, err := strconv.Atoi(numbers[i])
		if err != nil {
			log.Fatalf("Error converting %s", numbers[i])
		}
		nums = append(nums, n)
	}
	part1 = MemoryGame(nums, 2020)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
