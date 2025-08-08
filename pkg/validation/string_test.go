package validation_test

import (
	"github.com/rfanazhari/hris/pkg/validation"
	"testing"
)

func TestCharacterLong(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		length   int
		expected string
	}{
		{
			name:     "single_letter_key",
			key:      "A",
			length:   5,
			expected: "A must be at least 5 characters long",
		},
		{
			name:     "empty_key",
			key:      "",
			length:   10,
			expected: " must be at least 10 characters long",
		},
		{
			name:     "longer_key",
			key:      "Username",
			length:   15,
			expected: "Username must be at least 15 characters long",
		},
		{
			name:     "key_with_spaces",
			key:      "First Name",
			length:   20,
			expected: "First Name must be at least 20 characters long",
		},
		{
			name:     "zero_length",
			key:      "Password",
			length:   0,
			expected: "Password must be at least 0 characters long",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.CharacterLong(tt.key, tt.length)
			if err == nil {
				t.Fatalf("expected an error, got nil")
			}
			if err.Error() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, err.Error())
			}
		})
	}
}
