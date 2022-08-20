package do

import (
	"strconv"
	"testing"

	"github.com/domcermak/generally/do"
)

func TestProcessInBatches(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		src      []int
		expected []string
	}{
		{
			name:     "empty",
			src:      []int{},
			expected: []string{},
		},
		{
			name:     "short",
			src:      []int{1, 2, 3, 4, 5},
			expected: []string{"1", "2", "3", "4", "5"},
		},
		{
			name: "long",
			src: func() []int {
				slice := make([]int, 1000)
				for i := range slice {
					slice[i] = i
				}
				return slice
			}(),
			expected: func() []string {
				slice := make([]string, 1000)
				for i := range slice {
					slice[i] = strconv.Itoa(i)
				}
				return slice
			}(),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dst := make([]string, len(tc.src))
			do.ProcessInBatches(len(tc.src), func(from, to int) {
				for i, n := range tc.src[from:to] {
					dst[from+i] = strconv.Itoa(n)
				}
			})

			for i, s := range dst {
				if tc.expected[i] != s {
					t.Fatalf("%d) expected %s but got %s", i, tc.expected[i], s)
				}
			}
		})
	}
}
