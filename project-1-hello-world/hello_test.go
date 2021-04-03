package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	got := Hello()
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
	println(got)
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	got := Proverb()
	if got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
	println(got)
}

func TestMyHello(t *testing.T) {
	want := "Hello"
	got := MyHello()
	if got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
	println(got)
}
