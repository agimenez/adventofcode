package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	// . "github.com/agimenez/adventofcode/utils"
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

func hackHash(secret string) int {
	i := 0
	for {
		dbg("Trying %d", i)
		padding := strconv.Itoa(i)
		hash := md5.Sum([]byte(secret + padding))
		str := hex.EncodeToString(hash[:])
		if strings.HasPrefix(str, "00000") {
			break
		}
		i++
	}

	return i

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

func solve1(s []string) int {
	res := hackHash(s[0])

	return res
}

func solve2(s []string) int {
	res := 0

	return res
}
