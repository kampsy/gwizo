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


/*
Step 1 deals with plurals and past participles. The subsequent steps are
much more straightforward.
Step 1A according the stemmer doc.
======================================*/
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

    // For TIONAL suffix. TIONAL -> TION
    tional := strings.HasSuffix(a.Word, "tional")
    if tional == true {
      pre := strings.TrimSuffix(a.Word, "tional")
      str = pre + "tion"
    }

    // For ENCI suffix. ENCI -> ENCE
    enci := strings.HasSuffix(a.Word, "enci")
    if enci == true {
      pre := strings.TrimSuffix(a.Word, "enci")
      str = pre + "ence"
    }

    // For ANCI suffix. ANCI -> ANCE
    anci := strings.HasSuffix(a.Word, "anci")
    if anci == true {
      pre := strings.TrimSuffix(a.Word, "anci")
      str = pre + "ance"
    }

    // For IZER suffix. IZER -> IZE
    izer := strings.HasSuffix(a.Word, "izer")
    if izer == true {
      pre := strings.TrimSuffix(a.Word, "izer")
      str = pre + "ize"
    }

    // For ABLI suffix. ABLI -> ABLE
    abli := strings.HasSuffix(a.Word, "abli")
    if abli == true {
      pre := strings.TrimSuffix(a.Word, "abli")
      str = pre + "able"
    }

    // For ALLI suffix. ALLI -> AL
    alli := strings.HasSuffix(a.Word, "alli")
    if alli == true {
      pre := strings.TrimSuffix(a.Word, "alli")
      str = pre + "al"
    }

    // For ENTLI suffix. ENTLI -> ENT
    entli := strings.HasSuffix(a.Word, "entli")
    if entli == true {
      pre := strings.TrimSuffix(a.Word, "entli")
      str = pre + "ent"
    }

    // For ELI suffix. ELI -> E
    eli := strings.HasSuffix(a.Word, "eli")
    if eli == true {
      pre := strings.TrimSuffix(a.Word, "eli")
      str = pre + "e"
    }

    // For OUSLI suffix. OUSLI -> OUS
    ousli := strings.HasSuffix(a.Word, "ousli")
    if ousli == true {
      pre := strings.TrimSuffix(a.Word, "ousli")
      str = pre + "ous"
    }

    // For IZATION suffix. IZATION -> IZE
    ization := strings.HasSuffix(a.Word, "ization")
    if ization == true {
      pre := strings.TrimSuffix(a.Word, "ization")
      str = pre + "ize"
    }

    // For ATION suffix. ATION -> ATE
    ation := strings.HasSuffix(a.Word, "ation")
    if ation == true {
      pre := strings.TrimSuffix(a.Word, "ation")
      str = pre + "ate"
    }

    // For ATOR suffix. ATOR -> ATE
    ator := strings.HasSuffix(a.Word, "ator")
    if ator == true {
      pre := strings.TrimSuffix(a.Word, "ator")
      str = pre + "ate"
    }

    // For ALISM suffix. ALISM -> AL
    alism := strings.HasSuffix(a.Word, "alism")
    if alism == true {
      pre := strings.TrimSuffix(a.Word, "alism")
      str = pre + "al"
    }

    // For IVENESS suffix. IVENESS -> IVE
    iveness := strings.HasSuffix(a.Word, "iveness")
    if iveness == true {
      pre := strings.TrimSuffix(a.Word, "iveness")
      str = pre + "ive"
    }

    // For FULNESS suffix. FULNESS -> FUL
    fulness := strings.HasSuffix(a.Word, "fulness")
    if fulness == true {
      pre := strings.TrimSuffix(a.Word, "fulness")
      str = pre + "ful"
    }

    // For OUSNESS suffix. OUSNESS -> OUS
    ousness := strings.HasSuffix(a.Word, "ousness")
    if ousness == true {
      pre := strings.TrimSuffix(a.Word, "ousness")
      str = pre + "ous"
    }

    // For ALITI suffix. ALITI -> AL
    aliti := strings.HasSuffix(a.Word, "aliti")
    if aliti == true {
      pre := strings.TrimSuffix(a.Word, "aliti")
      str = pre + "al"
    }

    // For IVITI suffix. IVITI -> IVE
    iviti := strings.HasSuffix(a.Word, "iviti")
    if iviti == true {
      pre := strings.TrimSuffix(a.Word, "iviti")
      str = pre + "ive"
    }

    // For BILITI suffix. BILITI -> BLE
    biliti := strings.HasSuffix(a.Word, "biliti")
    if biliti == true {
      pre := strings.TrimSuffix(a.Word, "biliti")
      str = pre + "ble"
    }

  }

  return str
}


/*Step 3 according the stemmer doc.
=====================================*/
func (a *Analyse) Step_3() string {
  var str string

  if a.HasMeaGreater0() == true {
    // For ICATE suffix. ICATE -> IC
    icate := strings.HasSuffix(a.Word, "icate")
    if icate == true {
      pre := strings.TrimSuffix(a.Word, "icate")
      str = pre + "ic"
    }

    // For ATIVE suffix. ATIVE ->
    ative := strings.HasSuffix(a.Word, "ative")
    if ative == true {
      pre := strings.TrimSuffix(a.Word, "ative")
      str = pre
    }

    // For ALIZE suffix. ALIZE -> AL
    alize := strings.HasSuffix(a.Word, "alize")
    if alize == true {
      pre := strings.TrimSuffix(a.Word, "alize")
      str = pre + "al"
    }

    // For ICITI suffix. ICITI -> IC
    iciti := strings.HasSuffix(a.Word, "iciti")
    if iciti == true {
      pre := strings.TrimSuffix(a.Word, "iciti")
      str = pre + "ic"
    }

    // For ICAL suffix. ICAL -> IC
    ical := strings.HasSuffix(a.Word, "ical")
    if ical == true {
      pre := strings.TrimSuffix(a.Word, "ical")
      str = pre + "ic"
    }

    // For FUL suffix. FUL ->
    ful := strings.HasSuffix(a.Word, "ful")
    if ful == true {
      pre := strings.TrimSuffix(a.Word, "ful")
      str = pre
    }

    // For NESS suffix. NESS ->
    ness := strings.HasSuffix(a.Word, "ness")
    if ness == true {
      pre := strings.TrimSuffix(a.Word, "ness")
      str = pre
    }

  }

  return str
}



/*Step 4 according the stemmer doc.
The suffixes will now be removed
=====================================*/
func (a *Analyse) Step_4() string {
  
}
