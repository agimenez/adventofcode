package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sort"
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
	//dbg("lines: %#v", lines)
	leftSlice := sort.IntSlice{}
	rightSlice := sort.IntSlice{}
	rightCounts := map[int]int{}
	for _, l := range lines {
		parts := strings.Fields(l)
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		leftSlice = append(leftSlice, left)
		rightSlice = append(rightSlice, right)
		//log.Printf("  -> left: %v, right: %v", left, right)
		rightCounts[right]++
	}
	leftSlice.Sort()
	rightSlice.Sort()
	for i := range leftSlice {
		part1 += utils.Abs(leftSlice[i] - rightSlice[i])
		part2 += leftSlice[i] * rightCounts[leftSlice[i]]
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
