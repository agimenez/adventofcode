package main

import (
	"flag"
	"fmt"
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
}

type network map[string]map[rune]string

func (n network) findNode(start string, ins []rune, match func(string) bool) []string {
	path := []string{}
	next := start

	for i := 0; !match(next); i = (i + 1) % len(ins) {
		dbg("Trying  %c (%d)", ins[i], i)

		path = append(path, next)
		next = n[next][ins[i]]
	}

	return path
}

func parseMap(s []string) ([]rune, network) {
	net := network{}

	ins := []rune(s[0])
	for _, l := range s[2:] {
		var node, left, right string
		fmt.Sscanf(l, "%s = (%3s, %3s)", &node, &left, &right)

		net[node] = map[rune]string{
			'L': left,
			'R': right,
		}

	}

	return ins, net
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func sliceLCM(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}
	return result
}

func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)
	instructions, network := parseMap(lines)
	dbg("ins: %s\nnetwork: %v", string(instructions), network)
	path := network.findNode("AAA", instructions, func(s string) bool { return s == "ZZZ" })
	part1 = len(path)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
