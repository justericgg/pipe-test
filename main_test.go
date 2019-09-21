package main

import "testing"

func increase(n int) int {
	return n + 1
}

func square(n int) int {
	return n * n
}

func TestPipe(t *testing.T) {

	t.Run("Test increase twice", func(t *testing.T) {

		number := 1
		want := 3

		got := pipe(number, increase, increase)

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("Test increase and then square it", func(t *testing.T) {
		number := 2
		want := 9

		got := pipe(number, increase, square)

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
