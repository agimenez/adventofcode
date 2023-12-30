package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	debug int
)

func dbg(level int, fmt string, v ...interface{}) {
	if debug >= level {
		log.Printf(fmt+"\n", v...)
	}
}

func init() {
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.Parse()
}

type component struct {
	name string
	mult int
}

type reaction struct {
	mult  int // number of output produced
	input []component
}

type factory map[string]reaction

func toInt(s string) int {
	v, _ := strconv.Atoi(s)

	return v
}

func getReactions(in io.Reader) factory {
	scanner := bufio.NewScanner(in)
	f := factory{}

	for scanner.Scan() {
		line := scanner.Text()
		dbg(1, "Line: %v", line)
		parts := strings.Split(line, " => ")

		product := strings.Fields(parts[1])
		r := reaction{
			mult:  toInt(product[0]),
			input: []component{},
		}
		for _, p := range strings.Split(parts[0], ", ") {
			product := strings.Fields(p)
			r.input = append(r.input, component{
				name: product[1],
				mult: toInt(product[0]),
			})
		}

		f[product[1]] = r

		dbg(1, "Parts: %#v\n", parts)
		dbg(1, "factory: %v\n", f)

	}

	return f
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func (f factory) getORE(target string, mult int) int {
	spare := map[string]int{}
	need := map[string]int{
		"FUEL": 1,
	}

	for {
		if len(need) == 1 && need["ORE"] != 0 {
			return need["ORE"]
		}

		dbg(1, "=== NEW LOOP ===")
		for prod, q := range need {
			dbg(1, "Need %d %s", q, prod)
			if prod == "ORE" {
				continue
			}

			r := f[prod]
			qtyNeeded := int(math.Ceil(float64(q-spare[prod]) / float64(r.mult)))
			dbg(1, " -> Spare: %d, output: %d -> Need %d", spare[prod], r.mult, qtyNeeded)
			for _, c := range r.input {
				need[c.name] += qtyNeeded * c.mult
				dbg(1, " -> Added need: %d %s", qtyNeeded*c.mult, c.name)
			}

			spare[prod] += (r.mult * qtyNeeded) - q
			delete(need, prod)
			dbg(1, "Pending needs: %v", need)
		}
		dbg(1, "Pending need: %v", need)

	}

}

func main() {
	reactions := getReactions(os.Stdin)
	part1 := reactions.getORE("FUEL", 1)

	fmt.Printf("Reaction chains: %v\n", reactions)
	fmt.Printf("Part 1: %v\n", part1)
}
