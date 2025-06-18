package average_test

import (
	"github.com/Farm-Art/go-utils/average"
	"math/rand"
	"sync"
	"testing"
)

const (
	routines = 10
	inserts  = 10
)

func TestConcurrentAverage(t *testing.T) {
	a := average.Concurrent(average.New(nil))

	var wg sync.WaitGroup
	wg.Add(routines)
	for range routines {
		go func() {
			defer wg.Done()

			for range inserts {
				a.Add(rand.Float64())

				a.Window()

				a.Sum()
				a.Count()
				a.Avg()

				a.Min()
				a.Max()
			}

			a.Clear()
		}()
	}
	wg.Wait()
}
