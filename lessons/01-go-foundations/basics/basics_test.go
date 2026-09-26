package basics

import "testing"

func TestSum(t *testing.T) {
	got := Sum(7, 5)
	want := 12

	if got != want {
		t.Errorf("Sum(7, 5) = %d; want %d", got, want)
	}
}

func TestGreeting(t *testing.T) {
	got := Greeting("David")
	want := "Olá, David"

	if got != want {
		t.Errorf("Greeting(%q) = %q; want %q", "David", got, want)
	}
}

func TestRectangleArea(t *testing.T) {
	got := RectangleArea(4.5, 2)
	want := 9.0

	if got != want {
		t.Errorf("RectangleArea(4.5, 2) = %v; want %v", got, want)
	}
}
