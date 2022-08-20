package main

import (
	"fmt"

	"github.com/domcermak/generally/extend"
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
	extendedContainers := extend.Slice(containers)
	extendedContainers.Map(func(c container) container {
		c.value += 5

		return c
	}).Filter(func(c container) bool {
		return c.value < 8
	}).Each(func(c container) {
		fmt.Println(c.value)
	})
}

// Output:
// 6
// 7
