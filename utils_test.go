package main

import (
	"bytes"
	"testing"
)

func TestSafeDecode(t *testing.T) {
	t.Run("decode 4s", func(t *testing.T) {
		input := []byte("R2hNVlhoVlVr")
		expect := []byte("GhMVXhVUk")
		if output := safeDecode(input); bytes.Compare(output, expect) != 0 {
			t.Errorf("Expect: %s, got: %s", expect, output)
		}
	})
	t.Run("decode none 4s", func(t *testing.T) {
		input := []byte("R2hNVlhoVlVrRQo")
		expect := []byte("GhMVXhVUkE\n")
		if output := safeDecode(input); bytes.Compare(output, expect) != 0 {
			t.Errorf("Expect: %s, got: %s", expect, output)
		}
	})
}
