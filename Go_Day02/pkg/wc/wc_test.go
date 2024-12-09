package wc

import (
	"testing"
)

func TestAll(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
		wantErr  bool
	}{
		{
			name:     "count lines",
			input:    []string{"-l", "../../README.md"},
			expected: []string{"150\t../../README.md"},
			wantErr:  false,
		},
		{
			name:     "count words",
			input:    []string{"-w", "../../README.md"},
			expected: []string{"1190\t../../README.md"},
			wantErr:  false,
		},
		{
			name:     "count characters",
			input:    []string{"-m", "../../README.md"},
			expected: []string{"7763\t../../README.md"},
			wantErr:  false,
		},
		{
			name:     "multiple files",
			input:    []string{"-w", "../../.gitignore", "../../go.mod"},
			expected: []string{"2\t../../.gitignore", "4\t../../go.mod"},
			wantErr:  false,
		},
		{
			name:    "no flags",
			input:   []string{""},
			wantErr: true,
		},
		{
			name:    "invalid file path",
			input:   []string{"-w", "nonexistent.txt"},
			wantErr: true,
		},
		{
			name:    "conflicting flags",
			input:   []string{"-l", "-m", "../../README.md"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := make(map[string]struct{})
			for _, val := range tt.expected {
				m[val] = struct{}{}
			}

			w, err := New(&tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("unexpected error state: got %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			ch := Output(tt.input, w)

			for str := range ch {
				if _, ok := m[str]; !ok {
					t.Errorf("unexpected output: got %s", str)
					return
				}
			}
		})
	}
}
