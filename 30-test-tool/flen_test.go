package main

import (
	"fmt"
	"testing"
)

var blackhole int

func TestFileLen(t *testing.T) {
	result, err := FileLen("./coverage.html", 1)
	if err != nil {
		t.Fatal(err)
	}

	if result != 2540 {
		t.Error("Expected 2540, got", result)
	}
}

func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := FileLen("/home/simon/bin/kubectl", v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}
