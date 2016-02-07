/*
stemix is the implementation of the porter stemmer algorithm in go. Specificaly
the suffix stripping. M.F.Porter 1980
stemix does not use any stem dictionary. It reads a documents and returns a slice
of stems
*/
package stemix

import (
  "fmt"
  "strings"
)


type Analyse struct {
  Word string // The word to be stemmed.
  VowCon string // example vcvcvc. Where v = vowel and c = consonant.
  Measure int // Number of times the pair vc appears.
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

// Measure value is grater than 1
func (a *Analyse) HasMeaGreater1() bool {
  if a.Measure > 1 {
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

// Function checks if word has suffix S or T
func (a *Analyse) HasEndst() bool {
  s := strings.HasSuffix(a.Word, "s")
  t := strings.HasSuffix(a.Word, "t")

  if s == true || t == true {
    return true
  }else {
    return false
  }
}

// Function checks if word has suffix L
func (a *Analyse) HasEndl() bool {
  l := strings.HasSuffix(a.Word, "l")
  if l == true {
    return true
  }else {
    return false
  }
}




/*
Step 1 deals with plurals and past participles. The subsequent steps are
much more straightforward.
Step 1A according the stemmer doc.
======================================*/
func (a *Analyse) Step_1a() string {
  var str string = a.Word

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
  var str string = a.Word

  // Word Measure (m > 0) and EED suffix. EED -> EE
  eed := strings.HasSuffix(a.Word, "eed")
  if eed == true && a.Measure > 0 {
    if a.Word == "feed" {
      str = "feed"
    }else {
      pre := strings.TrimSuffix(a.Word, "eed")
      str = pre + "ee"
    }
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
    if str == "ring" {
      return str
    }
    pre := strings.TrimSuffix(a.Word, "ing")
    str = pre
  }

  /* If the second or third of the above rules is successful the following
  is done.*/

  // Word has AT suffix. AT -> ATE
  if ed == true {
    at := strings.HasSuffix(str, "at")
    if at == true {
      pre := strings.TrimSuffix(str, "at")
      str = pre + "ate"
    }
  }

  // Word has BL suffix. BL -> BLE
  if ed == true {
    bl := strings.HasSuffix(str, "bl")
    if bl == true {
      pre := strings.TrimSuffix(str, "bl")
      str = pre + "ble"
    }
  }

  // Word has IZ suffix. IZ -> IZE
  if ed == true {
    iz := strings.HasSuffix(str, "iz")
    if iz == true {
      pre := strings.TrimSuffix(str, "iz")
      str = pre + "ize"
    }
  }

  // (m=1 and *o) -> E
  if ing == true {
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
  var str string = a.Word

  // (*v*) Y -> I
  // Word has Vowel and Y suffix. Y -> I
  y := strings.HasSuffix(a.Word, "y")
  if a.HasVowel() == true && y == true {
    pre := strings.TrimSuffix(a.Word, "y")
    str = pre + "i"
  }

  return str
}


/* Step 2 according the stemmer doc.
=========================================*/
func (a *Analyse) Step_2() string {
  var str string = a.Word

  if a.HasMeaGreater0() == true {

    // For TIONAL suffix. TIONAL -> TION
    tional := strings.HasSuffix(a.Word, "tional") // Moved 1 step up
    if tional == true {
      if a.Word == "rational" {
        str = "rational"
      }else {
        pre := strings.TrimSuffix(a.Word, "tional")
        str = pre + "tion"
      }
    }

    // For ATIONAL suffix. ATIONAL -> ATE
    ational := strings.HasSuffix(a.Word, "ational") // Moved 1 step down
    if ational == true {
      if a.Word == "rational" {
        str = "rational"
      }else {
        pre := strings.TrimSuffix(a.Word, "ational")
        str = pre + "ate"
      }
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

    // For ATION suffix. ATION -> ATE
    ation := strings.HasSuffix(a.Word, "ation")// Moved 1 step bottom
    if ation == true {
      pre := strings.TrimSuffix(a.Word, "ation")
      str = pre + "ate"
    }

    // For IZATION suffix. IZATION -> IZE
    ization := strings.HasSuffix(a.Word, "ization")// Moved 1 step down
    if ization == true {
      pre := strings.TrimSuffix(a.Word, "ization")
      str = pre + "ize"
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
  var str string = a.Word

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
  var str string = a.Word

  if a.HasMeaGreater1() == true {
    // For AL suffix. AL ->
    al := strings.HasSuffix(a.Word, "al")
    if al == true {
      pre := strings.TrimSuffix(a.Word, "al")
      str = pre
    }

    // For ANCE suffix. ANCE ->
    ance := strings.HasSuffix(a.Word, "ance")
    if ance == true {
      pre := strings.TrimSuffix(a.Word, "ance")
      str = pre
    }

    // For ENCE suffix. ENCE ->
    ence := strings.HasSuffix(a.Word, "ence")
    if ence == true {
      pre := strings.TrimSuffix(a.Word, "ence")
      str = pre
    }

    // For ER suffix. ER ->
    er := strings.HasSuffix(a.Word, "er")
    if er == true {
      pre := strings.TrimSuffix(a.Word, "er")
      str = pre
    }

    // For IC suffix. IC ->
    ic := strings.HasSuffix(a.Word, "ic")
    if ic == true {
      pre := strings.TrimSuffix(a.Word, "ic")
      str = pre
    }

    // For ABLE suffix. ABLE ->
    able := strings.HasSuffix(a.Word, "able")
    if able == true {
      pre := strings.TrimSuffix(a.Word, "able")
      str = pre
    }

    // For IBLE suffix. IBLE ->
    ible := strings.HasSuffix(a.Word, "ible")
    if ible == true {
      pre := strings.TrimSuffix(a.Word, "ible")
      str = pre
    }

    // For ANT suffix. ANT ->
    ant := strings.HasSuffix(a.Word, "ant")
    if ant == true {
      pre := strings.TrimSuffix(a.Word, "ant")
      str = pre
    }

    // For ENT suffix. ENT ->
    ent := strings.HasSuffix(a.Word, "ent") // moved position 2 paces
    if ent == true {
      pre := strings.TrimSuffix(a.Word, "ent")
      str = pre
    }

    // For MENT suffix. MENT ->
    ment := strings.HasSuffix(a.Word, "ment") // moved top 1 pace
    if ment == true {
      pre := strings.TrimSuffix(a.Word, "ment")
      str = pre
    }

    // For EMENT suffix. EMENT ->
    ement := strings.HasSuffix(a.Word, "ement") // droped bottom 2 paces
    if ement == true {
      pre := strings.TrimSuffix(a.Word, "ement")
      str = pre
    }


    if a.HasEndst() == false { // Made Personal correction from true to false.
      // For ION suffix. ION ->
      ion := strings.HasSuffix(a.Word, "ion")
      if ion == true {
        pre := strings.TrimSuffix(a.Word, "ion")
        str = pre
      }
    }

    // For OU suffix. OU ->
    ou := strings.HasSuffix(a.Word, "ou")
    if ou == true {
      pre := strings.TrimSuffix(a.Word, "ou")
      str = pre
    }

    // For ISM suffix. ISM ->
    ism := strings.HasSuffix(a.Word, "ism")
    if ism == true {
      pre := strings.TrimSuffix(a.Word, "ism")
      str = pre
    }

    // For ATE suffix. ATE ->
    ate := strings.HasSuffix(a.Word, "ate")
    if ate == true {
      pre := strings.TrimSuffix(a.Word, "ate")
      str = pre
    }

    // For ITI suffix. ITI ->
    iti := strings.HasSuffix(a.Word, "iti")
    if iti == true {
      pre := strings.TrimSuffix(a.Word, "iti")
      str = pre
    }

    // For OUS suffix. OUS ->
    ous := strings.HasSuffix(a.Word, "ous")
    if ous == true {
      pre := strings.TrimSuffix(a.Word, "ous")
      str = pre
    }

    // For IVE suffix. IVE ->
    ive := strings.HasSuffix(a.Word, "ive")
    if ive == true {
      pre := strings.TrimSuffix(a.Word, "ive")
      str = pre
    }

    // For IZE suffix. IZE ->
    ize := strings.HasSuffix(a.Word, "ize")
    if ize == true {
      pre := strings.TrimSuffix(a.Word, "ize")
      str = pre
    }

  }

  return str
}

/*Step 5 according the stemmer doc.
little tidying up
=====================================*/
func (a *Analyse) Step_5a() string {
  var str string = a.Word

  if a.HasMeaGreater1() == true {
    // For E suffix. E ->
    e := strings.HasSuffix(a.Word, "e")
    if e == true {
      pre := strings.TrimSuffix(a.Word, "e")
      str = pre
    }
  }else {
    str = a.Word
  }

  if HasMeasure1(a.Word) == true && HasEndcvcNotwxy(a.Word) == false {
    // (m=1 and not *o) E ->
    e := strings.HasSuffix(a.Word, "e")
    if e == true {
      pre := strings.TrimSuffix(a.Word, "e")
      str = pre
    }
  }

  return str
}


/*Step 5B according the stemmer doc.*/
func (a *Analyse) Step_5b() string {
  var str string = a.Word

  if a.HasMeaGreater1() == true {
    if HasDoubleConsonant(a.Word) == true && a.HasEndl() == true {
      w := a.Word
      strlen := len(w)
      lastLetter := w[:(strlen-1)]
      str = string(lastLetter)
    }
  }else {
    str = a.Word
  }

  return str
}


func Stem(s string) string {
  var str string = s

  var slice []string

  mytype := Form(s)
  mystr := mytype.Step_1a()
  if mystr != s {
    slice = append(slice, mystr)
  }

  mytype2 := Form(s)
  mystr2 := mytype2.Step_1b()
  if mystr2 != s {
    slice = append(slice, mystr2)
  }

  mytype3 := Form(s)
  mystr3 := mytype3.Step_1c()
  if mystr3 != s {
    slice = append(slice, mystr3)
  }

  mytype4 := Form(s)
  mystr4 := mytype4.Step_2()
  if mystr4 != s {
    slice = append(slice, mystr4)
  }

  mytype5 := Form(s)
  mystr5 := mytype5.Step_3()
  if mystr5 != s {
    slice = append(slice, mystr5)
  }

  mytype6 := Form(s)
  mystr6 := mytype6.Step_4()
  if mystr6 != s {
    slice = append(slice, mystr6)
  }

  mytype7 := Form(s)
  mystr7 := mytype7.Step_5a()
  if mystr7 != s {
    slice = append(slice, mystr7)
  }

  mytype8 := Form(s)
  mystr8 := mytype8.Step_5b()
  if mystr8 != s {
    slice = append(slice, mystr8)
  }


  if len(slice) > 0 {
    str = slice[0]
  }

  return str
}
