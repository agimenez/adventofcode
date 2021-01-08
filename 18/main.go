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

var depth int = -1

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

type Stackable interface {
	Value() int
	Str() string
}

type Stack []Stackable

func (s Stack) Push(e Stackable) Stack {
	return append(s, e)
}

func (s Stack) Pop() (Stackable, Stack) {
	return s[len(s)-1], s[:len(s)-1]
}

func (s Stack) Empty() bool {
	return len(s) == 0
}

func (s Stack) Top() Stackable {
	return s[len(s)-1]
}

type Token string

func (t Token) Value() int {
	return 0
}

func (t Token) Str() string {
	return string(t)
}

type Number int

func NewNumber(s string) Number {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Cannot parse '%s", s)
	}

	return Number(n)
}
func (n Number) Value() int {
	return int(n)
}
func (n Number) Str() string {
	return fmt.Sprintf("%d", n)
}

type OpsMap map[string]struct {
	Precedence int
	Eval       func(a, b int) int
}

// Shunting-yard algorithm (https://en.wikipedia.org/wiki/Shunting-yard_algorithm)
func Solve2(tokens []string) (int, []string) {
	ops := Stack{}
	res := Stack{}
	opMap := OpsMap{
		"+": {1, Sum},
		"*": {2, Mul},
	}
	for _, tok := range tokens {
		switch tok {
		case "(":
			ops = ops.Push(Token(tok))
		case ")":
			top := ops.Top().Str()
			for top != "(" {

				var op, op1, op2 Stackable
				op, ops = ops.Pop()
				op1, res = res.Pop()
				op2, res = res.Pop()

				r := opMap[op.Str()].Eval(op1.Value(), op2.Value())
				res = res.Push(Number(r))

				top = ops.Top().Str()
			}
		case "+", "*":
			if !ops.Empty() {
				top := ops.Top().Str()
				for _, isOp := opMap[top]; isOp && opMap[top].Precedence >= opMap[tok].Precedence; {
					var op, op1, op2 Stackable

					op, ops = ops.Pop()
					op1, res = res.Pop()
					op2, res = res.Pop()

					r := opMap[op.Str()].Eval(op1.Value(), op2.Value())
					res = res.Push(Number(r))

					if ops.Empty() {
						break
					}

					top = ops.Top().Str()
				}
			}

			ops = ops.Push(Token(tok))

		default:
			res = res.Push(NewNumber(tok))
		}

	}

	for len(ops) > 0 {
		var op, op1, op2 Stackable

		op, ops = ops.Pop()
		op1, res = res.Pop()
		op2, res = res.Pop()

		r := opMap[op.Str()].Eval(op1.Value(), op2.Value())
		res = res.Push(Number(r))
	}
	dbg("Ops: %v", ops)
	dbg("Res: %v", res)

	return res[0].Value(), []string{}
}

func SolveExpression(e string, solve func([]string) (int, []string)) int {
	e = strings.ReplaceAll(e, "(", "( ")
	e = strings.ReplaceAll(e, ")", " )")
	tokens := strings.Split(e, " ")

	res, _ := solve(tokens)
	fmt.Printf("'%s' = %d\n", e, res)
	return res
}

func main() {

	part1, part2 := 0, 0
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		//part1 += SolveExpression(s.Text(), Solve)
		p2 := SolveExpression(s.Text(), Solve2)
		part2 += p2
	}

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
