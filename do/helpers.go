package do

import (
	"math"
	"runtime"
	"sync"
)

// ProcessInBatches splits given data length into equally large intervals according to the number of logical CPUs
// and executes the given function with the interval in parallel.
func ProcessInBatches(dataLength int, f func(from, to int)) {
	routines := runtime.NumCPU() * 2
	step := math.Max(1, float64(dataLength)/float64(routines))
	nextIndex := func(i int) int {
		return int(math.Min(math.Round(float64(i)+step), float64(dataLength)))
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < dataLength; i = nextIndex(i) {
		wg.Add(1)
		go func(from, to int) {
			f(from, to)
			wg.Done()
		}(i, nextIndex(i))
	}
	wg.Wait()
}

func async[T any, E any](src []T, dst []E, f func(T) E) {
	if len(src) == 0 {
		return
	}
	if len(src) < 20 {
		syncIterate(src, dst, 0, len(src), f)
		return
	}

	ProcessInBatches(len(src), func(from, to int) {
		syncIterate(src, dst, from, to, f)
	})
}

func asyncWithIndex[T any, E any](src []T, dst []E, f func(T, int) E) {
	if len(src) == 0 {
		return
	}
	if len(src) < 20 {
		syncIterateWithIndex(src, dst, 0, len(src), f)
		return
	}

	ProcessInBatches(len(src), func(from, to int) {
		syncIterateWithIndex(src, dst, from, to, f)
	})
}

func syncIterate[T any, E any](src []T, dst []E, from, to int, f func(T) E) {
	for i := from; i < to; i++ {
		dst[i] = f(src[i])
	}
}

func syncIterateWithIndex[T any, E any](src []T, dst []E, from, to int, f func(T, int) E) {
	for i := from; i < to; i++ {
		dst[i] = f(src[i], i)
	}
}
