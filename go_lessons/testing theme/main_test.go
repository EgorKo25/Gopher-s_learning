package main

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		name    string
		value   float64
		trueRes float64
	}{
		{
			name:    "test small negative number",
			value:   -0.00000000006,
			trueRes: 0.00000000006,
		},
		{
			name:    "test small positive number",
			value:   0.00000000006,
			trueRes: 0.00000000006,
		},
		{
			name:    "simple test №1",
			value:   3,
			trueRes: 3,
		},
		{
			name:    "simple test №2",
			value:   -3,
			trueRes: 3,
		},
		{
			name:    "simple test №3",
			value:   2.00000001,
			trueRes: 2.00000001,
		},
		{
			name:    "simple test №1",
			value:   -2.00000001,
			trueRes: 2.00000001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := Abs(tt.value); res != tt.trueRes {
				t.Errorf("Abs: %f, want: %f", res, tt.trueRes)
			}
		})
	}
}
