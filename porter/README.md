<h1 align="center">Porter</h1>
<br>
![home](https://github.com/kampsy/gwizo/blob/master/img/gwizo.png)

Package porter implement Porter, M. "An algorithm for suffix stripping."
Program 14.3 (1980): 130-137.
Martin Porter, the algorithm's inventor, maintains a web page about the
algorithm at http://www.tartarus.org/~martin/PorterStemmer/

## Installation

To install, simply run in a terminal:

    go get github.com/kampsy/gwizo/porter

# Usage

## Stem

DeepStem: The ingested word goes through every step in the algorithm.
```go
package main

import (
  "fmt"
  "github.com/kampsy/gwizo/porter"
)

func main() {
  stem := porter.Stem("abilities")
  fmt.Printf("Stem: %s\n", stem)
}
```
```shell
$ go run main.go

Stem: able
```

## Vowels, Consonants and Measure

porter returns a type Token VowCon which is the vowel consonut pattern 
and the Measure value [v]vc{m}[c]
```go
  package main

  import (
    "fmt"
    "github.com/kampsy/gwizo/porter"
    "strings"
  )

func main() {
  word := "abilities"
  token := porter.Parse(word)

  // VowCon
  fmt.Printf("%s has Pattern %s \n", word, token.VowCon)

  // Measure value [v]vc{m}[c]
  fmt.Printf("%s has Measure value %d \n", word, token.Measure)

  // Number of Vowels
  v := strings.Count(token.VowCon, "v")
  fmt.Printf("%s Has %d Vowels \n", word, v)

  // Number of Consonants
  c := strings.Count(token.VowCon, "c")
  fmt.Printf("%s Has %d Consonants\n", word, c)
}
```

```bash
$ go run main.go

abilities has Pattern vcvcvcvvc
abilities has Measure value 4
abilities Has 5 Vowels
abilities Has 4 Consonants
```

## File Stem Performance.

```go
  package main

  import (
    "fmt"
    "github.com/kampsy/gwizo/porter"
    "bufio"
    "io/ioutil"
    "strings"
    "os"
    "time"
  )

  func main() {
    curr := time.Now()
    writeOut()
    elaps := time.Since(curr)
    fmt.Println("============================")
    fmt.Println("Done After:", elaps)
    fmt.Println("============================")
  }

  func writeOut() {
    re, err := ioutil.ReadFile("input.txt")
    if err != nil {
      fmt.Println(err)
    }

    file := strings.NewReader(fmt.Sprintf("%s", re))
    scanner := bufio.NewScanner(file)
    out, err := os.Create("stem.txt")
    if err != nil {
      fmt.Println(err)
    }
    defer out.Close()
    for scanner.Scan() {
      txt := scanner.Text()
      stem := porter.Stem(txt)
      out.WriteString(fmt.Sprintf("%s\n", stem))
      fmt.Println(txt, "--->", str)
    }
    if err := scanner.Err(); err != nil {
      fmt.Println(err)
    }
  }
```
```shell
$ go run main.go

```

## License
BSD style - see license file.

## Developer
kampamba chanda (a.k.a kampsy).
<br>
Twitter: ***@kampsy***
<br>
Google+: ***google.com/+kampambachanda***
<br>
email: ***kampambachanda@gmail.com***