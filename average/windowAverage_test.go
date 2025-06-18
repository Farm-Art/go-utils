package average_test

import (
	"github.com/Farm-Art/go-utils/average"
	"math"
	"slices"
	"testing"
	"unsafe"
)

func TestWindowAverage(t *testing.T) {
	buf := make([]float64, 0, 3)

	wa := average.New(buf)

	if !math.IsNaN(wa.Min()) {
		t.Error("min: expected NaN, got", wa.Min())
	}
	if !math.IsNaN(wa.Max()) {
		t.Error("max: expected NaN, got", wa.Max())
	}

	wa.Add(5, 1, 2, 3, 4)
	if !slices.Contains(wa.Window(), 4) {
		t.Error("add: expected 4 to be in shared buf, got", buf)
	}
	if slices.Contains(wa.Window(), 1) {
		t.Error("add: window not properly overwritten, got", buf)
	}

	if unsafe.SliceData(wa.Window()) != unsafe.SliceData(buf) {
		t.Error("add: window reallocated")
	}

	if wa.Min() != 2 {
		t.Error("min: expected 3, got", wa.Min())
	}
	if wa.Max() != 4 {
		t.Error("max: expected 4, got", wa.Max())
	}

	if wa.Min() != 2 {
		t.Error("min: inconsistent value")
	}
	if wa.Max() != 4 {
		t.Error("max: inconsistent value")
	}

	if wa.Sum() != 9 {
		t.Error("sum: expected 9, got", wa.Sum())
	}
	if wa.Count() != 3 {
		t.Error("count: expected 3, got", wa.Count())
	}
	if wa.Avg() != 3 {
		t.Error("avg: expected 3, got", wa.Avg())
	}

	wa.Clear()

	if !math.IsNaN(wa.Max()) {
		t.Error("max: expected NaN, got", wa.Max())
	}
	if !math.IsNaN(wa.Min()) {
		t.Error("min: expected NaN, got", wa.Min())
	}
}
