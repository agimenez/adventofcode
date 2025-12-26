package main

import (
	"flag"
	"fmt"
	. "github.com/agimenez/adventofcode/utils"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var (
	debug bool
)

func dbg(f string, v ...interface{}) {
	if debug {
		fmt.Printf(f+"\n", v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
}
func main() {
	flag.Parse()

	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	part1, part2, dur1, dur2 := solve(lines)
	log.Printf("Part 1 (%v): %v\n", dur1, part1)
	log.Printf("Part 2 (%v): %v\n", dur2, part2)

}

func solve(lines []string) (int, int, time.Duration, time.Duration) {
	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 := solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 := solve2(lines)
	dur[1] = time.Since(now)

	return part1, part2, dur[0], dur[1]

}

func lookAndSay(seq []int) []int {
	ret := []int{}
	for i := 0; i < len(seq); i++ {
		num := seq[i]
		count := 1
		for j := i + 1; j < len(seq) && seq[j] == num; j++ {
			count++
		}
		i += count - 1

		ret = append(ret, count, num)

	}

	return ret
}

func solve1(s []string) int {
	res := 0

	nums := CSVToIntSlice(s[0], "")
	for range 40 {
		nums = lookAndSay(nums)
	}
	res = len(nums)

	return res
}

func solve2(s []string) int {
	res := 0

	nums := CSVToIntSlice(s[0], "")
	for range 50 {
		nums = lookAndSay(nums)
	}
	res = len(nums)

	return res
}
