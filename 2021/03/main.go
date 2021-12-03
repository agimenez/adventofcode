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
	dbg("Balance for %v: gamma %0b rate %d, epsilon %0b, rate %d", b, gamma, gamma, epsilon, epsilon)

	return gamma, epsilon
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

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
