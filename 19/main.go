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
	Eval(string) (bool, []string)
}

type Rule struct {
	id       string
	set      *Matcher
	subrules [][]string
}

var depth int = -1

func NewRule(s string, m *Matcher, id string) Rule {
	sets := strings.Split(s, " | ")
	r := Rule{
		id:       id,
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
func (r Rule) Eval(in string) (bool, []string) {
	depth++
	defer func() { depth-- }()

	dbg("")
	dbg("%*s[%2s] Input '%s'. Rule eval: %+v", 3*depth, " ", r.id, in, r.subrules)
	if len(in) == 0 {
		return false, nil
	}

	match := false
	var groupRemaining []string

	for set := range r.subrules {
		possibleRemaining := []string{in}
		dbg("%*s[%2s] Subgroup: %+v", 3*depth, " ", r.id, r.subrules[set])

		// For every rule, we use the previous rules' output
		groupMatch := true
		for _, rule := range r.subrules[set] {
			// Remainings to check agains the same group's next rule
			var nextRem []string

			ruleMatch := false
			for _, str := range possibleRemaining {
				ok, rem := (*r.set)[rule].Eval(str)
				if !ok {
					continue
				}
				dbg("%*s[%2s]     -> Subrule %s match: possible rem = '%s'", 3*depth, " ", r.id, rule, rem)
				nextRem = append(nextRem, rem...)
				ruleMatch = true
			}

			possibleRemaining = nextRem
			dbg("%*s[%2s]  -> possible rem for next rule = '%v'", 3*depth, " ", r.id, possibleRemaining)
			if !ruleMatch {
				groupMatch = false
			}

		}
		if groupMatch {
			match = true
		}

		groupRemaining = append(groupRemaining, possibleRemaining...)

	}

	dbg("%*s[%2s] RET %v '%s'", 3*depth, " ", r.id, match, groupRemaining)
	return match, groupRemaining
}

type Literal rune

func NewLiteral(s string) Literal {
	return Literal(s[1])
}

func (l Literal) Eval(in string) (bool, []string) {
	depth++
	defer func() { depth-- }()
	rid := "4"
	if string(l) == "a" {
		rid = "3"
	}

	dbg("%*s[%2s] Literal eval", 3*depth, " ", rid)
	if len(in) > 0 && in[0] == byte(l) {
		//dbg("%*s -> Match! rem = '%s'", 3*depth, " ", in[1:])
		dbg("%*s[%2s] Outpt '%s'", 3*depth, " ", rid, in[1:])
		return true, []string{in[1:]}
	}

	//dbg("%*s  -> NO Match", 3*depth, " ")
	dbg("%*s[%2c] NOMCH '%s'", 3*depth, " ", l, in)
	return false, nil
}

func NewValidator(s string, m *Matcher, id string) Validator {
	if s[0] == '"' {
		return NewLiteral(s)
	}

	return NewRule(s, m, id)
}

type Matcher map[string]Validator

func (m Matcher) Match(in string) bool {
	ok, rem := m["0"].Eval(in)

	// The first entry in the rem slice will have the shortest string not consumed.
	// If its length is 0, then we have a match (all input chars are consumed).
	//log.Printf("Match '%s' returned %v %v (%v)", in, ok, rem, len(rem))
	return ok && len(rem[0]) == 0
}

func (m *Matcher) AddRule(s string) {
	parts := strings.Split(s, ": ")
	(*m)[parts[0]] = NewValidator(parts[1], m, parts[0])
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

	l++

	for _, l := range lines[l:] {
		if m.Match(l) {
			part1++
		}

	}
	log.Printf("Part 1: %v\n", part1)

	m.AddRule("8: 42 | 42 8")
	m.AddRule("11: 42 31 | 42 11 31")
	for _, l := range lines[l:] {
		if m.Match(l) {
			part2++
		}

	}

	log.Printf("Part 2: %v\n", part2)

}
