package main

import "testing"

func BenchmarkP1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		enchance(2)
	}
}

func BenchmarkP2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		enchance(50)
	}
}
