package main

import "testing"

func BenchmarkP1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p1()
	}
}

func BenchmarkP2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p2()
	}
}
