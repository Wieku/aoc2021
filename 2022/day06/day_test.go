package main

import "testing"

func BenchmarkP1Iteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p1()
	}
}

func BenchmarkP2Iteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p2()
	}
}

func BenchmarkP1Array(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p1A()
	}
}

func BenchmarkP2Array(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p2A()
	}
}

func BenchmarkP1Array2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p1A2()
	}
}

func BenchmarkP2Array2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p2A2()
	}
}

func BenchmarkP1Bits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p1B()
	}
}

func BenchmarkP2Bits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p2B()
	}
}
