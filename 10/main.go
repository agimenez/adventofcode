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

// This is shameless stolen from Reddit. Thanks, kaur_virunurm
// It seems you can use memoization, recursivity, and something called "Tribonacci",
// but I can't be bothered.
// https://www.reddit.com/r/adventofcode/comments/ka8z8x/2020_day_10_solutions/gfbo61q/?utm_source=reddit&utm_medium=web2x&context=3
func getCombinations(nums []int) int {
	nums = append([]int{0}, nums...)

	c := make(map[int]int)
	c[0] = 1

	for _, n := range nums {
		dbg("n: %v", n)
		c[n+1] += c[n]
		c[n+2] += c[n]
		c[n+3] += c[n]
		dbg("%#v", c)
	}

	return c[nums[len(nums)-1]+3]
}

func main() {

	s := bufio.NewScanner(os.Stdin)
	part1, part2 := 0, 0
	nums := make([]int, 0, 100)
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

	part2 = getCombinations(nums)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
