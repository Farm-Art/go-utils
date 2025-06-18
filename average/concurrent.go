package average

import "sync"

// Concurrent wraps the given Average with a sync.RWMutex, making it threadsafe-ish.
// Specifically, it guarantees a valid Average.Average at any point.
func Concurrent(avg Average) Average {
	return &ConcurrentAverage{avg: avg}
}

// ConcurrentAverage is a wrapper over Average that uses a RWMutex
// to make it threadsafe-ish.
type ConcurrentAverage struct {
	mu  sync.RWMutex
	avg Average
}

func (c *ConcurrentAverage) Add(f ...float64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.avg.Add(f...)
}

func (c *ConcurrentAverage) Sum() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.avg.Sum()
}

func (c *ConcurrentAverage) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.avg.Count()
}

func (c *ConcurrentAverage) Window() []float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.avg.Window()
}

func (c *ConcurrentAverage) Avg() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.avg.Avg()
}

func (c *ConcurrentAverage) Min() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.avg.Min()
}

func (c *ConcurrentAverage) Max() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.avg.Max()
}

func (c *ConcurrentAverage) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.avg.Clear()
}
