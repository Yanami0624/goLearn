package test_test

import (
	"Gocodes/tool"
	"testing"
)

func TestSum(t *testing.T) {
	a, b := 1, 2
	expected := a + b
	result := tool.Sum(a, b)
	if expected != result {
		t.Fatalf("expect %d, got %d", expected, result)
	}
}