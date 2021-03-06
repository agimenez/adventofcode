package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
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
	flag.Parse()
}

type costFunction func(int, int) int

func linearCost(a, b int) int {
	return Abs(a - b)
}

func triangularCost(a, b int) int {
	dist := linearCost(a, b)
	cost := (dist * (dist + 1)) / 2

	return cost
}

func minFuel(pos []int, costFn costFunction) int {
	minFuel := math.MaxInt32

	for i := pos[0]; i < pos[len(pos)-1]; i++ {
		fuel := 0
		for _, cost := range pos {
			fuel += costFn(cost, i)
		}

		minFuel = Min(minFuel, fuel)
	}

	return minFuel
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	fuels := strings.Split(strings.Split(string(p), "\n")[0], ",")
	pos := []int{}
	for _, f := range fuels {
		n, _ := strconv.Atoi(f)
		pos = append(pos, n)
	}
	sort.Ints(pos)

	part1 = minFuel(pos, linearCost)
	part2 = minFuel(pos, triangularCost)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
