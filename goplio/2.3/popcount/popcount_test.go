package popcount_test

import (
	"fmt"
	"testing"

	"github.com/composit/exerciser/goplio/2.3/popcount"
)

const input = 0x1234567890ABCDEF

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(input)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(input)
	}
}

func TestPopCountsAreEqual(t *testing.T) {
	var (
		first  = popcount.PopCount(input)
		second = popcount.PopCountLoop(input)
	)

	if first != second {
		t.Fatalf("they're not the same: %d != %d", first, second)
	}

	fmt.Printf("they're the same: %d == %d\n", first, second)
}
