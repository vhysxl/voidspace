package common

import "testing"

func TestNullIfEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "empty string returns nil",
			input:    "",
			expected: nil,
		},
		{
			name:     "whitespace string returns nil",
			input:    "   ",
			expected: nil,
		},
		{
			name:     "non-empty string returns itself",
			input:    "test",
			expected: "test",
		},
		{
			name:     "string with spaces returns itself",
			input:    "  test  ",
			expected: "  test  ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NullIfEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("NullIfEmpty(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
