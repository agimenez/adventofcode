package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

var (
	debug int
)

func dbg(level int, fmt string, v ...interface{}) {
	if debug >= level {
		log.Printf(fmt+"\n", v...)
	}
}

func init() {
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.Parse()
}

func getReactions(in io.Reader) int {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		dbg(1, "Line: %v", line)
		pattern := regexp.MustCompile(`(\d+) (\w+)(, (\d+) (\w))* => (\d+) (\w+)`)
		reacts := pattern.FindStringSubmatch(line)

		fmt.Printf("Reactions: %#v\n", reacts)

	}

	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {
	reactions := getReactions(os.Stdin)

	fmt.Printf("Reaction chains: %d\n", reactions)
}
