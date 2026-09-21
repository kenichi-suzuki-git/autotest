package main

import "testing"

// コメント修正
func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "even" {
		t.Errorf("Expected 'even', actual '%s'", result)
	}
}

