package main

import (
  "fmt"
  "github.com/kampsy/porter/stemix"
)

func main() {
  val := stemix.Form("goodness")
  fmt.Println(val.Step_3())
}
