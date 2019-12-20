package main

import (
	"fmt"

	problem1 "./problem1"
)

func main() {
	fmt.Println("===================")
	fmt.Println("Day 1:")
	inputs := problem1.LoadInput()
	mass := problem1.CalcFuel(inputs)
	fmt.Printf("%f", mass)
}
