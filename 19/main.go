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

type Validator interface {
	Eval(string) (bool, string)
}

type Rule struct {
	set      *Matcher
	subrules [][]string
}

func NewRule(s string, m *Matcher) Rule {
	sets := strings.Split(s, " | ")
	r := Rule{
		set:      m,
		subrules: make([][]string, len(sets)),
	}
	for i := range sets {
		r.subrules[i] = []string{}
		for _, v := range strings.Split(sets[i], " ") {
			r.subrules[i] = append(r.subrules[i], v)
		}

	}

	return r
}
func (r Rule) Eval(in string) (bool, string) {
	if len(in) == 0 {
		return false, ""
	}

	for set := range r.subrules {
		rem := in
		match := true
		for _, rule := range r.subrules[set] {
			var ok bool
			if ok, rem = (*r.set)[rule].Eval(rem); !ok {
				match = false
				break
			}

		}

		if match {
			return true, rem
		}
	}

	return false, in
}

type Literal rune

func NewLiteral(s string) Literal {
	return Literal(s[1])
}

func (l Literal) Eval(in string) (bool, string) {
	if len(in) > 0 && in[0] == byte(l) {
		return true, in[1:]
	}

	return false, in
}

func NewValidator(s string, m *Matcher) Validator {
	if s[0] == '"' {
		return NewLiteral(s)
	}

	return NewRule(s, m)
}

type Matcher map[string]Validator

func (m Matcher) Match(in string) bool {
	ok, rem := m["0"].Eval(in)

	return ok && len(rem) == 0
}

func (m *Matcher) AddRule(s string) {
	parts := strings.Split(s, ": ")
	(*m)[parts[0]] = NewValidator(parts[1], m)
}

func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]

	m := Matcher{}
	l := 0
	for ; lines[l] != ""; l++ {
		m.AddRule(lines[l])
	}

	dbg("rules:%+v", &m)

	l++

	for _, l := range lines[l:] {
		if m.Match(l) {
			part1++
			dbg("%v matched!", l)
		} else {
			dbg("%v did not match", l)
		}

	}
	log.Printf("Part 1: %v\n", part1)

	m.AddRule("8: 42 | 42 8")
	m.AddRule("11: 42 31 | 42 11 31")
	for _, l := range lines[l:] {
		if m.Match(l) {
			part2++
			dbg("%v matched!", l)
		} else {
			dbg("%v did not match", l)
		}

	}

	log.Printf("Part 2: %v\n", part2)

}
