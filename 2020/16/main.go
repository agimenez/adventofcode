package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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

type Range struct {
	min int
	max int
}

type Validator []Range

func NewValidator(rule string) Validator {
	v := Validator{}
	ranges := strings.Split(rule, " or ")
	for _, r := range ranges {
		vals := strings.Split(r, "-")
		min, err := strconv.Atoi(vals[0])
		if err != nil {
			log.Fatalf("Could not parse min %v", vals[0])
		}

		max, err := strconv.Atoi(vals[1])
		if err != nil {
			log.Fatalf("Could not parse max %v", vals[1])
		}

		v = append(v, Range{min: min, max: max})
	}

	dbg("Validator: %#v", v)

	return v
}

func (v Validator) Valid(num int) bool {
	for _, r := range v {
		if r.min <= num && num <= r.max {
			return true
		}
	}

	return false
}

type Ticket []int

func NewTicket(fields string) Ticket {
	parts := strings.Split(fields, ",")
	t := Ticket{}
	for _, f := range parts {
		n, err := strconv.Atoi(f)
		if err != nil {
			log.Fatalf("Can't parse field %v", f)
		}

		t = append(t, n)
	}

	return t
}

func InvalidValues(rules map[string]Validator, t Ticket) []int {
	invalid := []int{}
	dbg("checking ticket %#v", t)
	for _, f := range t {
		dbg(" -> %v", f)
		valid := false
		for _, r := range rules {
			dbg("  -> rule %#v", r)
			if r.Valid(f) {
				valid = true
				break
			}
		}
		if !valid {
			invalid = append(invalid, f)
		}
	}

	return invalid
}

func Sum(nums []int) int {
	s := 0
	for i := range nums {
		s += nums[i]
	}

	return s
}

func CopyRules(rules map[string]Validator) map[string]Validator {
	c := make(map[string]Validator, len(rules))

	for k, v := range rules {
		c[k] = v
	}

	return c
}

// This is not the most efficient since we could ignore fields already processed
func GetRuleCandidates(field int, rules map[string]Validator, tickets []Ticket) []string {
	c := []string{}

RulesLoop:
	for rule, v := range rules {
		for _, ticket := range tickets {
			if !v.Valid(ticket[field]) {
				continue RulesLoop
			}
		}

		c = append(c, rule)
	}

	return c
}
func main() {

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n")
	lines = lines[:len(lines)-1]

	rules := make(map[string]Validator)
	l := 0
	for ; lines[l] != ""; l++ {
		parts := strings.Split(lines[l], ": ")
		rules[parts[0]] = NewValidator(parts[1])
	}

	//jump newline + "your ticket"
	l += 2
	myTicket := NewTicket(lines[l])

	//jump my ticket + newline + "nearby tickets"
	validNearby := []Ticket{}
	for _, l := range lines[l+3:] {
		t := NewTicket(l)
		inv := InvalidValues(rules, t)
		part1 += Sum(inv)

		if len(inv) == 0 {
			validNearby = append(validNearby, t)
		}

	}

	log.Printf("Part 1: %v\n", part1)

	remaining := CopyRules(rules)
	mapping := map[string]int{}
	for len(remaining) > 0 {
		for field := 0; field < len(myTicket); field++ {
			dbg("Field %d", field)
			c := GetRuleCandidates(field, remaining, validNearby)
			dbg("Got candidates: %#v", c)

			if len(c) == 1 {
				mapping[c[0]] = field
				delete(remaining, c[0])
			}
		}
	}

	// we could do the calculations inline in the previous loop, but I like to
	// separate the different logics
	dbg("mapping: %#v", mapping)
	part2 = 1
	for field := range rules {
		if strings.HasPrefix(field, "departure") {
			part2 *= myTicket[mapping[field]]
		}
	}

	log.Printf("Part 2: %v\n", part2)

}
