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

func gammaBalance(b []int) (int64, int64) {
	var g string
	var e string
	for i := range b {
		if b[i] >= 0 {
			g += "1"
			e += "0"
		} else {
			g += "0"
			e += "1"
		}
	}
	gamma, _ := strconv.ParseInt(g, 2, 32)
	epsilon, _ := strconv.ParseInt(e, 2, 32)
	//dbg("Balance for %v: gamma %0b rate %d, epsilon %0b, rate %d", b, gamma, gamma, epsilon, epsilon)

	return gamma, epsilon
}

const (
	MostCommon = iota
	LeastCommon
)

func recurseRating(in []string, pos int, criteria int) []string {
	if len(in) <= 1 {
		return in
	}

	split := map[byte][]string{}

	balance := 0
	dbg("=== Position %d ===", pos)
	for _, v := range in {
		split[v[pos]] = append(split[v[pos]], v)

		if v[pos] == '0' {
			balance--
		} else {
			balance++
		}
		dbg("%s [%d] -> %c: %d", v, pos, v[pos], balance)
	}

	var out []string
	switch criteria {
	case MostCommon:
		if balance >= 0 {
			out = split['1']
		} else {
			out = split['0']
		}

	case LeastCommon:
		if balance >= 0 {
			out = split['0']
		} else {
			out = split['1']
		}
	}
	dbg("Kept: %v", out)

	return recurseRating(out, pos+1, criteria)
}

func Rating(report []string, criteria int) int64 {

	r := recurseRating(report, 0, criteria)

	v, _ := strconv.ParseInt(r[0], 2, 32)
	return v
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

	balance := make([]int, len(lines[0]))
	for i := range lines {
		dbg("Line: %+v", lines[i])
		for j := range lines[i] {
			if lines[i][j] == '0' {
				balance[j]--
			} else {
				balance[j]++
			}
		}
		dbg("Balance: %v", balance)
	}
	gamma, epsilon := gammaBalance(balance)
	part1 = int(gamma * epsilon)

	oxygen := Rating(lines, MostCommon)
	co2 := Rating(lines, LeastCommon)

	part2 = int(oxygen * co2)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
