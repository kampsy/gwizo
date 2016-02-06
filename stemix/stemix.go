/*
stemix is the implementation of the porter stemmer algorithm in go. Specificaly
the suffix stripping. M.F.Porter 1980
stemix does not use a stem dictionary. It reads a documents and returns a slice
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

// Measure value is grater than 0
func (a *Analyse) HasMeaGreater0() bool {
  fmt.Println(a.Measure)
  if a.Measure > 0 {
    return true
  }else {
    return false
  }
}

// Function checks if VowCon pattern ends with cvc, where second c is not
// W, X, Y
func HasEndcvcNotwxy(str string) bool {
  nest := Form(str)
  cvc := strings.HasSuffix(nest.VowCon, "cvc")
  fmt.Println(nest.VowCon)
  wlen := len(nest.Word)
  lastLetter := nest.Word[(wlen - 1)]
  word := string(lastLetter)
  w := strings.Contains(word, "w")
  x := strings.Contains(word, "x")
  y := strings.Contains(word, "y")

  if cvc == true && w == false && x == false && y == false {
    return true
  }else {
    return false
  }
}

// Measure value = 1
func HasMeasure1(str string) bool {
  nest := Form(str)
  fmt.Println(nest.Measure)
  if nest.Measure == 1 {
    return true
  }else {
    return false
  }
}

// Function accepts a string as an argument, checks if it has double consonant
// as suffix and returns a boolean
func HasDoubleConsonant(str string) bool {
  nest := Form(str)
  cc := strings.HasSuffix(nest.VowCon, "cc")
  return cc
}


/* Step 1A according the stemmer doc.*/
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


/* Step 1B according the stemmer doc.*/
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

  /* If the second or third of the above rules is successful the following
  is done.*/

  // Word has AT suffix. AT -> ATE
  if ed == true || ing == true {
    at := strings.HasSuffix(str, "at")
    if at == true {
      pre := strings.TrimSuffix(str, "at")
      str = pre + "ate"
    }
  }

  // Word has BL suffix. BL -> BLE
  if ed == true || ing == true {
    bl := strings.HasSuffix(str, "bl")
    if bl == true {
      pre := strings.TrimSuffix(str, "bl")
      str = pre + "ble"
    }
  }

  // Word has IZ suffix. IZ -> IZE
  if ed == true || ing == true {
    iz := strings.HasSuffix(str, "iz")
    if iz == true {
      pre := strings.TrimSuffix(str, "iz")
      str = pre + "ize"
    }
  }

  // (m=1 and *o) -> E
  if ed == true || ing == true {
    if HasMeasure1(str) == true && HasEndcvcNotwxy(str) == true {
      str = str + "e"
    }
  }

  // (*d and not (*L or *S or *Z)) -> single letter at the end
  if ed == true || ing == true {
    if HasDoubleConsonant(str) == true {
      ll := strings.HasSuffix(str, "ll")
      ss := strings.HasSuffix(str, "ss")
      zz := strings.HasSuffix(str, "zz")
      if ll == true || ss == true || zz == true {
        str = str
      }else {
        strlen := len(str)
        lastLetter := str[:(strlen-1)]
        str = string(lastLetter)
      }
    }
  }

  return str
}


/* Step 1c according the stemmer doc.*/
func (a *Analyse) Step_1c() string {
  var str string

  // Word has Vowel and Y suffix. Y -> I
  y := strings.HasSuffix(a.Word, "y")
  if a.HasVowel() == true && y == true {
    pre := strings.TrimSuffix(a.Word, "y")
    str = pre + "i"
  }else {
    str = a.Word
  }

  return str
}


/* Step 2 according the stemmer doc.
=========================================*/
func (a *Analyse) Step_2() string {
  var str string

  if a.HasMeaGreater0() == true {
    // For ATIONAL suffix. ATIONAL -> ATE
    ational := strings.HasSuffix(a.Word, "ational")
    if ational == true {
      pre := strings.TrimSuffix(a.Word, "ational")
      str = pre + "ate"
    }

  }

  return str
}
