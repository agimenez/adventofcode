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
		dbg(2, "Line: %v", line)
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

		dbg(2, "Parts: %#v\n", parts)
		dbg(2, "factory: %v\n", f)

	}

	return f
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func (f factory) getORE(target string, mult int64) int64 {
	spare := map[string]int64{}
	need := map[string]int64{
		"FUEL": mult,
	}

	for {
		if len(need) == 1 && need["ORE"] != 0 {
			return need["ORE"]
		}

		dbg(2, "=== NEW LOOP ===")
		for prod, q := range need {
			dbg(2, "Need %d %s", q, prod)
			if prod == "ORE" {
				continue
			}

			r := f[prod]
			qtyNeeded := int64(math.Ceil(float64(q-spare[prod]) / float64(r.mult)))
			dbg(2, " -> Spare: %d, output: %d -> Need %d", spare[prod], r.mult, qtyNeeded)
			for _, c := range r.input {
				need[c.name] += qtyNeeded * int64(c.mult)
				dbg(2, " -> Added need: %d %s", qtyNeeded*int64(c.mult), c.name)
			}

			spare[prod] += (int64(r.mult) * qtyNeeded) - q
			delete(need, prod)
			dbg(2, "Pending needs: %v", need)
		}
		dbg(2, "Pending need: %v", need)

	}

}

func tryFuels(f factory, maxOre int64, orePerFuel int64) int64 {
	triedFuels := map[int64]bool{}
	tryFuel := float64(maxOre / orePerFuel)
	loops := 0
	for {
		dbg(1, "Trying with %f fuel", tryFuel)
		ore := f.getORE("FUEL", int64(tryFuel))
		dbg(1, " -> Got %v ore", ore)
		if _, ok := triedFuels[ore]; !ok {
			triedFuels[ore] = true
		} else if ore <= maxOre {
			//	return int64(tryFuel)
		}

		test := float64(maxOre / ore)
		dbg(1, " -> Ratio: %f", test)
		tryFuel *= test

		if loops > 2 {
			return int64(tryFuel)
		}
		loops++
	}
}

func main() {
	reactions := getReactions(os.Stdin)
	part1 := reactions.getORE("FUEL", 1)

	//part2 := tryFuels(reactions, 1e12, part1)
	part2 := tryFuels(reactions, 1000000000000, part1)

	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}
