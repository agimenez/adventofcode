package main

import (
	"fmt"

	"github.com/agimenez/adventofcode2019/intcode"
	"github.com/agimenez/adventofcode2019/utils"
)

type Queue []int

func (q Queue) Enqueue(d int) {
	q = append(q, d)
	utils.Dbg(2, " -> Enqueue: %d", d)
}

func (q Queue) Dequeue() int {
	if len(q) == 0 {
		utils.Dbg(2, " -> Dequeue: return -1")
		return -1
	}

	d := q[0]
	q = q[1:len(q)]
	utils.Dbg(2, " -> Dequeue: %d", d)
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

		}()

		nic.input <- i
		go func() {
			for {
				nic.input <- nic.packetQ.Dequeue()
				utils.Dbg(1, "NIC %d, dequeued", i)
			}
		}()
	}

	for {
		dst := <-n.output
		x := <-n.output
		y := <-n.output
		utils.Dbg(1, "Got packet (%d, %d) for %d", x, y, dst)
		n.nics[dst].packetQ.Enqueue(x)
		n.nics[dst].packetQ.Enqueue(y)
	}
}

func main() {

	var in string
	fmt.Scan(&in)

	n := newNetwork(in)
	n.Run()

}
