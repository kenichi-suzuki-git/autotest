package main

import "testing"

// コメント追加
func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "even" {
		t.Errorf("Expected 'even', actual '%s'", result)
	}
}

