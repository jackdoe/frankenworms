package main

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

type muscle struct {
	pull    int32
	oldPull int32
	name    string
}

func (m *muscle) ping(weight int32) {
	atomic.AddInt32(&m.pull, weight)
}

func (m *muscle) id() string {
	return fmt.Sprintf("MUSCLE_%s", m.name)
}

func (m *muscle) delta() int32 {
	current := atomic.SwapInt32(&m.pull, 0)
	diff := current - m.oldPull
	m.oldPull = current
	return diff
}

func (m *muscle) activity() uint32 {
	return 1
}

func (m *muscle) connect(to receiver, weight int32) {
	panic("dont connect the *muscle to anything")
}

type body struct {
	left, right *muscle
}

func (b *body) ticker() {
	go func() {
		for {
			delta := b.left.delta()
			if delta > 0 {
				log.Printf("move left: %d", delta)
			}
			delta = b.right.delta()
			if delta > 0 {
				log.Printf("move right: %d", delta)
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()
}

func newBody() *body {
	return &body{
		left:  &muscle{name: "left"},
		right: &muscle{name: "right"},
	}
}
