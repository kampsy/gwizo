![gwizo](https://github.com/kampsy/gwizo/img/gwizo.png)
## gwizo

gwizo |pronounced as [guizo]| is the Next generation Native Go implementation of the
Porter Stemming algorithm (An algorithm for suffix stripping M.F.Porter 1980 see: http://tartarus.org/martin/PorterStemmer/def.txt).
Gwizo is Better than other Go implementations of the algorithm because it Provides
powerful features Like the Measure([v]vc{m}[c]) of the word, how many Vowel and Consonants
the word has or had?, enables developers to choose which step(Step_1a, Step_2 etc)
to use on a word etc. Gwizo was built from the ground up with modularity in mind.
Developers have direct access to the Core parts of the algorithm, which results in doing
lots of cool, powerful and amazing things (see examples below).

Gwizo is built for people doing Machine Learning algorithms, specifically
Natural language processing (NLP).

Note: I made a few modification to gwizo for it to pass all tests. The original algorithm
in http://tartarus.org/martin/PorterStemmer/def.txt) has a few mistakes.

Installation
------------
<pre>
  go get github.com/kampsy/gwizo
</pre>

[[[[[ Examples ]]]]]

The Stem
------------------------------
This returns the stem of your string.
<pre>
  package main

  import (
    "fmt"
    "github.com/kampsy/gwizo"
    )

    func main() {
      str := gwizo.Stem("Consonants")
      fmt.Println(str)
    }
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
      data := gwizo.Form("abatements")

      // Measure value [v]vc{m}[c]
      fmt.Printf("Your word %s has measure value %d \n", data.Word, data.Measure)

      // Number of Vowels
      v := strings.Count(data.VowCon, "v")
      fmt.Printf("Your word %s Has %d Vowels \n", data.Word, v)

      // Number of Consonants
      c := strings.Count(data.VowCon, "c")
      fmt.Printf("Your word %s Has %d Consonants\n", data.Word, c)

      // Or just print all of the values
      fmt.Println(data) // Type Analyse implements the Stringer interface.
    }
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
      data := gwizo.Form("abatements")

      // Stem only with Step_1a
      fmt.Println(data.Step_1a())

      // Stem only with Step_4
      fmt.Println(data.Step_4())

      // Stem only with Step_1b
      fmt.Println(data.Step_1b())

      // You get the idea!
    }
</pre>
