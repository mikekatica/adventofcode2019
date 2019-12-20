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
	inputs := strings.Split(sdata, "\n")
	return inputs
}

func CalcFuel(masses []string) float64 {
	sum := float64(0)
	for _, x := range masses {
		num, _ := strconv.ParseFloat(x, 32)
		if num <= 0 {
			continue
		}
		fmt.Println("Mass: ", num)
		valu := doMassToFuelConversion(num)
		fmt.Println("Fuel: ", valu)
		sum += valu
	}
	fmt.Println("====================")
	return sum
}

func doMassToFuelConversion(mass float64) float64 {
	fuel := math.Floor(mass/3) - 2
	return fuel
}
