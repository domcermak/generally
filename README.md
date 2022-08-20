# Generally

Generally is a library leveraging generics introduced in go 1.18 to make life easier.
It allows to you to:
- Extend basic types to be able to call methods on them.
- Improve code readability by using higher order functions with generic parameters.
- Set implementation with generics.

and probably more in the future depending on what I (or others) need.

## How to use

```go
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
```

See more in [examples](/examples)