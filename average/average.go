package average

import "math"

type Average interface {
	// Add values to Average
	Add(...float64)

	// Sum of current values
	Sum() float64
	// Count of current values. Never grows above cap(Window), unless
	// cap(Window) == 0, in which case grows indefinitely.
	Count() int

	// Avg of values so far
	Avg() float64

	// Window of current values. Read-only, any write will break the Average.
	// Values are out of insertion order.
	//
	// If cap(Window) == 0, stores no values and treats Window as infinite.
	Window() []float64

	// Min value in window or NaN if Count == 0
	Min() float64
	// Max value in window or NaN if Count == 0
	Max() float64

	// Clear all values, retain Window buffer.
	Clear()
}

// New creates a new Average that now owns the passed window.
// cap(window) is used to determine its size.
//
// If len(window) != 0, contents will be used as initial data, including zeroes.
// If instead you need a blank buffer, make sure the len is 0.
//
// A window with capacity > 0 stores N=capacity values, Average.Count never
// grows above window size, and only stored values are used in Average.Avg,
// Average.Sum, Average.Min, Average.Max calculations.
// A window with capacity = 0 is considered infinitely large, but will not
// store any of the values. Thus, Average will account for every Average.Add since last
// Average.Clear.
func New(window []float64) Average {
	if cap(window) == 0 {
		return &AccumulatingAverage{}
	}

	return &WindowAverage{window: window}
}

// AccumulatingAverage is an Average with an infinite window, but no stored values.
type AccumulatingAverage struct {
	sum   float64
	count int

	max float64
	min float64
}

func (a *AccumulatingAverage) Add(values ...float64) {
	// Check for first Add
	if a.count == 0 && len(values) > 0 {
		a.min = values[0]
		a.max = values[0]
	}

	a.count += len(values)

	for _, v := range values {
		if v < a.min {
			a.min = v
		}
		if v > a.max {
			a.max = v
		}

		a.sum += v
	}
}

func (a *AccumulatingAverage) Window() []float64 {
	return nil
}

func (a *AccumulatingAverage) Avg() float64 {
	return a.sum / float64(a.count)
}

func (a *AccumulatingAverage) Sum() float64 {
	return a.sum
}

func (a *AccumulatingAverage) Count() int {
	return a.count
}

func (a *AccumulatingAverage) Clear() {
	a.sum = 0
	a.count = 0
}

func (a *AccumulatingAverage) Min() float64 {
	if a.count == 0 {
		return math.NaN()
	}
	return a.min
}

func (a *AccumulatingAverage) Max() float64 {
	if a.count == 0 {
		return math.NaN()
	}
	return a.max
}
