package main

import (
  "io/ioutil"
  "strings"
  "fmt"
)

func main() {
  dat, _ := ioutil.ReadFile("./input")
  sdata := string(dat)
  inputs := strings.SplitN(sdata, "\n", -1)
  fmt.Println(inputs)
}
