package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sort"
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

func sumCals(in []string) []int {
	totals := []int{}

	elfSum := 0
	for _, c := range in {
		if c == "" {
			totals = append(totals, elfSum)
			elfSum = 0
			continue
		}

		cals, _ := strconv.Atoi(c)
		elfSum += cals

	}
	totals = append(totals, elfSum)

	return totals
}

func sumInts(in []int) int {

	total := 0
	for _, v := range in {
		total += v
	}

	return total
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

	totals := sumCals(lines)
	dbg("totals: %#v", totals)
	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	dbg("totals: %#v", totals)
	part1 = totals[0]
	part2 = sumInts(totals[:3])

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
