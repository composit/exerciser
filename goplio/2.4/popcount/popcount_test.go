package popcount_test

import (
	"fmt"
	"testing"

	"github.com/composit/exerciser/goplio/2.4/popcount"
)

const input = 0x1234567890ABCDEF

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(input)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountShift(input)
	}
}

func TestPopCountsAreEqual(t *testing.T) {
	var (
		first  = popcount.PopCount(input)
		second = popcount.PopCountShift(input)
	)

	if first != second {
		t.Fatalf("they're not the same: %d != %d", first, second)
	}

	fmt.Printf("they're the same: %d == %d\n", first, second)
}
