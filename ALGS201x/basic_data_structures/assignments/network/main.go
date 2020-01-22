package main

import "fmt"

func main() {
	fmt.Println(simulate(1, 0, [][]int{}))
	fmt.Println(simulate(1, 1, [][]int{
		[]int{0, 0},
	}))
	fmt.Println(simulate(1, 2, [][]int{
		[]int{0, 1},
		[]int{0, 1},
	}))
	fmt.Println(simulate(1, 2, [][]int{
		[]int{0, 1},
		[]int{1, 1},
	}))
	fmt.Println(simulate(1, 3, [][]int{
		[]int{0, 2},
		[]int{1, 1},
		[]int{2, 1},
	}))

}

func simulate(bufferSize int, expectedPackets int, packets [][]int) []int {
	q := newQueue(bufferSize, expectedPackets)

	for index, pk := range packets {
		p := newPacket(pk[0], pk[1], index)
		q.takePacket(p)
	}

	for len(q.buffer) > 0 {
		q.clearCompletePackets(100)
	}

	return q.output
}

type packet struct {
	hasStarted  bool
	index       int
	arrivalTime int
	processTime int
}

func (p *packet) processOneTick() bool {
	if !p.hasStarted {
		p.hasStarted = true
	}
	p.processTime--
	return p.isFinished()
}

func (p *packet) isFinished() bool {
	return p.processTime <= 0
}
func newPacket(arrivalTime int, processTime int, index int) *packet {
	return &packet{
		false,
		index,
		arrivalTime,
		processTime,
	}
}

type queue struct {
	output      []int
	maxCapacity int
	currentTime int
	buffer      []*packet
}

func (q *queue) takePacket(p *packet) {
	q.clearCompletePackets(p.arrivalTime)

	if len(q.buffer) >= q.maxCapacity {
		q.output[p.index] = -1
	} else {
		q.buffer = append(q.buffer, p)
	}

}

func (q *queue) clearCompletePackets(nextArrivalTime int) {
	processingTime := nextArrivalTime - q.currentTime

	for {
		if len(q.buffer) == 0 {
			break
		}
		nextPacket := q.buffer[0]

		for i := 0; i < processingTime; processingTime-- {
			if !nextPacket.hasStarted {
				q.output[nextPacket.index] = nextArrivalTime - processingTime
			}

			if nextPacket.processOneTick() {
				break
			}
		}

		if nextPacket.isFinished() {
			q.buffer = q.buffer[1:]
		}

		if processingTime == 0 {
			break
		}
	}

	q.currentTime = nextArrivalTime
}

func newQueue(bufferSize int, expectedPackets int) *queue {
	return &queue{
		make([]int, expectedPackets, expectedPackets),
		bufferSize,
		0,
		[]*packet{},
	}
}
