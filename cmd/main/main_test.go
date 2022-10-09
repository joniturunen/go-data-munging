package main

import (
	"testing"
)

func Test_example(t *testing.T) {
	t.Run("Constructed with one parameter", func(t *testing.T) {
		got := "Assertion"
		expected := "Expected result"

		if got != expected {
			t.Errorf("GOT: %q\nEXPECTED: %q\n", got, expected)
		}
	})
}

