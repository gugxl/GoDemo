package main

import "testing"

func TestLoop(t *testing.T) {
	t.Logf("Loop:%v", Loop(uint64(32)))
}

func TestFactorial(t *testing.T) {
	t.Logf("Factorial:%v", Factorial(uint64(32)))
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop(uint64(40))
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(uint64(40))
	}
}

// 当前包下执行 go test --test.bench=.* 即可
