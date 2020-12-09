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

type XMAS struct {
	window        int
	currentWindow map[int]int
	data          []int
	baseIndex     int
}

func NewXMAS(data []string, window int) *XMAS {
	x := &XMAS{
		window:        window,
		currentWindow: make(map[int]int, window),
		data:          make([]int, len(data)),
		baseIndex:     0,
	}

	for i := range data {
		val, err := strconv.Atoi(data[i])
		if err != nil {
			panic("error parsing input data")
		}
		x.data[i] = val
		if i < window {
			x.currentWindow[val] = i
		}
	}

	return x
}

func (x *XMAS) isNextInvalid() bool {
	next := x.next()
	for i := x.baseIndex; i < x.baseIndex+x.window; i++ {
		if idx, ok := x.currentWindow[next-x.data[i]]; ok && idx != i {
			return false
		}
	}

	return true
}

func (x *XMAS) slideWindow() {
	first := x.data[x.baseIndex]
	delete(x.currentWindow, first)
	x.baseIndex++

	last := x.data[x.baseIndex+x.window]
	x.currentWindow[last] = x.baseIndex + x.window
}

func (x *XMAS) next() int {
	return x.data[x.baseIndex+x.window]
}

func (x *XMAS) findInvalid() int {

	for {
		if x.isNextInvalid() {
			return x.next()
		}
		x.slideWindow()
	}
	return 0
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	x := NewXMAS(lines, 25)

	part1 = x.findInvalid()

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
