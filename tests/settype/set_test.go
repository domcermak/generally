package settype

import (
	"domcermak/stl/containers/settype"
	"testing"
)

func TestContains_String(t *testing.T) {
	for _, tc := range []struct {
		name     string
		slice    []string
		item     string
		expected bool
	}{
		{
			name:     "empty",
			slice:    []string{},
			item:     "",
			expected: false,
		},
		{
			name:     "non-empty but does not contain",
			slice:    []string{"hello"},
			item:     "",
			expected: false,
		},
		{
			name:     "non-empty and contains",
			slice:    []string{"hello", "world", "foo", "bar"},
			item:     "foo",
			expected: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			actual := settype.FromSlice(tc.slice).Contains(tc.item)

			if tc.expected != actual {
				t.Fatalf("expected %v but got %v", tc.expected, actual)
			}
		})
	}
}
