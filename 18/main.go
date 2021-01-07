package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

var (
	debug bool
)

func dbg(fmt string, v ...interface{}) {
	if debug {
		log.Printf(fmt, v...)
	}
}

// This is disappointing from Go :(
// https://stackoverflow.com/a/58192326/4735682
var _ = func() bool {
	testing.Init()
	return true
}()

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug")
	flag.Parse()
}

func Sum(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

var depth int

func Solve(tokens []string) (int, []string) {
	acc := 0
	rem := tokens

	op := Sum

	for len(rem) > 0 {
		dbg("%*sTokens: %v", 3*depth, " ", rem)
		dbg("%*stoken: %v", 3*depth, " ", rem[0])
		switch rem[0] {
		case "+":
			op = Sum
		case "*":
			op = Mul
		case "(":
			depth++
			dbg("%*sCalling Solve on %v", 3*depth, " ", rem)
			num, remaining := Solve(rem[1:])
			dbg("%*sReturn remaining: %v", 3*depth, " ", remaining)
			acc = op(acc, num)
			rem = remaining
		case ")":
			depth--
			return acc, rem
		default:
			num, err := strconv.Atoi(rem[0])
			if err != nil {
				fmt.Errorf("Can't parse %v", rem[0])
			}

			acc = op(acc, num)

		}

		rem = rem[1:]
	}

	return acc, []string{}
}

func SolveExpression(e string) int {
	e = strings.ReplaceAll(e, "(", "( ")
	e = strings.ReplaceAll(e, ")", " )")
	tokens := strings.Split(e, " ")

	res, _ := Solve(tokens)
	return res
}

func main() {

	part1, part2 := 0, 0
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		part1 += SolveExpression(s.Text())
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
