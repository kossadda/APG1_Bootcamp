package anscombe

import (
	"fmt"
	"testing"
)

func TestMean(t *testing.T) {
	tests := []struct {
		input    []int
		expected float64
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0},
		{[]int{10, 20, 30, 40, 50}, 30.0},
		{[]int{-1, -2, -3, -4, -5}, -3.0},
		{[]int{}, 0.0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			result := Mean(tt.input)
			if result != tt.expected {
				t.Errorf("Mean(%v) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMode(t *testing.T) {
	tests := []struct {
		input    []int
		expected float64
	}{
		{[]int{1, 2, 2, 3, 4}, 2.0},
		{[]int{1, 1, 2, 2, 3, 3, 4}, 1.0},
		{[]int{1, 2, 3, 4, 5}, 1.0},
		{[]int{}, 0.0},
	}

	for _, tt := range tests {
		result := Mode(tt.input)
		if result != tt.expected {
			t.Errorf("Mean(%v) = %f; want %f", tt.input, result, tt.expected)
		}
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		input    []int
		expected float64
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0},
		{[]int{1, 2, 3, 4, 5, 6}, 3.5},
		{[]int{10, 20, 30, 40, 50}, 30.0},
		{[]int{10, 20, 30, 40, 50, 60}, 35.0},
		{[]int{}, 0.0},
	}

	for _, tt := range tests {
		result := Median(tt.input)
		if result != tt.expected {
			t.Errorf("Median(%v) = %f; want %f", tt.input, result, tt.expected)
		}
	}
}

func TestDeviation(t *testing.T) {
	tests := []struct {
		input    []int
		expected float64
	}{
		{[]int{1, 2, 3, 4, 5}, 1.4142135623730951},
		{[]int{10, 20, 30, 40, 50}, 14.142135623730951},
		{[]int{-1, -2, -3, -4, -5}, 1.4142135623730951},
		{[]int{}, 0.0},
	}

	for _, tt := range tests {
		result := Deviation(tt.input)
		if result != tt.expected {
			t.Errorf("Deviation(%v) = %f; want %f", tt.input, result, tt.expected)
		}
	}
}
