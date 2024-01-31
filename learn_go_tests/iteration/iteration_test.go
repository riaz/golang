package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		fmt.Errorf("expeted %q but got %q", repeated, expected)
	}
}

func BenchmarkReport(b *testing.B) {
	for i := 0; i < b.N; i++ { // this will run the code b.N times to check for latency
		Repeat("a", 4)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 3)
	fmt.Println(repeat)
	// Output: aaa
}
