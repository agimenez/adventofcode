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
}

type part struct {
	x, m, a, s int
}

type cmpFunc func(part) bool
type step struct {
	cmp  cmpFunc
	dest string
}

type workflow map[string][]step

func parseWF(s string) (string, []string) {
	name, stepStr, _ := strings.Cut(s, "{")
	steps := strings.Split(strings.TrimSuffix(stepStr, "}"), ",")
	dbg("Workflow %q: name: %q, rest: %q", s, name, steps)

	return name, steps

}

func gt(a, b int) bool {
	return a > b
}

func lt(a, b int) bool {
	return a < b
}

func toInt(s string) int {
	v, _ := strconv.Atoi(s)

	return v
}

var op2func = map[byte]func(int, int) bool{
	'>': gt,
	'<': lt,
}

func parseStep(in string) step {
	parts := strings.Split(in, ":")

	// Terminal step: always jump to dest
	if len(parts) == 1 {
		return step{
			dest: parts[0],
			cmp:  func(part) bool { return true },
		}
	}

	opIndex := strings.IndexAny(parts[0], "<>")
	field := parts[0][:opIndex]
	op := parts[0][opIndex]
	cmpValue := parts[0][opIndex+1:]
	dbg("creating step for %v -> %v (%s %c %s)", parts[0], parts[1], field, op, cmpValue)
	val := toInt(cmpValue)
	cmpFun := func(p part) bool {
		switch field {
		case "x":
			return op2func[op](p.x, val)
		case "m":
			return op2func[op](p.m, val)
		case "a":
			return op2func[op](p.a, val)
		case "s":
			return op2func[op](p.s, val)

		}

		return false
	}

	return step{
		cmp:  cmpFun,
		dest: parts[1],
	}
}

func (w workflow) addWorkflow(wf string) {
	name, steps := parseWF(wf)
	w[name] = []step{}

	for _, s := range steps {
		w[name] = append(w[name], parseStep(s))
	}
}

func (w workflow) run(p part, start string) bool {
	dbg("Workflow %s", start)
	for _, s := range w[start] {
		if s.cmp(p) {
			switch s.dest {
			case "A":
				return true
			case "R":
				return false
			}

			return w.run(p, s.dest)
		}
	}

	return false
}

func parsePart(s string) part {
	p := part{}
	s = strings.Trim(s, "{}")

	for _, value := range strings.Split(s, ",") {
		keyVal := strings.Split(value, "=")
		v := toInt(keyVal[1])
		switch keyVal[0] {
		case "x":
			p.x = v
		case "m":
			p.m = v
		case "a":
			p.a = v
		case "s":
			p.s = v
		}
	}

	return p
}

func (p part) sumRatings() int {
	return p.x + p.m + p.a + p.s
}

func main() {
	flag.Parse()

	part1, part2 := 0, 0
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input")
	}
	lines := strings.Split(string(p), "\n\n")
	w := workflow{}
	for _, wf := range strings.Split(lines[0], "\n") {
		dbg("Workflow: %q", wf)
		w.addWorkflow(wf)

	}
	dbg("Workflow: %v", w)

	lines[1] = lines[1][:len(lines[1])-1]
	for _, part := range strings.Split(lines[1], "\n") {
		dbg("Running part %v", part)
		p := parsePart(part)
		if w.run(p, "in") {
			part1 += p.sumRatings()
		}
	}

	//dbg("lines: %#v", lines)

	log.Printf("Part 1: %v\n", part1)
	log.Printf("Part 2: %v\n", part2)

}
