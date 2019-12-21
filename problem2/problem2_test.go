package problem2

import (
	"fmt"
	"strings"
	"testing"
)

func TestOp(t *testing.T) {
	tset := [][]int64{{1, 0, 0, 0, 99}, {2, 3, 0, 3, 99}, {2, 4, 4, 5, 99, 0}, {1, 1, 1, 4, 99, 5, 6, 0, 99}}
	tsetdes := [][]int64{{2, 0, 0, 0, 99}, {2, 3, 0, 6, 99}, {2, 4, 4, 5, 99, 9801}, {30, 1, 1, 4, 2, 5, 6, 0, 99}}
	for i := range tset {
		OpAndDesiredResult(tset[i], tsetdes[i], t)
	}
}

func TestParse(t *testing.T) {
	parsed := parseArrayFromProblemInput("./input_test")
	comp := []int64{1, 2, 3, 4, 5}
	fmt.Println(parsed)
	fmt.Println(comp)
	if len(parsed) != len(comp) {
		t.Errorf("Lengths are not the same, Test: %v and Desried: %v", len(parsed), len(comp))
	} else {
		for i := range comp {
			if parsed[i] != comp[i] {
				t.Errorf("Test: %v does not equal Want: %v", parsed, comp)
			}
		}
	}
}

func OpAndDesiredResult(op []int64, des []int64, t *testing.T) {
	postOp := processSlice(op)
	if len(postOp) != len(des) {
		t.Error("Lengths of the two did not match")
	} else {
		for i := range des {
			if postOp[i] != des[i] {
				var b strings.Builder
				b.WriteString("Wanted ")
				b.WriteString(fmt.Sprintf("%v", des))
				b.WriteString(", got ")
				b.WriteString(fmt.Sprintf("%v", op))
				t.Error(b.String())
			}
		}
	}
}
