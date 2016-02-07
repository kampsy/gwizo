package main

import (
  "fmt"
  "github.com/kampsy/porter/gwizo"
  "bufio"
  "io/ioutil"
  "strings"
  "os"
)

func main() {
  writeOut()
  fmt.Println("Done")
}

func writeOut() {
  re, err := ioutil.ReadFile("words.txt")
  if err != nil {
    fmt.Println(err)
  }

  file := strings.NewReader(fmt.Sprintf("%s", re))
  scanner := bufio.NewScanner(file)
  out, err := os.Create("stemout.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer out.Close()

  for scanner.Scan() {
    txt := scanner.Text()
    str := stemix.Stem(txt)
    out.WriteString(fmt.Sprintf("%s\n", str))
    fmt.Println(txt, "--->", str)
  }
  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }
}
