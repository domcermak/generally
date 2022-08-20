package is

import (
	"testing"

	"github.com/domcermak/generally/is"
)

func TestAny(t *testing.T) {
	for _, tc := range []struct {
		name     string
		args     []bool
		expected bool
	}{
		{
			name:     "empty",
			args:     []bool{},
			expected: false,
		},
		{
			name:     "only true",
			args:     []bool{true, true, true},
			expected: true,
		},
		{
			name:     "one false",
			args:     []bool{true, false, true},
			expected: true,
		},
		{
			name:     "all false",
			args:     []bool{false, false, false},
			expected: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			actual := is.Any(tc.args, func(a bool) bool { return a })

			if actual != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, actual)
			}
		})
	}
}
