package main

import (
	"testing"
)

func BenchmarkDayOneTaskOne(b *testing.B) {
	testData := []string{"543cshpxrfnnhonetkbhxtmlgczdndqjscb2mpftseven44five8nineeightrmgrljrljb8hxxfmdpbbvmblltxfive6mdsmm7mmpknprsix6five4vkqrsixmjkjps"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		taskOne(testData)
	}
}

func BenchmarkDayOneTaskTwo(b *testing.B) {
	testData := []string{"543cshpxrfnnhonetkbhxtmlgczdndqjscb2mpftseven44five8nineeightrmgrljrljb8hxxfmdpbbvmblltxfive6mdsmm7mmpknprsix6five4vkqrsixmjkjps"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		taskTwo(testData)
	}
}
