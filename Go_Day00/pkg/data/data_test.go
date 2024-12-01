package data

import (
	"os"
	"testing"
)

func TestNumberData(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"1\n2\n3\n4\n", []int{1, 2, 3, 4}},
		{"10\n20\n30\n", []int{10, 20, 30}},
		{"1\n2\n3\n\n", []int{1, 2, 3}},
		{"145235\n2\n3\n\n", []int{2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := oneCase(t, tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("Length mismatch: got %v, expected %v", len(result), len(tt.expected))
				return
			}
			for i := 0; i < len(tt.expected); i++ {
				if result[i] != tt.expected[i] {
					t.Errorf("Number mismatch: got %v, expected %v", result[i], tt.expected[i])
				}
			}
		})
	}
}

func oneCase(t *testing.T, input string) []int {
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	_, err = w.Write([]byte(input))
	if err != nil {
		t.Fatalf("Failed to write to pipe: %v", err)
	}
	w.Close()

	os.Stdin = r

	return NumberData()
}
