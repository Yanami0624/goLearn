package test_test

import (
	hello "Gocodes"
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	data := "Alice"
	expect := fmt.Sprintf("Hello, %s!", data)
	result := hello.Hello(data)
	if result != expect {
		t.Fatalf("expected %s, but got %s", expect, result)
	}
}

func ExampleHello() {
	hello.Hello("Alice")
}