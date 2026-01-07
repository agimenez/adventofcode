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

func OpSwapPos(pass []byte, p1, p2 int) []byte {

	pass[p1], pass[p2] = pass[p2], pass[p1]
	return pass
}

func RotateRelative(pass []byte, c byte) []byte {
	idx := bytes.IndexByte(pass, c)
	rot := idx + 1
	if idx >= 4 {
		rot++
	}
	rot = rot % len(pass)

	return RotateImmediate(pass, -rot)
}

func RotateImmediate(pass []byte, n int) []byte {
	split := n
	if n < 0 {
		split = len(pass) + n
	}

	var out bytes.Buffer
	out.Write(pass[split:])
	out.Write(pass[0:split])

	return out.Bytes()
}

func Reverse(pass []byte, from, to int) []byte {
	for ; from < to; from, to = from+1, to-1 {
		pass[from], pass[to] = pass[to], pass[from]
	}

	return pass
}

func Move(pass []byte, from, to int) []byte {
	var out bytes.Buffer
	dbg("MOVE: %q %d -> %d", string(pass), from, to)
	if from < to {
		out.Write(pass[:from])
		out.Write(pass[from+1 : to+1])
		out.WriteByte(pass[from])
		out.Write(pass[to+1:])
	} else {
		out.Write(pass[:to])
		out.WriteByte(pass[from])
		out.Write(pass[to:from])
		out.Write(pass[from+1:])
	}

	return out.Bytes()
}

func ApplyOp(s string, pass []byte) []byte {

	parts := strings.Fields(s)
	switch parts[0] {
	case "swap":
		var p1, p2 int
		if parts[1] == "position" { //swap position 4 with position 0
			p1 = ToInt(parts[2])
			p2 = ToInt(parts[5])

		} else { // swap letter d with letter b
			p1 = bytes.IndexByte(pass, parts[2][0])
			p2 = bytes.IndexByte(pass, parts[5][0])
		}
		return OpSwapPos(pass, p1, p2)

	case "rotate":
		// rotate based on position of letter b
		if parts[1] == "based" {
			return RotateRelative(pass, parts[6][0])
		}
		//rotate left 1 step
		rot := ToInt(parts[2])
		if parts[1] == "right" {
			rot = -rot
		}

		return RotateImmediate(pass, rot)

	case "reverse":
		// reverse positions 0 through 4
		return Reverse(pass, ToInt(parts[2]), ToInt(parts[4]))
	case "move":
		// move position 1 to position 4
		return Move(pass, ToInt(parts[2]), ToInt(parts[5]))

	default:
		panic("Unknown operation: " + parts[0])
	}
}

func Scramble(s []string, pass string) string {
	p := []byte(pass)

	dbg("== Scramble %q", string(p))
	for _, op := range s {
		p = ApplyOp(op, p)
		dbg("  > %q", string(p))
	}

	return string(p)
}

func solve1(s []string) int {
	res := 0

	fmt.Println(Scramble(s, "abcdefgh"))

	dbg("")
	return res
}

func solve2(s []string) int {
	res := 0
	dbg("========== PART 2 ===========")

	target := "fbgdceah"
	for pass := range Permutations([]byte("abcdefgh")) {
		scrambled := Scramble(s, string(pass))
		if scrambled == target {
			fmt.Println(string(pass))
			break
		}
	}

	return res
}
