package conditionals

import "testing"

func TestIsAdult(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want bool
	}{
		{name: "below boundary", age: 17, want: false},
		{name: "at boundary", age: 18, want: true},
		{name: "above boundary", age: 25, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAdult(tt.age); got != tt.want {
				t.Errorf("IsAdult(%d) = %t; want %t", tt.age, got, tt.want)
			}
		})
	}
}

func TestLarger(t *testing.T) {
	tests := []struct {
		name        string
		left, right int
		want        int
	}{
		{name: "left is larger", left: 9, right: 4, want: 9},
		{name: "right is larger", left: 3, right: 8, want: 8},
		{name: "values are equal", left: 5, right: 5, want: 5},
		{name: "negative values", left: -3, right: -8, want: -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Larger(tt.left, tt.right); got != tt.want {
				t.Errorf("Larger(%d, %d) = %d; want %d", tt.left, tt.right, got, tt.want)
			}
		})
	}
}

func TestCanAccess(t *testing.T) {
	tests := []struct {
		name          string
		age           int
		accountActive bool
		want          bool
	}{
		{name: "adult with active account", age: 18, accountActive: true, want: true},
		{name: "minor with active account", age: 17, accountActive: true, want: false},
		{name: "adult with inactive account", age: 30, accountActive: false, want: false},
		{name: "minor with inactive account", age: 12, accountActive: false, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanAccess(tt.age, tt.accountActive); got != tt.want {
				t.Errorf("CanAccess(%d, %t) = %t; want %t", tt.age, tt.accountActive, got, tt.want)
			}
		})
	}
}
