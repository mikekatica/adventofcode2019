package main

import (
	"fmt"

	problem1 "./problem1"
)

func main() {
	inputs := problem1.LoadInput()
	mass := problem1.CalcFuel(inputs)
	fmt.Println(mass)
}
