package average_test

import (
	"github.com/Farm-Art/go-utils/average"
	"math"
	"testing"
)

func TestNew(t *testing.T) {
	buf := make([]float64, 0, 1)
	a := average.New(nil)

	_, ok := a.(*average.AccumulatingAverage)
	if !ok {
		t.Error("expected *average.AccumulatingAverage")
	}

	wa := average.New(buf)

	_, ok = wa.(*average.WindowAverage)
	if !ok {
		t.Error("expected *average.WindowAverage")
	}
}

func TestAccumulatingAverage(t *testing.T) {
	a := average.New(nil)

	if !math.IsNaN(a.Min()) {
		t.Error("min: expected NaN, got", a.Min())
	}
	if !math.IsNaN(a.Max()) {
		t.Error("max: expected NaN, got", a.Max())
	}

	a.Add(10, 15, 5)
	if a.Count() != 3 {
		t.Error("count: expected", 3, "got", a.Count())
	}

	if a.Sum() != 30 {
		t.Error("sum: expected", 30, "got", a.Sum())
	}

	if a.Avg() != 10 {
		t.Error("avg: expected", 10, "got", a.Avg())
	}

	if a.Min() != 5 {
		t.Error("min: expected", 5, "got", a.Min())
	}

	if a.Max() != 15 {
		t.Error("max: expected", 15, "got", a.Max())
	}

	if a.Window() != nil {
		t.Error("window: expected", nil, "got", a.Window())
	}

	a.Clear()
	if a.Count() != 0 {
		t.Error("clear: expected", 0, "got", a.Count())
	}

	if !math.IsNaN(a.Avg()) {
		t.Error("avg: expected", a.Avg(), "got", a.Avg())
	}

	if !math.IsNaN(a.Min()) {
		t.Error("post-clear min: expected NaN, got", a.Min())
	}
	if !math.IsNaN(a.Max()) {
		t.Error("post-clear max: expected NaN, got", a.Max())
	}
}
