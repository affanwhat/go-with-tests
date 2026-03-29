package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat 'c' 50 times", func(t *testing.T) {
		repeated := Repeat("c", 50)
		expected := "cccccccccccccccccccccccccccccccccccccccccccccccccc"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat 'd' 0 times", func(t *testing.T) {
		repeated := Repeat("d", 0)
		expected := ""
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat 'e' -1 times", func(t *testing.T) {
		repeated := Repeat("e", -1)
		expected := ""
		assertCorrectMessage(t, repeated, expected)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat (b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a",5)
	fmt.Println(repeated)
	// Output: aaaaa
}