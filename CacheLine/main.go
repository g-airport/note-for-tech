package main

import (
	"fmt"
	"sync/atomic"
)

//-----------------------------------------*
// CacheLineSize = 64
//-----------------------------------------*

type Atomic interface {
	IncreaseAllEles()
	IncreaseA()
	IncreaseB()
}

type NoPad struct {
	a uint64
	b uint64
	c uint64
}

func (atom *NoPad) IncreaseAllEles() {
	atomic.AddUint64(&atom.a, 1)
	atomic.AddUint64(&atom.b, 1)
	atomic.AddUint64(&atom.c, 1)
}

func (atom *NoPad) IncreaseA() {
	atomic.AddUint64(&atom.a, 1)
	//atomic.AddUint64(&atom.b, 1)
	//atomic.AddUint64(&atom.c, 1)
}

func (atom *NoPad) IncreaseB() {
	//atomic.AddUint64(&atom.a, 1)
	atomic.AddUint64(&atom.b, 1)
	//atomic.AddUint64(&atom.c, 1)
}

type Pad struct {
	a   uint64
	_p1 [8]uint64
	b   uint64
	_p2 [8]uint64
	c   uint64
	_p3 [8]uint64
}

func (atom *Pad) IncreaseAllEles() {
	atomic.AddUint64(&atom.a, 1)
	atomic.AddUint64(&atom.b, 1)
	atomic.AddUint64(&atom.c, 1)
}
func (atom *Pad) IncreaseA() {
	atomic.AddUint64(&atom.a, 1)
	//atomic.AddUint64(&atom.b, 1)
	//atomic.AddUint64(&atom.c, 1)
}
func (atom *Pad) IncreaseB() {
	//atomic.AddUint64(&atom.a, 1)
	atomic.AddUint64(&atom.b, 1)
	//atomic.AddUint64(&atom.c, 1)
}

func main() {
	np := &NoPad{}
	np.IncreaseAllEles()
	fmt.Println(np.a, np.b, np.c)
}
