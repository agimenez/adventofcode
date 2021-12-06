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

type fishCounter map[int]int

func NewFish(in []string) fishCounter {
	f := fishCounter{}

	for _, v := range in {
		n, _ := strconv.Atoi(v)
		f[n]++
	}

	return f
}

func (f fishCounter) simulateDays(n int) int {
	for day := 1; day <= n; day++ {
		prev := f
		f = fishCounter{}
		for counter, num := range prev {
			// move all fish that were on counter "0" to new 6-day counter slot
			// and spawn an equal number of new lantern fish
			if counter == 0 {
				f[6] += num
				f[8] += num

				continue
			}

			// "decrease" counter -> move them to the previous slot
			f[counter-1] += num
		}
	}

	count := 0
	for _, v := range f {
		count += v
	}

	return count
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	line := strings.Split(strings.Split(string(p), "\n")[0], ",")
	dbg("lines: %#v", line)

	fish := NewFish(line)
	part1 = fish.simulateDays(80)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
