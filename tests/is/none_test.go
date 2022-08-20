package is

import (
	"testing"

	"github.com/domcermak/generally/is"
)

func TestNone(t *testing.T) {
	for _, tc := range []struct {
		name     string
		args     []bool
		expected bool
	}{
		{
			name:     "empty",
			args:     []bool{},
			expected: true,
		},
		{
			name:     "only true",
			args:     []bool{true, true, true},
			expected: false,
		},
		{
			name:     "one false",
			args:     []bool{true, false, true},
			expected: false,
		},
		{
			name:     "all false",
			args:     []bool{false, false, false},
			expected: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			actual := is.None(tc.args, func(a bool) bool { return a })

			if actual != tc.expected {
				t.Fatalf("expected %v but got %v", tc.expected, actual)
			}
		})
	}
}
