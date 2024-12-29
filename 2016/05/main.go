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

func hackHash(secret string, zeros int) string {
	i := 0
	zeroCmp := strings.Repeat("0", zeros)
	var pwd strings.Builder
	for pwd.Len() < 8 {
		padding := strconv.Itoa(i)
		hash := md5.Sum([]byte(secret + padding))
		str := hex.EncodeToString(hash[:])
		if strings.HasPrefix(str, zeroCmp) {
			pwd.WriteByte(str[5])
			dbg("Found (%d): hash: %v, pwd: %v", i, str, pwd.String())
		}
		i++
	}
	dbg("FINAL: %d = %v", i, pwd.String())

	return pwd.String()

}

func hackHash2(secret string, zeros int) string {
	i := 0
	zeroCmp := strings.Repeat("0", zeros)
	pwd := [8]rune{}
	found := 0
	for i = 0; found < 8; i++ {
		padding := strconv.Itoa(i)
		hash := md5.Sum([]byte(secret + padding))
		str := hex.EncodeToString(hash[:])
		if strings.HasPrefix(str, zeroCmp) {
			dbg("CANDIDATE (%d): hash: %v, pwd: %q", i, str, string(pwd[:]))
			pos := str[5] - '0'

			// position out of bounds
			if pos > 7 {
				continue
			}

			// already set
			if pwd[pos] != 0 {
				dbg(" -> pos %d already set!", pos)
				continue
			}

			pwd[pos] = rune(str[6])
			dbg("Found (%d): hash: %v, pwd: %q", i, str, string(pwd[:]))
			found++
		}
	}
	dbg("FINAL: %d = %v", i, string(pwd[:]))

	return string(pwd[:])

}

func solve(lines []string) (string, string, time.Duration, time.Duration) {
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

func solve1(s []string) string {
	res := hackHash(s[0], 5)

	return res
}

func solve2(s []string) string {
	res := hackHash2(s[0], 5)

	return res
}
