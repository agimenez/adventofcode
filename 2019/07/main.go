package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

var (
	debug int
)

func dbg(level int, fmt string, v ...interface{}) {
	if debug >= level {
		log.Printf(fmt, v...)
	}
}

type program struct {
	mem []int
	pc  int
}

func newProgram(p string) *program {

	pr := &program{
		mem: []int{},
		pc:  0,
	}

	pSlice := strings.Split(p, ",")
	for _, b := range pSlice {
		i, err := strconv.Atoi(b)
		if err != nil {
			log.Fatal(err)
		}

		pr.mem = append(pr.mem, i)
	}

	return pr
}

func (p *program) run(input <-chan int, output chan<- int) {

	for op := p.mem[p.pc]; op != 99; {
		dbg(3, "MEM = %v", p.mem)
		dbg(3, "pc = %d; op = %v", p.pc, op)

		opcode := op % 100
		dbg(3, "opcode = %v", opcode)

		switch opcode {
		case 1: // Add
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+4])
			a, b, c := p.fetchParameter(1), p.fetchParameter(2), p.mem[p.pc+3]
			dbg(3, " ADD %d %d -> %d", a, b, c)
			p.mem[c] = a + b
			p.pc += 4
		case 2: // Mul
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+4])
			a, b, c := p.fetchParameter(1), p.fetchParameter(2), p.mem[p.pc+3]
			dbg(3, " MUL %d %d -> %d", a, b, c)
			p.mem[c] = a * b
			p.pc += 4
		case 3: // In
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+2])
			var in, dst int
			in, dst = <-input, p.mem[p.pc+1]
			dbg(3, " IN  %d -> %d", in, dst)
			p.mem[dst] = in
			p.pc += 2
		case 4: // Out
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+2])
			src := p.fetchParameter(1)
			dbg(3, " OUT %d", src)

			output <- src
			p.pc += 2

		case 5: // JMP IF TRUE
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+3])
			tst, newpc := p.fetchParameter(1), p.fetchParameter(2)
			dbg(3, " JMP %d if %d != 0", newpc, tst)
			if tst != 0 {
				p.pc = newpc
			} else {
				p.pc += 3
			}

		case 6: // JMP IF FALSE
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+3])
			tst, newpc := p.fetchParameter(1), p.fetchParameter(2)
			dbg(3, " JMP %d if %d == 0", newpc, tst)
			if tst == 0 {
				p.pc = newpc
			} else {
				p.pc += 3
			}

		case 7: // LT
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+4])
			first, second, dst := p.fetchParameter(1), p.fetchParameter(2), p.mem[p.pc+3]
			dbg(3, " LT %d %d %d", first, second, dst)
			if first < second {
				p.mem[dst] = 1
			} else {
				p.mem[dst] = 0
			}

			p.pc += 4

		case 8: // EQ
			dbg(3, " INSTR = %v", p.mem[p.pc:p.pc+4])
			first, second, dst := p.fetchParameter(1), p.fetchParameter(2), p.mem[p.pc+3]
			dbg(3, " EQ %d %d %d", first, second, dst)
			if first == second {
				p.mem[dst] = 1
			} else {
				p.mem[dst] = 0
			}
			p.pc += 4

		default:
			log.Fatalf("Bad opcode = %v", op)
		}
		dbg(3, " MEM = %v", p.mem)

		op = p.mem[p.pc]
	}
}

func (p *program) fetchParameter(n int) int {
	opcode := p.mem[p.pc]
	parameter := p.mem[p.pc+n]
	mode := opcode / int(math.Pow10(n+1)) % 10

	dbg(3, "  (fetch) param[%d](%d) mode: %d", n, parameter, mode)
	if mode == 0 {
		return p.mem[parameter]
	}

	return parameter
}

// Heap's algorith
// Credits to https://stackoverflow.com/a/30226442/4735682
func slicePermutations(s []int) [][]int {
	perms := [][]int{}

	// fwd declaration to be able to call itself recursively
	var helper func(arr []int, n int)
	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(s))
			copy(tmp, s)
			perms = append(perms, tmp)
			return
		}

		for i := 0; i < n; i++ {
			helper(s, n-1)

			if n%2 == 1 {
				arr[i], arr[n-1] = arr[n-1], arr[i]
			} else {
				arr[0], arr[n-1] = arr[n-1], arr[0]
			}
		}
	}

	helper(s, len(s))

	return perms
}

func init() {
	flag.IntVar(&debug, "debug", 0, "debug level")
	flag.Parse()
}

func main() {

	var in string
	fmt.Scan(&in)

	perms := slicePermutations([]int{5, 6, 7, 8, 9})
	max := maxAmpsOutput(in, perms)

	log.Printf("Max output: %v", max)

}

func maxAmpsOutput(program string, perms [][]int) int {
	maxOutput := 0
	for _, perm := range perms {
		thrustersSignal := runPermutation(program, perm)
		dbg(1, "Got signal from perm %v: %v", perm, thrustersSignal)
		if thrustersSignal > maxOutput {
			maxOutput = thrustersSignal
		}
	}

	return maxOutput
}

func runPermutation(in string, phaseSettings []int) int {
	var channels []chan int

	// The first channel is special since the last loop of the last amp will write to
	// this channel (intended to be input of the first one), but in this case, the
	// first amp would have finished its work, thus leaving the channel blocked.
	// Buffer the first channel so the last amp can write to it and we can read from
	// it when exiting to provide the max value
	channels = append(channels, make(chan int, 1))

	for i := 1; i < len(phaseSettings); i++ {
		channels = append(channels, make(chan int))
	}

	// indicates that the whole system halted (need to wait for a signal from every
	// amp. If we don't sync with this, after the loop exits from launching the
	// goroutines, we'd get some random value from channels[0] while the system
	// didn't finish its whole processing.
	done := make(chan bool, len(phaseSettings))

	for i, ampSetting := range phaseSettings {
		program := newProgram(in)
		chanIn := channels[i]
		chanOut := channels[(i+1)%len(channels)]
		id := i
		go func() {
			dbg(1, "Spawn %d with input %v and output %v", id, chanIn, chanOut)
			program.run(chanIn, chanOut)
			dbg(1, "Spawn %d done", id)
			done <- true
		}()
		channels[i] <- ampSetting

	}
	channels[0] <- 0
	for range phaseSettings {
		<-done
	}

	return <-channels[0]
}
