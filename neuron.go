package main

import (
	"log"
	"math/rand"
	"time"
)

type receiver interface {
	ping(weight int32)
	id() string
	activity() uint32
	connect(receiver, int32)
}

type connection struct {
	to     receiver
	weight int32
}

func (c connection) ping() {
	c.to.ping(c.weight)
}

type neuron struct {
	connections []*connection
	meta        []string
	threshold   int32
	nActive     uint32
	nReceived   int32
}

func (n *neuron) connect(to receiver, weight int32) {
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

func (n *neuron) ping(weight int32) {
	n.nReceived += weight
}

func (n *neuron) ticker() {
	go func() {
		rand.Seed(time.Now().Unix())
		for {
			if n.nReceived >= n.threshold {
				if DEBUG {
					log.Printf("%s activated, nReceived: %d", n.id(), n.nReceived)
				}

				n.nActive++
				for _, conn := range n.connections {
					conn.ping()
				}
			}
			n.nReceived = 0
			time.Sleep(200 * time.Millisecond)
		}
	}()
}

func newEmptyNeuron(id string, meta []string, threshold int32, b *body) *neuron {
	n := &neuron{
		connections: []*connection{},
		threshold:   threshold,
		meta:        meta,
	}
	ID_TO_NEURON[id] = n

	// Anterior harsh body touch: "FLPL", "FLPR", "BDUL", "BDUR", "SDQR":
	// Posterior harsh body touch: "PVDL", "PVDR", "PVCL", "PVCR":
	// Nose touch: "ASHL", "ASHR", "FLPL", "FLPR", "OLQDL", "OLQDR", "OLQVL", "OLQVR"
	// Food: "ADFL", "ADFR", "ASGL", "ASGR", "ASIL", "ASIR", "ASJL", "ASJR", "AWCL", "AWCR", "AWAL", "AWAR"

	return n
}
