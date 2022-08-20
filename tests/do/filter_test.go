package do

import (
	"testing"

	"domcermak/stl/do"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name       string
		filterFunc func(int) bool
		numbers    []int
		expected   []int
	}{
		{
			name: "empty",
			filterFunc: func(i int) bool {
				return true
			},
			numbers:  []int{},
			expected: []int{},
		},
		{
			name: "short - filter out all",
			filterFunc: func(i int) bool {
				return false
			},
			numbers:  []int{1, 2, 3, 4, 5, 6, 6, 8},
			expected: []int{},
		},
		{
			name: "short - filter out some",
			filterFunc: func(i int) bool {
				return i < 5
			},
			numbers:  []int{1, 2, 3, 4, 5, 6, 6, 8},
			expected: []int{1, 2, 3, 4},
		},
		{
			name: "short - filter out none",
			filterFunc: func(i int) bool {
				return true
			},
			numbers:  []int{1, 2, 3, 4, 5, 6, 6, 8},
			expected: []int{1, 2, 3, 4, 5, 6, 6, 8},
		},
		{
			name: "long - filter out some",
			filterFunc: func(i int) bool {
				return i > 499
			},
			numbers: func() []int {
				ary := make([]int, 1000)
				for i := range ary {
					ary[i] = i
				}

				return ary
			}(),
			expected: func() []int {
				ary := make([]int, 500)
				for i := range ary {
					ary[i] = i + 500
				}

				return ary
			}(),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			actual := do.Filter(tc.expected, tc.filterFunc)

			if !do.Match(tc.expected, actual) {
				t.Fatalf("expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}
