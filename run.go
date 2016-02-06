package main

import (
  "fmt"
  "github.com/kampsy/porter/stemix"
)

func main() {
  val := stemix.Form("conflated")
  fmt.Println(val.Step_1b())
}
