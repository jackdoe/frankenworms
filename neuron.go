package main

import "log"

type receiver interface {
	ping(receiver)
	tick()
	id() string
	activity() uint32
}

type neuron struct {
	peers     []receiver
	meta      []string
	threshold uint32
	nActive   uint32
	nReceived uint32
}

func (n *neuron) activity() uint32 {
	return n.nActive
}

func (n *neuron) id() string {
	return n.meta[0]
}

func (n *neuron) ping(from receiver) {
	n.nReceived++
	if from == nil {
		log.Printf("%s: pinged from external entity, nReceived: %d", n.id(), n.nReceived)
	} else {
		log.Printf("%s: pinged from %s, nReceived: %d", n.id(), from.id(), n.nReceived)
	}
}

func (n *neuron) tick() {
	if n.nReceived > n.threshold {
		n.nActive++
		for _, peer := range n.peers {
			peer.ping(n)
		}
	}
	n.nReceived = 0
}

func connect(id string, peer string) {
	n := ID_TO_NEURON[id]
	n.peers = append(n.peers, ID_TO_NEURON[peer])
}

func newEmptyNeuron(id string, meta []string, threshold uint32) *neuron {
	n := &neuron{
		peers:     []receiver{},
		threshold: threshold,
		meta:      meta,
	}
	ID_TO_NEURON[id] = n
	return n
}
