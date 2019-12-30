package main

import (
	"fmt"

	"github.com/agimenez/adventofcode2019/intcode"
	"github.com/agimenez/adventofcode2019/utils"
)

type Queue []int

func (q Queue) Enqueue(d int) {
	q = append(q, d)
}

func (q Queue) Dequeue() int {
	if len(q) == 0 {
		return -1
	}

	d := q[0]
	q = q[1:len(q)]
	return d
}

type NIC struct {
	cpu     *intcode.Program
	input   chan int
	packetQ Queue
}

type Network struct {
	nics   []NIC
	output chan int
}

func newNetwork(code string) Network {
	n := Network{
		output: make(chan int),
	}
	for i := 0; i < 50; i++ {
		nic := NIC{
			cpu:     intcode.NewProgram(code),
			input:   make(chan int),
			packetQ: Queue{},
		}

		n.nics = append(n.nics, nic)
	}

	return n
}

func (n *Network) Run() {
	for i, nic := range n.nics {
		go func() {
			nic.cpu.Run(nic.input, n.output)
			nic.input <- i
		}()

		go func() {
			for {
				select {
				case nic.input <- nic.packetQ.Dequeue():
				default:

				}
			}
		}()
	}

	for {
		c := <-n.output
		utils.Dbg(1, "Got packet for %d", c)
		n.nics[c].input <- (<-n.output)
		n.nics[c].input <- (<-n.output)
	}
}

func main() {

	var in string
	fmt.Scan(&in)

	n := newNetwork(in)
	n.Run()

}
