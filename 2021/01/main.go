package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
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

func sliceSum(nums []int) int {
	sum := 0

	for i := range nums {
		sum += nums[i]
	}

	return sum
}

func solveWindow(nums []int, w int) int {
	prev := sliceSum(nums[0:w])
	out := 0

	for i := 1; i < len(nums)-w+1; i++ {
		cur := sliceSum(nums[i : i+w])
		dbg("Slice [%d:%d]: %v: %d (prev %d)", i, i+w, nums[i:i+w], cur, prev)

		if cur > prev {
			out++
		}
		prev = cur
	}

	return out
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	dbg("lines: %#v", lines)
	nums := make([]int, len(lines))
	for i := range lines {
		num, _ := strconv.Atoi(lines[i])
		nums[i] = num
	}
	part1 = solveWindow(nums, 1)
	part2 = solveWindow(nums, 3)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
