<h1 align="center">Gwizo</h1>
<br>
![home](https://github.com/kampsy/gwizo/blob/master/img/gwizo.png)

gwizo |pronounced as [guizo]| is the Next generation Native Go implementation of the
Porter Stemmer algorithm (An algorithm for suffix stripping M.F.Porter 1980 see:
(http://tartarus.org/martin/PorterStemmer/def.txt).
The uniqueness of gwizo is not that it is open source. Its that it's well designed.
It is designed to be extensible, so that developers can easily create
new experiences.(see examples below).

Gwizo is an awesome tool for projects involving:
1) Machine Learning algorithms, specifically Natural language processing (NLP).
2) Inverted indices for Information Retrieval Systems eg Search Engines.


The string that the Ingest() function takes is case insensitive

### Installation

To install, simply run in a terminal:

    go get github.com/kampsy/gwizo

# Usage

## DeepStem, ShallowStem, ShallowStemmed.

DeepStem: The ingested word goes through every step in the algorithm.
```go
package main

import (
  "fmt"
  "github.com/kampsy/gwizo"
)

func main() {
  octopus := gwizo.Ingest("abilities")
  str := octopus.DeepStem()
  fmt.Printf("Stem: %s\n", str)
}
```
```sh
$ go run main.go

Stem: able
```

## ShallowStem.

The word Goes through each step in accending order just like DeepStem. But The
difference is that it return when the original word is changed.
```go
package main

import (
  "fmt"
  "github.com/kampsy/gwizo"
)

func main() {
  octopus := gwizo.Ingest("abilities")
  str := octopus.ShallowStem()
  fmt.Printf("Stem: %s\n", str)
}
```

```sh
$ go run main.go

Stem: abiliti
```

## ShallowStemmed.

Works exactly like ShallowStem. The difference is that it returns
the Step that was used instead of the stem.
```go
package main

import (
  "fmt"
  "github.com/kampsy/gwizo"
)

func main() {
  octopus := gwizo.Ingest("abilities")

  str := octopus.ShallowStemmed()
  fmt.Printf("Step Used: %s\n", str)
}
```

```sh
$ go run main.go

Steps used: Step1a()
```

## Vowels, Consonants and Measure

gwizo returns a type called Octopus which has the following fields; The ingested
string Word, VowCon which is the vowel consonut pattern and the Measure value
[v]vc{m}[c]
```go
  package main

  import (
    "fmt"
    "github.com/kampsy/gwizo"
    "strings"
  )

func main() {
  octopus := gwizo.Ingest("abilities")

  // VowCon
  fmt.Printf("%s has Pattern %s \n", octopus.Word, octopus.VowCon)

  // Measure value [v]vc{m}[c]
  fmt.Printf("%s has Measure value %d \n", octopus.Word, octopus.Measure)

  // Number of Vowels
  v := strings.Count(octopus.VowCon, "v")
  fmt.Printf("%s Has %d Vowels \n", octopus.Word, v)

  // Number of Consonants
  c := strings.Count(octopus.VowCon, "c")
  fmt.Printf("%s Has %d Consonants\n", octopus.Word, c)

  // Or just print all of the values
  fmt.Println("values: ", octopus)// Type Octopus implements the Stringer interface.
}
```

```bash
$ go run main.go

abilities has Pattern vcvcvcvvc
abilities has Measure value 4
abilities Has 5 Vowels
abilities Has 4 Consonants
values:  {abilities vcvcvcvvc 4}
```

## Access Any Step Directly

gwizo is so extensible that it allows you to use its core components.
you can explicitly specify which Step to use on an ingested string.
```go
  package main

  import (
    "fmt"
    "github.com/kampsy/gwizo"
  )

func main() {
  octopus := gwizo.Ingest("troubled")

  // Stem only with Step1b
  fmt.Println(octopus.Step1b())

  octopus.Word = "vietnamization"
  // Stem only with Step2
  fmt.Println(octopus.Step2())

  octopus.Word = "electriciti"
  // Stem only with Step3
  fmt.Println(octopus.Step3())

  // You get the idea!
}
```
```sh
$ go run main.go

trouble
vietnamize
electric
```

## File Stem Performance.

gwizo stemmed a file input.txt containing 23531 words in 1.814791104s
on AMD C655 Laptop.
```go
  package main

  import (
    "fmt"
    "github.com/kampsy/gwizo"
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
      octopus := gwizo.Ingest(txt)
      str := octopus.DeepStem()
      out.WriteString(fmt.Sprintf("%s\n", str))
      fmt.Println(txt, "--->", str)
    }
    if err := scanner.Err(); err != nil {
      fmt.Println(err)
    }
  }
```
```sh
$ go run main.go

Done After: 1.814791104s
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
email: ***kampsycode@gmail.com***
