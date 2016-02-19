![gwizo](https://github.com/kampsy/gwizo/img/gwizo.png)
## gwizo

gwizo |pronounced as [guizo]| is the Next generation Native Go implementation of the
Porter Stemmer algorithm (An algorithm for suffix stripping M.F.Porter 1980 see:
(http://tartarus.org/martin/PorterStemmer/def.txt).
The uniqueness of gwizo is not that it is open source. Its that it's well designed and   
extensible. It is designed to be extensible, so that developers can easily create
new experiences.(see examples below).

Gwizo is built for people doing:
1) Machine Learning algorithms, specifically Natural language processing (NLP).
2) An Inverted index for an Information Retrieval System or Search Engine.


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
  Steps used: Step_1a() then Step_2()
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
  Steps used: Step_1a()
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

    // Stem only with Step_1b
    fmt.Println(octopus.Step_1b())

    octopus.Word = "vietnamization"
    // Stem only with Step_2
    fmt.Println(octopus.Step_2())

    octopus.Word = "electriciti"
    // Stem only with Step_3
    fmt.Println(octopus.Step_3())

    // You get the idea!
  }
  Results
  ---------------------
  trouble
  vietnamize
  electric
</pre>

License
==========
BSD style - see license file.

Developer
===============
kampamba chanda (a.k.a kampsy).
email: kampsycode@gmail.com
