package tcptest

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, World!"
	got := "Hello, World!"

	if got != expected {
		t.Errorf("Hello() = %q; want %q", got, expected)
	}
}
