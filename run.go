package main

import (
  "fmt"
  "github.com/kampsy/porter/stemix"
)

func main() {
  val := stemix.Form("monitoring")
  fmt.Println(val.Step_1b())
}
