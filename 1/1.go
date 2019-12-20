package 1

import (
  "io/ioutil"
  "strings"
  "fmt"
)

func loadInput() []string {
  dat, _ := ioutil.ReadFile("./1/input")
  sdata := string(dat)
  inputs := strings.SplitN(sdata, "\n", -1)
  return inputs
}

func calcFuel(masses []string) int {
  sum := 0
  for _, x := range masses {
    valu := math.Floor(x / 3) - 2
    sum += valu
  }
  return sum
}
