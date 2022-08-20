package do

import (
	"domcermak/stl/do"
	"testing"
)

func TestMatch(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name        string
		a, b        []int
		shouldMatch bool
	}{
		{
			name:        "both empty",
			a:           []int{},
			b:           []int{},
			shouldMatch: true,
		},
		{
			name:        "a longer",
			a:           []int{1},
			b:           []int{},
			shouldMatch: false,
		},
		{
			name:        "b longer",
			a:           []int{},
			b:           []int{1},
			shouldMatch: false,
		},
		{
			name:        "both short - non matching",
			a:           []int{1, 2, 3, 4},
			b:           []int{1, 2, 3, 5},
			shouldMatch: false,
		},
		{
			name:        "both short - matching",
			a:           []int{1, 2, 3, 4},
			b:           []int{1, 2, 3, 4},
			shouldMatch: true,
		},
		{
			name: "both long - non matching",
			a: func() []int {
				ary := make([]int, 1000)
				for i := range ary {
					ary[i] = i
				}

				return ary
			}(),
			b: func() []int {
				ary := make([]int, 1000)
				for i := range ary {
					if i == 800 {
						ary[i] = i + 1
					} else {
						ary[i] = i
					}
				}

				return ary
			}(),
			shouldMatch: false,
		},
		{
			name: "both long - matching",
			a: func() []int {
				ary := make([]int, 1000)
				for i := range ary {
					ary[i] = -i
				}

				return ary
			}(),
			b: func() []int {
				ary := make([]int, 1000)
				for i := range ary {
					ary[i] = -i
				}

				return ary
			}(),
			shouldMatch: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			actual := do.Match(tc.a, tc.b)
			if actual != tc.shouldMatch {
				if tc.shouldMatch {
					t.Fatalf("expected %v to match %v, but got %v", tc.a, tc.b, false)
				} else {
					t.Fatalf("expected %v not to match %v, but got %v", tc.a, tc.b, false)
				}
			}
		})
	}
}
