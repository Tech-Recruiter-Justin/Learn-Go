package hello_test

import (
	"learnGoWithTests/hello"
	"testing"
)

func TestSayHello(t *testing.T) {
	t.Run("say hello to a person", func(t *testing.T) {
		actual := hello.SayHello("Justin", "")
		expected := "Hello, Justin"
		assertEquals(t, actual, expected)
	})

	t.Run("say hello to the world when the person is absent", func(t *testing.T) {
		actual := hello.SayHello("", "")
		expected := "Hello, world"
		assertEquals(t, actual, expected)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		actual := hello.SayHello("Carlos", "Spanish")
		expected := "Hola, Carlos"
		assertEquals(t, actual, expected)
	})

	t.Run("say hello in French", func(t *testing.T) {
		actual := hello.SayHello("Alex", "French")
		expected := "Bonjour, Alex"
		assertEquals(t, actual, expected)
	})
}

func assertEquals(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("actual = %q, expected = %q", actual, expected)
	}
}
