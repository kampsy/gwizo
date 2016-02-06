package main

import (
  "fmt"
  "github.com/kampsy/porter/stemix"
)

func main() {
  val := stemix.Form("ponies")
  fmt.Println(val.Step_1a())
}
