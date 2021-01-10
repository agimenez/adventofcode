package main

import (
	"flag"
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

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
	flag.Parse()
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	dbg("lines: %+v", lines)

	l := 0
	for ; lines[l] != ""; l++ {
		parts := strings.Split(lines[l], ": ")
		//rules[parts[0]] = NewValidator(parts[1])
		dbg("%+v", parts)
	}

	l++

	//jump my ticket + newline + "nearby tickets"
	for _, l := range lines[l:] {
		dbg("%v", l)

	}

	log.Printf("Part 1: %v\n", part1)

	log.Printf("Part 2: %v\n", part2)

}
