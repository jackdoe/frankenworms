package main

import (
	"log"
	"math/rand"
	"time"
)

type receiver interface {
	ping(weight uint32)
	id() string
	activity() uint32
	connect(receiver, uint32)
}

type connection struct {
	to     receiver
	weight uint32
}

func (c connection) ping() {
	c.to.ping(c.weight)
}

type neuron struct {
	connections []*connection
	meta        []string
	threshold   uint32
	nActive     uint32
	nReceived   uint32
	wakeup      chan bool
}

func (n *neuron) connect(to receiver, weight uint32) {
	n.connections = append(n.connections, &connection{
		to:     to,
		weight: weight,
	})
}

func (n *neuron) activity() uint32 {
	return n.nActive
}

func (n *neuron) id() string {
	return n.meta[0]
}

func (n *neuron) ping(weight uint32) {
	n.nReceived += weight
	// select {
	// case n.wakeup <- true:
	// default:
	// }
}

func (n *neuron) ticker() {
	go func() {
		rand.Seed(time.Now().Unix())
		for {
			if n.nReceived > n.threshold {
				if DEBUG {
					log.Printf("%s activated, nReceived: %d", n.id(), n.nReceived)
				}

				n.nActive++
				for _, conn := range n.connections {
					conn.ping()
				}
			}
			n.nReceived = 0

			sleep := 50 + rand.Int31n(50)
			select {
			case <-n.wakeup:
			case <-time.After(time.Millisecond * time.Duration(sleep)):
			}
		}
	}()
}

func newEmptyNeuron(id string, meta []string, threshold uint32) *neuron {
	n := &neuron{
		connections: []*connection{},
		threshold:   threshold,
		meta:        meta,
		wakeup:      make(chan bool),
	}
	ID_TO_NEURON[id] = n
	return n
}
