package problem2

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func DoProblem2() {
	intcomp := parseArrayFromProblemInput("./problem2/input")
	fmt.Println(intcomp)
	processSlice(intcomp)
	fmt.Println(intcomp)
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
		fmt.Printf("String: %v, Int: %v\n", x, inputInt)
	}
	return inputs
}

func processSlice(r []int64) {
	exit := false
	for i := 0; i < len(r) && !exit; i += 4 {
		exit = evaluateSlice(i, r)
	}
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
