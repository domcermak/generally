package main

import (
	"fmt"

	"github.com/domcermak/generally/do"
)

type container struct {
	value int
}

func main() {
	containers := []container{
		{
			value: 1,
		},
		{
			value: 2,
		},
		{
			value: 3,
		},
		{
			value: 4,
		},
	}

	ints := do.Map(containers, func(c container) int {
		return c.value
	})
	lowerThanFour := do.Filter(ints, func(n int) bool {
		return n < 4
	})
	summed := do.Reduce(lowerThanFour, func(acc int, n int) int {
		return acc + n
	})
	fmt.Println(summed)
}

// Output:
// 6
