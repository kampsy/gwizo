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
  VowCon string // example vcvcvc. Where v = vowel and c = consonant
  Measure int // Number of times the pair vc appears
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
  anl.VowCon = str
  anl.Measure = strings.Count(str, "vc")

  return anl
}

// implementation of String Method and so Stringer interface
func (a *Analyse) String() string {
  return fmt.Sprintf("%s %s %s", a.Word, a.VowCon, a.Measure)
}

// Method HasVowel returns bool (*v*)
func (a *Analyse) HasVowel() bool {
  return strings.Contains(a.VowCon, "v")
}

// Method HasConsonant returns bool (*c*)
func (a *Analyse) HasConsonant() bool {
  return strings.Contains(a.VowCon, "c")
}

// Measure = 1 and VowCon pattern ends with cvc, where second c is not
// W, X, y
func (a *Analyse) HasM1Endscvc() bool {
  cvc := strings.HasSuffix(a.VowCon, "cvc")
  wlen := len(a.Word)
  lastLetter := a.Word[(wlen - 1)]
  str := string(lastLetter)
  w := strings.Contains(str, "w")
  x := strings.Contains(str, "x")
  y := strings.Contains(str, "y")

  if cvc == true && w == false && x == false && y == false {
    return true
  }else {
    return false
  }
}


// Step 1A according the stemmer doc.
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

  // For S suffix. S ->
  s := strings.HasSuffix(a.Word, "s")
  if s == true && sses == false && ies == false && ss == false {
    pre := strings.TrimSuffix(a.Word, "s")
    str = pre
  }

  return str
}


// Step 1B according the stemmer doc.
func (a *Analyse) Step_1b() string {
  var str string

  // Word Measure (m > 0) and EED suffix. EED -> EE
  eed := strings.HasSuffix(a.Word, "eed")
  if eed == true && a.Measure > 0 {
    pre := strings.TrimSuffix(a.Word, "eed")
    str = pre + "ee"
  }

  // Word has Vowel and ED suffix. ED ->
  ed := strings.HasSuffix(a.Word, "ed")
  if a.HasVowel() == true && ed == true && eed == false {
    // word exception
    if a.Word == "bled" {
      str = "bled"
    }else {
      pre := strings.TrimSuffix(a.Word, "ed")
      str = pre
    }
  }

  // Word has Vowel and ING suffix. ING ->
  ing := strings.HasSuffix(a.Word, "ing")
  if a.HasVowel() == true && ing == true {
    // word exception
    if a.Word == "sing" {
      str = "sing"
    }else {
      pre := strings.TrimSuffix(a.Word, "ing")
      str = pre
    }
  }

  // If the second or third of the above rules is successful the following
  // is done.
  at := strings.HasSuffix(a.Word, "at")
  if at == true || ing == true || ed == true && a.HasM1Endscvc() == true {
    str = str + "e"
  }

  return str
}
