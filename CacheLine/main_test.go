package main

import (
	"reflect"
	"runtime"
	"sync"
	"testing"
)

func TestUint64(t *testing.T) {
	i := uint64(1)
	typ := reflect.TypeOf(i)
	t.Log(typ.Size())
	t.Log(typ.Align())
}

func TestNoPad(t *testing.T) {
	i := NoPad{}
	typ := reflect.TypeOf(i)
	t.Log(typ.Size())
	t.Log(typ.Align())
}

func TestPad(t *testing.T) {
	i := Pad{}
	typ := reflect.TypeOf(i)
	t.Log(typ.Size())
	t.Log(typ.Align())
	n := typ.NumField()
	for j := 0; j < n; j++ {
		t.Log("Name:", typ.Field(j).Name,
			"Type:", typ.Field(j).Type,
			"Tag:", typ.Field(j).Tag,
			"Offset", typ.Field(j).Offset,
			"Size", typ.Field(j).Type.Size(),
			"Align", typ.Field(j).Type.Align())
	}
}

func testAtomicIncrease(Atomic Atomic) {
	runtime.GOMAXPROCS(4)
	paraNum := 10000
	addTimes := 10000
	var wg sync.WaitGroup
	wg.Add(paraNum * 2)
	for i := 0; i < paraNum; i++ {
		go func() {
			for j := 0; j < addTimes; j++ {
				Atomic.IncreaseA()
			}
			wg.Done()
		}()
	}
	for i := 0; i < paraNum; i++ {
		go func() {
			for j := 0; j < addTimes; j++ {
				Atomic.IncreaseB()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkNoPad(b *testing.B) {
	Atomic := &NoPad{}
	b.ResetTimer()
	testAtomicIncrease(Atomic)
}

func BenchmarkPad(b *testing.B) {
	Atomic := &Pad{}
	b.ResetTimer()
	testAtomicIncrease(Atomic)
}
