package main

import "testing"

func TestSayHello(t *testing.T) {
	actual := SayHello("Justin")
	expected := "Hello, Justin"

	if actual != expected {
		t.Errorf("actual = %q, expected = %q", actual, expected)
	}
}
