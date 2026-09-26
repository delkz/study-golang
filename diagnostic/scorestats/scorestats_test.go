package scorestats

import (
	"math"
	"reflect"
	"testing"
)

func TestAnalyzeValidScores(t *testing.T) {
	tests := []struct {
		name   string
		scores []int
		want   Summary
	}{
		{
			name:   "single score",
			scores: []int{80},
			want:   Summary{Count: 1, Total: 80, Average: 80, Min: 80, Max: 80},
		},
		{
			name:   "different scores",
			scores: []int{60, 75, 90, 77},
			want:   Summary{Count: 4, Total: 302, Average: 75.5, Min: 60, Max: 90},
		},
		{
			name:   "boundary values",
			scores: []int{0, 100},
			want:   Summary{Count: 2, Total: 100, Average: 50, Min: 0, Max: 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := append([]int(nil), tt.scores...)

			got, err := Analyze(tt.scores)
			if err != nil {
				t.Fatalf("Analyze(%v) returned unexpected error: %v", tt.scores, err)
			}

			if got.Count != tt.want.Count || got.Total != tt.want.Total ||
				got.Min != tt.want.Min || got.Max != tt.want.Max ||
				math.Abs(got.Average-tt.want.Average) > 1e-9 {
				t.Errorf("Analyze(%v) = %+v; want %+v", tt.scores, got, tt.want)
			}

			if !reflect.DeepEqual(tt.scores, original) {
				t.Errorf("Analyze modified its input: got %v; original %v", tt.scores, original)
			}
		})
	}
}

func TestAnalyzeRejectsInvalidInput(t *testing.T) {
	tests := []struct {
		name   string
		scores []int
	}{
		{name: "empty input", scores: nil},
		{name: "negative score", scores: []int{50, -1, 70}},
		{name: "score above maximum", scores: []int{50, 101, 70}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Analyze(tt.scores)
			if err == nil {
				t.Fatalf("Analyze(%v) returned nil error; want an error", tt.scores)
			}

			if got != (Summary{}) {
				t.Errorf("Analyze(%v) returned %+v on error; want zero Summary", tt.scores, got)
			}
		})
	}
}
