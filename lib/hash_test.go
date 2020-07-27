package lib

import (
	"testing"
)

func Test_Murmur32(t *testing.T) {
	h := Murmur32([]byte("hello world"))
	if h != 4008393376 {
		t.Errorf("Hash %d is not equal to %d", h, 1815237614)
	}
}

func Benchmark_Murmur32(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		Murmur32([]byte("hello world"))
	}
}
