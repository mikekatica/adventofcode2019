package problem1

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func LoadInput() []string {
	dat, _ := ioutil.ReadFile("./problem1/input")
	sdata := string(dat)
	inputs := strings.SplitN(sdata, "\n", -1)
	return inputs
}

func CalcFuel(masses []string) float64 {
	sum := float64(0)
	for _, x := range masses {
		num, _ := strconv.ParseFloat(x, 32)
		fmt.Println(num)
		valu := math.Floor(num/3) - 2
		sum += valu
	}
	return sum
}
