package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"
	"time"

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
}
func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 = solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 = solve2(lines)
	dur[1] = time.Since(now)

	log.Printf("Part 1 (%v): %v\n", dur[0], part1)
	log.Printf("Part 2 (%v): %v\n", dur[1], part2)

}

func solve1(s []string) int {
	res := 0

	for _, l := range s {
		secret := evolve(ToInt(l), 2000)
		res += secret
		dbg("%v -> %v", l, secret)
	}

	return res
}

func evolve(secret int, loops int) int {
	for loops > 0 {
		secret = nextSecret(secret)
		loops--
	}

	return secret
}

func nextSecret(secret int) int {
	secret = prune(mix(secret, secret*64))
	secret = prune(mix(secret, secret/32))
	secret = prune(mix(secret, secret*2048))

	return secret
}

func prune(secret int) int {
	return secret % (1 << 24)
}

func mix(secret int, value int) int {
	return secret ^ value
}

func solve2(s []string) int {
	res := 0

	return res
}
