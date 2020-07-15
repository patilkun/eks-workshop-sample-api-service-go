package main

import "testing"

func TestHello(t *testing.T) {
	result := hello()

	var answer string = "CI/CD Workshop today with everyone"

	if result != answer {
		t.Errorf("Error the result, expected %v , got %v", answer, result)
	}

}