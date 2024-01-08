package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected %q, got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	char := Repeat("a", 5)
	fmt.Println(char)
	// Output: aaaaa
}

func BenchmarkRepeated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
