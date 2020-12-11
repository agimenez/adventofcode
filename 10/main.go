package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
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

func getVDiff(nums []int) int {
	diffs := make(map[int]int)

	rating := 0
	for i := range nums {
		diff := nums[i] - rating
		diffs[diff]++
		rating = nums[i]
	}

	return (diffs[3] + 1) * diffs[1]
}

func main() {

	s := bufio.NewScanner(os.Stdin)
	part1, part2 := 0, 0
	nums := make([]int, 100)
	for s.Scan() {
		l := s.Text()
		n, err := strconv.Atoi(l)
		if err != nil {
			panic("can't parse input")
		}

		nums = append(nums, n)

	}

	sort.Ints(nums)
	part1 = getVDiff(nums)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
