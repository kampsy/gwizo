package main

import (
  "fmt"
  "github.com/kampsy/porter/stemix"
)

func main() {
  val := stemix.Form("relational")
  fmt.Println(val.Step_2())
}
