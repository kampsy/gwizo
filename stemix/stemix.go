/*
Arach is the implementation of the porter stemmer algorithm in go. Specificaly
the suffix stripping. M.F.Porter 1980
Arach does not use a stem dictionary. It reads a documents and returns a slice
of stems
*/
package stemix

import (
  "fmt"
  "strings"
)


type Analyse struct {
  Word string
  Form string // example vcvcvc. Where v = vowel and c = consonant
  Measure int // Number of times the pair vc appears
}
func (a *Analyse) String() string {
  return fmt.Sprintf("%s %s %s", a.Word, a.Form, a.Measure)
}

// returns type Analyse
func Form(w string) Analyse {
  var inx []string
  // Change the word to lowercase letters.
  wordLower := strings.ToLower(w)
  for i := 0; i < len(wordLower); i++ {
    if string(wordLower[i]) == "a" || string(wordLower[i]) == "e" ||
    string(wordLower[i]) == "i" || string(wordLower[i]) == "o" ||
    string(wordLower[i]) == "u" {
      inx = append(inx, "v")
    }else if string(wordLower[i]) != "a" || string(wordLower[i]) != "e" ||
    string(wordLower[i]) != "i" || string(wordLower[i]) != "o" ||
    string(wordLower[i]) != "u" {
      inx = append(inx, "c")
    }
  }

  str := strings.Join(inx, "")
  var anl Analyse // Instance of Analyse.
  anl.Word  = w
  anl.Form = str
  anl.Measure = strings.Count(str, "vc")

  return anl
}


func (a *Analyse) Step_1a() string {
  var str string

  // For SSES suffix. SSES -> SS
  sses := strings.HasSuffix(a.Word, "sses")
  if sses == true {
    pre := strings.TrimSuffix(a.Word, "sses")
    str = pre + "ss"
  }

  // For IES suffix. IES -> I
  ies := strings.HasSuffix(a.Word, "ies")
  if ies == true {
    pre := strings.TrimSuffix(a.Word, "ies")
    str = pre + "i"
  }

  // For SS suffix. SS -> SS
  ss := strings.HasSuffix(a.Word, "ss")
  if ss == true {
    pre := strings.TrimSuffix(a.Word, "ss")
    str = pre + "ss"
  }

  // For S suffix. S -> S
  s := strings.HasSuffix(a.Word, "s")
  if s == true && sses == false && ies == false && ss == false {
    pre := strings.TrimSuffix(a.Word, "s")
    str = pre
  }

  return str
}
