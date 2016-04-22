gwizo v 1.0
===========
<code>The stemmer with a magic touch </code>
<a href="https://youtu.be/At0orCwqHwM">Play Screencast</a>
<br>
<img src="https://github.com/kampsy/gwizo/blob/master/img/gwizo.png" height="200px" width="200px">
<br>
<code>image made by Renee French under Creative Commons 3.0 Attributions. Modified and improved by Olga Shalakhina osshalakhina@gmail.com</code>
<hr>

gwizo |pronounced as [guizo]| is the Next generation Native Go implementation of the
Porter Stemmer algorithm (An algorithm for suffix stripping M.F.Porter 1980 see:
(http://tartarus.org/martin/PorterStemmer/def.txt).
The uniqueness of gwizo is not that it is open source. Its that it's well designed.
It is designed to be extensible, so that developers can easily create
new experiences.(see examples below).

Gwizo is an awesome tool for projects involving:
1) Machine Learning algorithms, specifically Natural language processing (NLP).
2) Inverted indices for Information Retrieval Systems eg Search Engines.


Note: I made a few modification to gwizo for it to pass all tests. The original algorithm
at http://tartarus.org/martin/PorterStemmer/def.txt) has a few issues(opinion!).

The string that the Ingest() function takes is case insensitive

Installation
------------
<pre>
  go get github.com/kampsy/gwizo
</pre>

[[[[[ Examples ]]]]]

DeepStem, ShallowStem, ShallowStemmed
====================================================
DeepStem: The ingested word goes through every step in the algorithm.
<pre>
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
  Results
  ---------------------
  Steps used: Step1a() then Step2()
  Stem: able
</pre>

ShallowStem: The word Goes through each step, from top to bottom like in DeepStem. The
difference is that it bells out the moment a step Stems the ingested word.
<pre>
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
  Results
  ---------------------
  Steps used: Step_1a()
  Stem: abiliti
</pre>

ShallowStemmed: Works exactly like ShallowStem. The difference is that it returns
the Step that was used instead of the stem.
<pre>
  package main

  import (
    "fmt"
    "github.com/kampsy/gwizo"
  )

  func main() {
    octopus := gwizo.Ingest("abilities")

    str := octopus.ShallowStemmed()
    fmt.Printf("Stem: %s\n", str)
  }
  Results
  ---------------------
  Steps used: Step1a()
</pre>

Vowels, Consonants and Measure
====================================================
gwizo returns a type called Octopus which has the following fields; The ingested
string Word, VowCon which is the vowel consonut pattern and the Measure value
[v]vc{m}[c]
<pre>
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
  Results
  ---------------------
  abilities has Pattern vcvcvcvvc
  abilities has Measure value 4
  abilities Has 5 Vowels
  abilities Has 4 Consonants
  values:  {abilities vcvcvcvvc 4}
</pre>

Access Any Step Directly
====================================================
gwizo is so extensible that it allows you to use its core components.
you can explicitly specify which Step to use on an ingested string.
<pre>
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
  Results
  ---------------------
  trouble
  vietnamize
  electric
</pre>

File Stem Performance.
====================================================
gwizo stemmed a file input.txt containing 23531 words in 1.814791104s
<<<<<<< HEAD
on AMD C655 Laptop.
=======
on my computer
>>>>>>> 802c8414914e6b1a23c52d15bcd031cffab39589
<pre>
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
  Results
  ---------------------
  Done After: 1.814791104s
</pre>

License
==========
BSD style - see license file.

Developer
===============
kampamba chanda (a.k.a kampsy).
twitter @kampsy
google+ google.com/+kampambachanda
email: kampsycode@gmail.com
