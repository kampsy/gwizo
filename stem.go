package main

import (
  "fmt"
  "github.com/kampsy/porter/stemix"
)

func main() {
  str := stemix.Stem("revival")
  fmt.Println(str)
}
