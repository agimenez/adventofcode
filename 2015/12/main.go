package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	. "github.com/agimenez/adventofcode/utils"
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

func solve1(s []string) int {
	res := 0

	numRE := regexp.MustCompile(`-?\d+`)

	nums := numRE.FindAllString(s[0], -1)
	for _, n := range nums {
		res += ToInt(n)

	}

	return res
}

func processDoc(doc any) int {
	res := 0

	switch v := doc.(type) {
	case int:
		dbg("INT -> %v", v)
		return v
	case float64:
		dbg("FLOAT64 -> %v", v)
		return int(v)
	case string:
		return 0
	case []any:
		for _, u := range v {
			res += processDoc(u)
		}
	case map[string]any:
		for _, vv := range v {
			if vv == "red" {
				return 0
			}

			res += processDoc(vv)
		}

	default:
		panic(fmt.Sprintf("Unknown type: %T (%v)", v, v))

	}

	return res
}

func solve2(s []string) int {
	res := 0

	var doc any
	err := json.Unmarshal([]byte(s[0]), &doc)
	if err != nil {
		panic(err)
	}

	res = processDoc(doc)

	return res
}
