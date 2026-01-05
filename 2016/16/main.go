package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
	// . "github.com/agimenez/adventofcode/utils"
)

var (
	debug bool
)

func dbg(f string, v ...interface{}) {
	if debug {
		fmt.Printf(f+"\n", v...)
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
}

func main() {
	flag.Parse()

	p, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]
	//dbg("lines: %#v", lines)

	part1, part2, dur1, dur2 := solve(lines)
	log.Printf("Part 1 (%v): %v\n", dur1, part1)
	log.Printf("Part 2 (%v): %v\n", dur2, part2)

}

func solve(lines []string) (int, int, time.Duration, time.Duration) {
	var now time.Time
	var dur [2]time.Duration

	now = time.Now()
	part1 := solve1(lines)
	dur[0] = time.Since(now)

	now = time.Now()
	part2 := solve2(lines)
	dur[1] = time.Since(now)

	return part1, part2, dur[0], dur[1]

}

func DragonStep(s string) string {
	var b bytes.Buffer
	b.WriteString(s)
	b.WriteRune('0')

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' {
			b.WriteRune('0')
		} else {
			b.WriteRune('1')
		}
	}

	return b.String()
}

func DragonCurve(s string, length int) string {
	var data string
	for data = s; len(data) < length; data = DragonStep(data) {
		dbg("Data len: %v", len(data))
	}

	return data[:length]
}

func CheckSum(s string) string {
	dbg("Checksum len: %v", len(s))
	if len(s)%2 != 0 {
		return s
	}

	var b bytes.Buffer
	for i := 0; i < len(s)-1; i += 2 {
		if s[i] == s[i+1] {
			b.WriteRune('1')
		} else {
			b.WriteRune('0')
		}
	}

	return CheckSum(b.String())
}

func solve1(s []string) int {
	res := 0

	data := DragonCurve(s[0], 272)
	fmt.Println(CheckSum(data))

	return res
}

func solve2(s []string) int {
	res := 0

	data := DragonCurve(s[0], 35651584)
	fmt.Println(CheckSum(data))

	return res
}
