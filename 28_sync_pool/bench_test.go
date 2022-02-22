package main

import (
	"sync"
	"testing"
)

type Counter struct {
	A int
	B int
}

func (c *Counter) Inc() {
	c.A += 1
	c.B += 1
}

var counterPool = sync.Pool{
	New: func() interface{} { return new(Counter) },
}

func BenchmarkWithoutPool(b *testing.B) {
	var c *Counter
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			c = new(Counter)
			c.Inc()
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var c *Counter
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			c = counterPool.Get().(*Counter)
			c.Inc()
			counterPool.Put(c)
		}
	}
}
