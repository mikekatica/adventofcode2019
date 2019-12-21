package problem2

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func DoProblem2() {
	intcomp := parseArrayFromProblemInput("./problem2/input")
	res := processSlice(intcomp)
	fmt.Println(intcomp[0])
	fmt.Println(res[0])
}

func DoProblem2Part2() {
	baseMemory := parseArrayFromProblemInput("./problem2/input")
	result := processSlice(baseMemory)
	var o1 int
	var o2 int
	for i := 0; i < len(baseMemory) && result[0] != int64(19690720); i++ {
		for j := 0; j < i && j < len(baseMemory) && result[0] != int64(19690720); j++ {
			result = processSliceRepl(baseMemory, i, j)
			o2 = j
		}
		o1 = i
	}
	fmt.Printf("Op: %v%v", o1, o2)
}

func parseArrayFromProblemInput(path string) []int64 {
	dat, _ := ioutil.ReadFile(path)
	sdata := string(dat)
	trimData := strings.TrimSuffix(sdata, "\n")
	splitTrimData := strings.Split(trimData, ",")
	var inputs []int64
	for _, x := range splitTrimData {
		inputInt, _ := strconv.ParseInt(x, 10, 32)
		inputs = append(inputs, inputInt)
	}
	return inputs
}

func processSlice(r []int64) []int64 {
	exit := false
	copyOfr := make([]int64, len(r))
	copy(copyOfr, r)
	for i := 0; i < len(r) && !exit; i += 4 {
		exit = evaluateSlice(i, copyOfr)
	}
	return copyOfr
}

func processSliceRepl(r []int64, spot1 int, spot2 int) []int64 {
	exit := false
	copyOfr := make([]int64, len(r))
	copy(copyOfr, r)
	copyOfr[1] = int64(spot1)
	copyOfr[2] = int64(spot2)
	for i := 0; i < len(r) && !exit; i += 4 {
		exit = evaluateSlice(i, copyOfr)
	}
	return copyOfr
}

func evaluateSlice(startingIndex int, slice []int64) bool {
	switch slice[startingIndex] {
	case 99:
		return true
	case 1:
		arg1Index := slice[startingIndex+1]
		arg2Index := slice[startingIndex+2]
		targetIndex := slice[startingIndex+3]
		slice[targetIndex] = slice[arg1Index] + slice[arg2Index]
		return false
	case 2:
		arg1Index := slice[startingIndex+1]
		arg2Index := slice[startingIndex+2]
		targetIndex := slice[startingIndex+3]
		slice[targetIndex] = slice[arg1Index] * slice[arg2Index]
		return false
	default:
		return false
	}
}
