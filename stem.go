package main

import (
  "fmt"
  "github.com/kampsy/porter/gwizo"
  "strings"
)

func main() {
  analyse := gwizo.Form("abatements")
  val := strings.Count(analyse.VowCon, "c")
  fmt.Printf("%s Has %d consonants\n", analyse.Word, val)
}
