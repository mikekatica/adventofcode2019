package problem1

import (
	"testing"
)

func TestFuelConversion(t *testing.T) {
	got := doMassToFuelConversion(1969)
	if got != 966 {
		t.Errorf("Got %f, needed to get 966 from fuel conversion", got)
	}
}
