package main

import (
	"fmt"

	problem1 "./problem1"
	problem2 "./problem2"
	problem3 "./problem3"
)

func main() {
	fmt.Println("===================")
	fmt.Println("Day 1:")
	inputs := problem1.LoadInput()
	mass := problem1.CalcFuel(inputs)
	fmt.Printf("%f", mass)
	fmt.Println("\n===================")
	fmt.Println("Day 2:")
	problem2.DoProblem2()
	problem2.DoProblem2Part2()
	fmt.Println("\n===================")
	fmt.Println("Day 3:")
	dist := problem3.DoProblem3()
	fmt.Println(dist)
}
