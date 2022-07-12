package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := repeated("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeated("a", 3)
	}
}

func ExampleRepeat() {
	str := repeated("f", 6)
	fmt.Println(str)
	// Output: ffffff
}
