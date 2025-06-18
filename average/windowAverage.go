package average

import "math"

// WindowAverage is an implementation of Average with a finite window of values,
// where new values > cap(Average.Window) overwrite previously known values.
//
// Zero value is unusable, WindowAverage.window must be set.
type WindowAverage struct {
	avg    AccumulatingAverage
	window []float64
	index  int
	// validMinMax determines if min/max values have to be recalculated.
	validMinMax bool
}

func (w *WindowAverage) Add(values ...float64) {
	if len(values) > 0 {
		w.validMinMax = false
	}

	w.avg.Add(values...)
	for _, value := range values {
		// Value is added, no need to alter anything
		if len(w.window) < cap(w.window) {
			w.window = append(w.window, value)
			continue
		}

		// Value is replaced, subtract and overwrite
		w.avg.sum -= w.window[w.index]
		w.window[w.index] = value

		w.avg.count--

		w.index = (w.index + 1) % len(w.window)
	}
}

func (w *WindowAverage) Sum() float64 {
	return w.avg.Sum()
}

func (w *WindowAverage) Count() int {
	return w.avg.Count()
}

func (w *WindowAverage) Avg() float64 {
	return w.avg.Avg()
}

func (w *WindowAverage) Window() []float64 {
	return w.window
}

func (w *WindowAverage) Min() float64 {
	if !w.validMinMax {
		min_, _ := w.recalcMinMax()
		return min_
	}
	return w.avg.Min()
}

func (w *WindowAverage) Max() float64 {
	if !w.validMinMax {
		_, max_ := w.recalcMinMax()
		return max_
	}
	return w.avg.Max()
}

func (w *WindowAverage) Clear() {
	w.avg.Clear()
	w.window = w.window[:0]
}

func (w *WindowAverage) recalcMinMax() (min, max float64) {
	// Should be threadsafe if wrapped with Concurrent RLock?
	// By the time validMinMax is set, both values are guaranteed correct;
	// No invalidation possible due to RLock surrounding Min/Max call;
	// Worst case should be duplicate calculation of same values, I think?
	if len(w.window) == 0 {
		min, max = math.NaN(), math.NaN()
		w.avg.min, w.avg.max = min, max
		return
	}

	max = w.window[0]
	min = w.window[0]
	for _, value := range w.window {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	w.avg.min = min
	w.avg.max = max
	w.validMinMax = true
	return
}
