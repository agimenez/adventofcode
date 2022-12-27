package main

import (
	"io/ioutil"
	"log"
	"os"
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

//func init() {
//	flag.BoolVar(&debug, "debug", false, "enable debug")
//	flag.Parse()
//}

type compartment map[rune]bool

type rucksack [2]compartment

func NewRuckSack(in string) rucksack {
	r := rucksack{
		compartment{},
		compartment{},
	}

	for i, v := range in {
		slot := i / (len(in) / 2)
		r[slot][v] = true
	}

	return r
}

func (r rucksack) getDuplicate() rune {
	for k := range r[0] {
		if _, ok := r[1][k]; ok {
			return k
		}
	}

	return 0
}

func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	} else {
		return int(r - 'A' + 27)
	}

}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	for _, line := range lines {
		r := NewRuckSack(line)
		d := r.getDuplicate()
		part1 += priority(d)
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
