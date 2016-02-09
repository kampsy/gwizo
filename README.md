![gwizo](https://github.com/kampsy/gwizo/img/gwizo.png)
## gwizo

gwizo |pronounced as [guizo]| is the Next generation Native Go implementation of the
Porter Stemmer algorithm (An algorithm for suffix stripping M.F.Porter 1980 see: http://tartarus.org/martin/PorterStemmer/def.txt).
Gwizo is Better than other Go implementations of the algorithm because it Provides
powerful features Like the value of Measure([v]vc{m}[c]) for a word, the number of Vowel
and Consonants that a word has or had?, providing information about which Step was used during
the stemming process and developer access to all steps(Step_1a, Step_2 etc) of the algorithm.
Gwizo was built from the ground up with modularity in mind. Developers have direct access to all Core components of the algorithm, which results in doing lots of cool, powerful and amazing things (see examples below).

Gwizo is built for people doing Machine Learning algorithms, specifically
Natural language processing (NLP).

Note: I made a few modification to gwizo for it to pass all tests. The original algorithm
at http://tartarus.org/martin/PorterStemmer/def.txt) has a few issues.

Installation
------------
<pre>
  go get github.com/kampsy/gwizo
</pre>

[[[[[ Examples ]]]]]

DeepStem, ShallowStem, ShallowStemmed
------------------------------------------------------
This returns the stem of your string and the Step that was used.
<pre>
  package main

  import (
    "fmt"
    "github.com/kampsy/gwizo"
  )

  func main() {
    octopus := gwizo.Ingest("Consonants")

    deepstr := octopus.DeepStem() // Returns the stem from using every step
    fmt.Printf("Stem: %s\n", deepstr)

    shallowstr := octopus.ShallowStem() // Returns the stem from using one step
    shallowstep := octopus.ShallowStemmed() // Returns the step used
    fmt.Printf("ShallowStem: %s\nShallowStep  Used: %s\n", shallowstr, shallowstep)
  }
  Results
  ---------------------
  ShallowStem: consonant
  Stemmed With: Step_1a
</pre>

Vowels, Consonants and Measure
------------------------------
This returns the Vowels, Consonants and Measure of your word
<pre>
  package main

  import (
    "fmt"
    "github.com/kampsy/gwizo"
    "strings"
  )

  func main() {
    octopus := gwizo.Ingest("gwizo")

    // Measure value [v]vc{m}[c]
    fmt.Printf("%s has Measure value %d \n", octopus.Word, octopus.Measure)

    // Number of Vowels
    v := strings.Count(octopus.VowCon, "v")
    fmt.Printf("%s Has %d Vowels \n", octopus.Word, v)

    // Number of Consonants
    c := strings.Count(octopus.VowCon, "c")
    fmt.Printf("%s Has %d Consonants\n", octopus.Word, c)

    // Or just print all of the values
    fmt.Println("values: ", octopus) // Type Analyse implements the Stringer interface.
  }
  Results
  ---------------------
  gwizo has Measure value 1
  gwizo Has 2 Vowels
  gwizo Has 3 Consonants
  values:  {gwizo ccvcv 1}
</pre>

Access Each Step Directly
------------------------------
This returns a stem that has been stemmed with a specific Step. Step_1a(),
Step_1b(), Step_1c(), Step_2(), Step_3(), Step_4(), Step_5a(), Step_5b()
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
