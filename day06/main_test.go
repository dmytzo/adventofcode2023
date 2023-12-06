package main

import "testing"

var testData = []string{
	"Time:      40   81   77   72",
	"Distance: 219 1012 1365 1089",
}

func BenchmarkTaskOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		taskOne(testData)
	}
}

func BenchmarkTaskTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		taskTwo(testData)
	}
}
