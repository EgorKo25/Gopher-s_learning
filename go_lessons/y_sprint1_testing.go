package main

import "testing"

// Sum возвращает сумму элементов.
func Sum(values ...int) int {
	var sum int
	for _, v := range values {
		sum += v
	}
	return sum
}

// TestSum - протестировали Sum()
func TestSum(t *testing.T) {
	tests := []struct {
		name       string
		values     []int
		wantResult int
	}{
		{
			name:       "simple test",
			values:     []int{1, 2, 3},
			wantResult: 6,
		},
		{
			name:       "a lot of number test",
			values:     []int{1, 2, 3, 78, 0},
			wantResult: 84,
		},
		{
			name:       "negative values",
			values:     []int{-3, 2, 4},
			wantResult: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if sum := Sum(tt.values...); sum != tt.wantResult {
				t.Errorf("Add: %v, want: %v", sum, tt.wantResult)
			}
		})
	}
}

func main() {
	var t *testing.T
	TestSum(t)
}
